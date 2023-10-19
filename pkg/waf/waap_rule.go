package waf

import (
	"fmt"
	"strings"
	"time"
)

type VPatchRule struct {
	//Those 2 together represent something like ARGS.foo
	//If only target is set, it's used for variables that are not a collection (REQUEST_METHOD, etc)
	Target   string `yaml:"target"`
	Variable string `yaml:"var"`

	Match     string       `yaml:"match"`           //@rx
	Equals    string       `yaml:"equals"`          //@eq
	Transform string       `yaml:"transform"`       //t:lowercase, t:uppercase, etc
	Detect    string       `yaml:"detect"`          //@detectXSS, @detectSQLi, etc
	Logic     string       `yaml:"logic,omitempty"` // "AND", "OR", or empty if not applicable
	SubRules  []VPatchRule `yaml:"sub_rules,omitempty"`
}

func (v *VPatchRule) String() string {
	return strings.Trim(v.constructRule(0), "\n")
}

func countTotalRules(rules []VPatchRule) int {
	count := 0
	for _, rule := range rules {
		count++
		if rule.Logic == "AND" {
			count += countTotalRules(rule.SubRules)
		}
	}
	return count
}

func (v *VPatchRule) constructRule(depth int) string {
	var result string
	result = v.singleRuleString()

	if len(v.SubRules) == 0 {
		return result + "\n"
	}

	switch v.Logic {
	case "AND":
		result = strings.TrimSuffix(result, `"`) + `,chain"` + "\n"
		for _, subRule := range v.SubRules {
			result += subRule.constructRule(depth + 1)
		}
	case "OR":
		skips := countTotalRules(v.SubRules) - 1
		if depth == 0 {
			skips++
		}
		result = strings.TrimSuffix(result, `"`) + fmt.Sprintf(`,skip:%d"`+"\n", skips)
		for _, subRule := range v.SubRules {
			skips--
			if skips > 0 {
				result += strings.TrimSuffix(subRule.singleRuleString(), `"`) + fmt.Sprintf(`,skip:%d"`+"\n", skips)
			} else {
				result += subRule.singleRuleString() + "\n"
			}
		}
	}
	return result
}

func (v *VPatchRule) singleRuleString() string {
	var operator string
	var ruleStr string

	if v.Match != "" {
		operator = fmt.Sprintf("@rx %s", v.Match)
	} else if v.Equals != "" {
		operator = fmt.Sprintf("@eq %s", v.Equals)
	} else {
		return ""
	}

	if v.Variable != "" {
		ruleStr = fmt.Sprintf(`SecRule %s:%s "%s"`, v.Target, v.Variable, operator)
	} else {
		ruleStr = fmt.Sprintf(`SecRule %s "%s"`, v.Target, operator)
	}

	//FIXME: phase2 should probably not be hardcoded
	//Find a better way than using time.Now().UnixMilli() to generate a unique ID
	actions := fmt.Sprintf(` "id:%d,deny,log,phase:2`, time.Now().UnixNano())

	// Handle transformation
	if v.Transform != "" {
		actions = actions + fmt.Sprintf(",t:%s", v.Transform)
	}
	actions = actions + `"`
	ruleStr = ruleStr + actions

	return ruleStr
}