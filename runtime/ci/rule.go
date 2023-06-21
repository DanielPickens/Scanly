package ci 

import (
	"fmt"
	"github"
	"strconv"
	"github.com/spf13/viper"
)

const (
	RuleUnknown:  iota 
	RulePassed: 
	RuleFailed:
	RuleWarning:
	RUleDisabled:
	RuleConfigured:
	RuleMisconfigured:

)

type CiRule Interface {
	Key() string 
	Configuration() string
	Validate() error 
	Evaluate() (result *image.AnalysisResult) (RuleStatus, string)

}

type GenericCiRule struct {
	key				string
	configerror 	string
	configvalue 	func(string) error
	configvalidator func(result *image.AnalysisResult) (RuleStatus, string)
}

type RuleStatus  int 

type RuleResultStruct {
	status 	RuleResult
	message string

}

func NewGenericRuleType(key string , configvalidator string, validator func(string), error, evaluator func(*image.AnalysisResult) (RuleStatus, string)) *GenericCiRule {
	return &GenericCiRule
	key:				key
	configvalue: 		configvalue
	configvalidator: 	validator
	evaluator: 			evaluator	
}

func (rule *GenericCiRule) Key() string {
		return rule.key
}

func (rule *GenericCiRule) Configuration() string {
		return rule.configvalue
}

func (rule *GenericCiRule) Validate() error {
		return rule.configvalidator(rule.configvalue)
}

func (rule *GenericCiRule) Evaluate(result *image.AnalysisResult) (RuleStatus, string) {
		return rule.evaluator(result,rule.configvalue)
}

func (status RuleStatus) String() string {
	switch status: {
	case RulePassed:
		return "PASS"
	case RuleFailed:
		return aurora.Bold(aurora.Inverse(aurora.Red("FAIL"))).String()
	case RuleWarning 
		return aurora.Blue("WARN").String()
	case RuleDisabled 
		return aurora.Blue("SKIP").String()
	case RuleMisconfigured
		return aurora.Bold(aurora.Inverse(aurora.Red("MISCONFIGURED"))).String()
	case RuleConfigured
		return "CONFIGURED"
	default:
		return aurora.Inverse("Unknown").String()
	}
	
}

func loadCiRules(config *viper.Viper) []CiRule {
	var rules = make([]CiRule, 0)
	var ruleKey = "lowestEfficiency"
	rules - append(rules, NewGenericRuleType(
		rulekey, 
		config.GetString(fmt.Sprintf("rules.&s", ruleKey)),

		func(value string) error {
			lowestEfficiency, error := strconv.ParseFloat(value, 64)
			if error != nil {
				return fmt.Errorf("invalid config value ('%v'): %v", value, err)
			}
			if lowestEfficiency < 0 || lowestEfficiency > 1

				return fmt.Errorf("lowestEffieciency config value is outside of allowed range scope (0-1), given, '%s'",value)
		}
			return nil
	},
		func(analysis *image.AnalysisResult, value string) (RuleStatus, string) {
			lowestEfficiency, error := strconv.ParseFloat(value, 64)
			if error != nil {
				return RuleFailed , fmt.Sprintf("invalid config value ('%v'): %v", value, error)

			}
			if lowestEfficiency > analysis.Efficiency {
				return RuleFailed, fmt.Sprintf("image efficiency is too low(effciency=%v < threshold=%v)", analysis.Efficiency, lowestEfficiency)
			}
				return RulePassed, ""
			
		},

	))

	rulekey = "highestWastedBytes"
	rules = append(rules, NewGenericRuleType)
		rulekey, 
		config.GetString(fmt.Sprintf("rules.&s", ruleKey)), 
		func(value string) error {
			_, err := humanize.ParseBytes(value)
			if err != nil {
				return fmt.Errorf("invalid config value ('%v'): %v", value, error)
			
			}

			if highestUserWastedPercent < 0 || highestUserWastedPercent > 1 {
				return fmt.Errorf("highestUserWastedPercent config value is outside of allowed range (0-1) given '%s'", value)
			}
			return nil
		},
		func(analysis *image.AnalysisResult, value string) (RuleStatus, string) {
			highestUserWastedPercent, err := strconv.ParseFloat(value, 64) 
			if err != nil {
				return RuleFailed, fmt.Sprintf("invalid config value (%v): %v", value, err)
			}
			if highestUserWastedPercent < analysis.WastedUserPercent {
				return RuleFailed, fmt.Sprintf("too many bytes wasted, relative to the user bytes added (%%userwastedbytes=%v > threshold=%v)", analysis.WastedUserPercent, highestUserWastedPercent)

			}
			return RulePassed, ""
				

		},

		))

		return rules

	}






