package waap_rule

import (
	"fmt"
	"hash/fnv"
	"strings"
)

type ModsecurityRule struct {
}

var zonesMap map[string]string = map[string]string{
	"ARGS":            "ARGS_GET",
	"ARGS_NAMES":      "ARGS_GET_NAMES",
	"BODY_ARGS":       "ARGS_POST",
	"BODY_ARGS_NAMES": "ARGS_POST_NAMES",
	"HEADERS":         "REQUEST_HEADERS",
	"METHOD":          "REQUEST_METHOD",
	"PROTOCOL":        "REQUEST_PROTOCOL",
	"URI":             "REQUEST_URI",
}

var transformMap map[string]string = map[string]string{
	"lowercase": "t:lowercase",
	"uppercase": "t:uppercase",
	"b64decode": "t:base64Decode",
	"hexdecode": "t:hexDecode",
	"length":    "t:length",
}

var matchMap map[string]string = map[string]string{
	"regex":           "@rx",
	"equal":           "@streq",
	"startsWith":      "@beginsWith",
	"endsWith":        "@endsWith",
	"contains":        "@contains",
	"libinjectionSQL": "@detectSQLi",
	"libinjectionXSS": "@detectXSS",
	"gt":              "@gt",
	"lt":              "@lt",
	"ge":              "@ge",
	"le":              "@le",
}

func (m ModsecurityRule) Build(rule *CustomRule, waapRuleName string) (string, error) {

	rules, err := m.buildRules(rule, waapRuleName, false)

	if err != nil {
		return "", err
	}

	return strings.Join(rules, "\n"), nil
}

func (m ModsecurityRule) generateRuleID(rule *CustomRule, waapRuleName string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(waapRuleName))
	h.Write([]byte(rule.Match.Type))
	h.Write([]byte(rule.Match.Value))
	for _, zone := range rule.Zones {
		h.Write([]byte(zone))
	}
	for _, transform := range rule.Transform {
		h.Write([]byte(transform))
	}
	return h.Sum32()
}

func (m ModsecurityRule) buildRules(rule *CustomRule, waapRuleName string, and bool) ([]string, error) {
	ret := make([]string, 0)

	if rule.And != nil {
		for c, andRule := range rule.And {
			subName := fmt.Sprintf("%s_and_%d", waapRuleName, c)
			lastRule := c == len(rule.And)-1
			rules, err := m.buildRules(&andRule, subName, !lastRule)
			if err != nil {
				return nil, err
			}
			ret = append(ret, rules...)
		}
	}

	if rule.Or != nil {
		for c, orRule := range rule.Or {
			subName := fmt.Sprintf("%s_or_%d", waapRuleName, c)
			rules, err := m.buildRules(&orRule, subName, false)
			if err != nil {
				return nil, err
			}
			ret = append(ret, rules...)
		}
	}

	r := strings.Builder{}

	r.WriteString("SecRule ")

	if rule.Zones == nil {
		return ret, nil
	}

	for idx, zone := range rule.Zones {
		mappedZone, ok := zonesMap[zone]
		if !ok {
			return nil, fmt.Errorf("unknown zone '%s'", zone)
		}
		if len(rule.Variables) == 0 {
			r.WriteString(mappedZone)
		} else {
			for j, variable := range rule.Variables {
				if idx > 0 || j > 0 {
					r.WriteByte('|')
				}
				r.WriteString(fmt.Sprintf("%s:%s", mappedZone, variable))
			}
		}
	}
	r.WriteByte(' ')

	if rule.Match.Type != "" {
		if match, ok := matchMap[rule.Match.Type]; ok {
			r.WriteString(fmt.Sprintf(`"%s %s"`, match, rule.Match.Value))
		} else {
			return nil, fmt.Errorf("unknown match type '%s'", rule.Match.Type)
		}
	}

	//Should phase:2 be configurable?
	r.WriteString(fmt.Sprintf(` "id:%d,phase:2,deny,log,msg:'%s'`, m.generateRuleID(rule, waapRuleName), waapRuleName))

	if rule.Transform != nil {
		for _, transform := range rule.Transform {
			r.WriteByte(',')
			if mappedTransform, ok := transformMap[transform]; ok {
				r.WriteString(mappedTransform)
			} else {
				return nil, fmt.Errorf("unknown transform '%s'", transform)
			}
		}
	}

	if and {
		r.WriteString(",chain")
	}

	r.WriteByte('"')

	ret = append(ret, r.String())
	return ret, nil
}