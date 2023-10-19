package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"slices"
	"sort"
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	"github.com/crowdsecurity/crowdsec/pkg/cwhub"
)


func selectItems(hub *cwhub.Hub, itemType string, args []string, installedOnly bool) ([]string, error) {
	itemNames := hub.GetItemNames(itemType)

	notExist := []string{}
	if len(args) > 0 {
		installedOnly = false
		for _, arg := range args {
			if !slices.Contains(itemNames, arg) {
				notExist = append(notExist, arg)
			}
		}
	}

	if len(notExist) > 0 {
		return nil, fmt.Errorf("item(s) '%s' not found in %s", strings.Join(notExist, ", "), itemType)
	}

	if len(args) > 0 {
		itemNames = args
	}

	if installedOnly {
		installed := []string{}
		for _, item := range itemNames {
			if hub.GetItem(itemType, item).Installed {
				installed = append(installed, item)
			}
		}
		return installed, nil
	}
	return itemNames, nil
}


func ListItems(out io.Writer, itemTypes []string, args []string, showType bool, showHeader bool, all bool) error {
	var err error

	hub, err := cwhub.GetHub()
	if err != nil {
		return err
	}

	items := make(map[string][]string)
	for _, itemType := range itemTypes {
		if items[itemType], err = selectItems(hub, itemType, args, !all); err != nil {
			return err
		}
	}
		
	if csConfig.Cscli.Output == "human" {
		for _, itemType := range itemTypes {
			listHubItemTable(out, "\n"+strings.ToUpper(itemType), itemType, items[itemType])
		}
	} else if csConfig.Cscli.Output == "json" {
		type itemHubStatus struct {
			Name         string `json:"name"`
			LocalVersion string `json:"local_version"`
			LocalPath    string `json:"local_path"`
			Description  string `json:"description"`
			UTF8Status   string `json:"utf8_status"`
			Status       string `json:"status"`
		}

		hubStatus := make(map[string][]itemHubStatus)
		for _, itemType := range itemTypes {
			// empty slice in case there are no items of this type
			hubStatus[itemType] = make([]itemHubStatus, len(items[itemType]))
			for i, itemName := range items[itemType] {
				item := hub.GetItem(itemType, itemName)
				status, emo := item.Status()
				hubStatus[itemType][i] = itemHubStatus{
					Name:         item.Name,
					LocalVersion: item.LocalVersion,
					LocalPath:    item.LocalPath,
					Description:  item.Description,
					Status:       status,
					UTF8Status:   fmt.Sprintf("%v  %s", emo, status),
				}
			}
			h := hubStatus[itemType]
			sort.Slice(h, func(i, j int) bool { return h[i].Name < h[j].Name })
		}
		x, err := json.MarshalIndent(hubStatus, "", " ")
		if err != nil {
			log.Fatalf("failed to unmarshal")
		}
		out.Write(x)
	} else if csConfig.Cscli.Output == "raw" {
		csvwriter := csv.NewWriter(out)
		if showHeader {
			header := []string{"name", "status", "version", "description"}
			if showType {
				header = append(header, "type")
			}
			err := csvwriter.Write(header)
			if err != nil {
				log.Fatalf("failed to write header: %s", err)
			}

		}
		for _, itemType := range itemTypes {
			for _, itemName := range items[itemType] {
				item := hub.GetItem(itemType, itemName)
				status, _ := item.Status()
				if item.LocalVersion == "" {
					item.LocalVersion = "n/a"
				}
				row := []string{
					item.Name,
					status,
					item.LocalVersion,
					item.Description,
				}
				if showType {
					row = append(row, itemType)
				}
				err := csvwriter.Write(row)
				if err != nil {
					log.Fatalf("failed to write raw output : %s", err)
				}
			}
		}
		csvwriter.Flush()
	}
	return nil
}

func InspectItem(name string, itemType string, noMetrics bool) error {
	hub, err := cwhub.GetHub()
	if err != nil {
		return err
	}

	hubItem := hub.GetItem(itemType, name)
	if hubItem == nil {
		return fmt.Errorf("can't find '%s' in %s", name, itemType)
	}

	var b   []byte

	switch csConfig.Cscli.Output {
	case "human", "raw":
		b, err = yaml.Marshal(*hubItem)
		if err != nil {
			return fmt.Errorf("unable to marshal item: %s", err)
		}
	case "json":
		b, err = json.MarshalIndent(*hubItem, "", " ")
		if err != nil {
			return fmt.Errorf("unable to marshal item: %s", err)
		}
	}

	fmt.Printf("%s", string(b))

	if noMetrics || csConfig.Cscli.Output == "json" || csConfig.Cscli.Output == "raw" {
		return nil
	}

	fmt.Printf("\nCurrent metrics: \n")
	ShowMetrics(hub, hubItem)

	return nil
}