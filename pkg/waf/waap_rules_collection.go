package waf

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	corazatypes "github.com/crowdsecurity/coraza/v3/types"
	"github.com/crowdsecurity/crowdsec/pkg/csconfig"
	"github.com/crowdsecurity/crowdsec/pkg/cwhub"
	"gopkg.in/yaml.v2"

	log "github.com/sirupsen/logrus"
)

// to be filled w/ seb update
type WaapCollection struct {
	collectionName string
	Rules          []string
}

// to be filled w/ seb update
type WaapCollectionConfig struct {
	Type              string   `yaml:"type"`
	Name              string   `yaml:"name"`
	SecLangFilesRules []string `yaml:"seclang_files_rules"`
	SecLangRules      []string `yaml:"seclang_rules"`
}

func LoadCollection(collection string) (WaapCollection, error) {

	//FIXME: do it once globally
	waapRules := make(map[string]WaapCollectionConfig)

	for _, hubWafRuleItem := range cwhub.GetItemMap(cwhub.WAF_RULES) {
		log.Infof("loading %s", hubWafRuleItem.LocalPath)
		if !hubWafRuleItem.Installed {
			continue
		}

		content, err := os.ReadFile(hubWafRuleItem.LocalPath)

		if err != nil {
			log.Warnf("unable to read file %s : %s", hubWafRuleItem.LocalPath, err)
			continue
		}

		var rule WaapCollectionConfig

		err = yaml.Unmarshal(content, &rule)

		if err != nil {
			log.Warnf("unable to unmarshal file %s : %s", hubWafRuleItem.LocalPath, err)
			continue
		}

		if rule.Type != "waf-rule" { //FIXME: rename to waap-rule when hub is properly updated
			log.Warnf("unexpected type %s instead of waap-rule for file %s", rule.Type, hubWafRuleItem.LocalPath)
			continue
		}
		log.Infof("Adding %s to waap rules", rule.Name)
		waapRules[rule.Name] = rule
	}

	if len(waapRules) == 0 {
		return WaapCollection{}, fmt.Errorf("no waap rules found in hub")
	}

	var loadedRule WaapCollectionConfig
	var ok bool

	if loadedRule, ok = waapRules[collection]; !ok {
		return WaapCollection{}, fmt.Errorf("no waap rules found for collection %s", collection)
	}

	waapCol := WaapCollection{
		collectionName: loadedRule.Name,
	}

	if loadedRule.SecLangFilesRules != nil {
		for _, rulesFile := range loadedRule.SecLangFilesRules {
			fullPath := filepath.Join(csconfig.DataDir, rulesFile)
			c, err := os.ReadFile(fullPath)
			if err != nil {
				log.Errorf("unable to read file %s : %s", rulesFile, err)
				continue
			}
			for _, line := range strings.Split(string(c), "\n") {
				if strings.HasPrefix(line, "#") {
					continue
				}
				if strings.TrimSpace(line) == "" {
					continue
				}
				waapCol.Rules = append(waapCol.Rules, line)
			}
		}
	}

	if loadedRule.SecLangRules != nil {
		waapCol.Rules = append(waapCol.Rules, loadedRule.SecLangRules...)
	}

	return waapCol, nil
}

func (wcc WaapCollectionConfig) LoadCollection(collection string) (WaapCollection, error) {
	return WaapCollection{}, nil
}

func (w WaapCollection) Check() error {
	return nil
}

func (w WaapCollection) Eval(req ParsedRequest) (*corazatypes.Interruption, error) {
	return nil, nil
}

func (w WaapCollection) GetDisplayName() string {
	return w.collectionName
}