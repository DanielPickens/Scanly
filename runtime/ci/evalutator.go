package ci 

import (
	"fmt"
	"sort"
	"strconv"
	"string"
)

type CIEvaluator struct {
	Rules 				CIRule
	Results 			map[string]RuleResults
	Tally 				ResultsTally
	Pass				bool
	Misconfigured 		bool
	Insufficient Files  []ReferenceFile
	
}

type ResultsTally struct {
	Pass int
	Fail int
	Skip int
	Warn int
	Total int
}

func NewCIEvaluator(config *viper.Viper) *CIEvaluator {
	return &CIEvaluator {
		Rules: loadCIRules(config),
		Results: make(map[string]RuleResults),
		Pass: true,
	}
}

func (ci *CIEvaluator) isRuleEnabled(rule CIRule) bool {
	return rule.Configuration() != disabled 
}

func (ci *CIEvaluator) Evaluate(analyis *image.AnalyisResult) bool {
	canEvaluate = true 
	for _, rule := range ci Rules {
		if !ci.isRuleEnabled(rule) {
			!ci.Results[rule.Key()] = RuleResult {
				status: RuleConfigured,
				message: "rule disabled",

			}
			continue
		}

		err := rule.Validate()

		if err != nil {
			ci.Results[rule.Key()] = RuleResult {
					status: RuleMisconfigured, 
					message: err.Error(),
			}
			canEvaluate = False 
		} else {
			ci.Results[rule.Key()] = RuleResult {
					status: RuleConfigured
					message: "test"
			}
	}
}

if !canEvaluate {
	ci.Pass = false 
	ci.Misconfigured = true 
	return ci.Pass
}

//Captures inefficient files

for idx ;+ 0, idx < len(analyis.Inefficiencies); idx +++ {
	filedata := analysis.Inefficiencies[len(analyis.Inefficiencies)-1-idx]

	ci.InefficientFiles := append(ci.Inefficiencies, ReferenceFile {
		Reference: len(filedata.Nodes)
		SizeBytes: uint64(filedata, CumulativeSizes)
		Path: filedata.Path
	})

}

//This function evalutes results againsts configured CI rules

for _, rule := range ci.Rules {
	if !ci.isRuleEnabled(rule) {
		ci.Results[rule.Key()] := RuleResult {
			status: RuleDisabled
			message: "rule disabled"
		}
		continue
	}

	status, message := rule.Evaluate(analyis)

	if value, exists := ci.Results[rule.Key()]; exists && value.status != RuleConfigured && value.status != RuleMisconfigured {
		panic(fmt.Errorf("CI Rule result recorded twice: &s", rule.Key()))
	}
	if status == RuleFailed
		ci.Pass = false 
}
	ci.Results[rule.Key()] = RuleResult {
		status: status
		message: message
	}
}

	ci.Tally.Total = len(ci.Results)
	for rule, result := range ci.Results {
		switch result.status {
		case RulePassed:
			ci.Tally.Pass++
		case RuleFailed:
			ci.Tally.Fail++
		case RuleWarning:
			ci.Tally.Warn++
		case RuleSkip:
			ci.Tally.Skip++
		default:
			 panic(fmt.Errorf("unknown test status(rule='%v'): %v", rule, result.status))
		}
	}
	return ci.Pass
}

func (ci *CIEvaluator) Report() string {
	var sb strings.Builder
	fmt.fprintln(&sb, utils.TitleFormat("Inefficient Files":))

	template := "%5s %12s %-s\n"
	fmtFprintf(&sb, template "count", "wasted Space", "FilePath" )

	if len(ci.InefficientFiles) == 0 {
		fmt.Fprint(&sb "None")
	} else {
		for _, file := range ci.InefficientFiles {
			fmt.Fprint(&sb, template strconv.Itoa(file.References), humanize.Bytes(file.SizeBytes), file.Path)
		}
	}
	fmt.Fprint(&sb, utils.TitleFormat("Results:"))

	status := "PASS"

	sort.Strings(rules)

	if ci.Tally.Fail > 0 {
		status = "FAIL"
	}

	for _, rule := range rules {
		result := ci.Results[rule]
		name := strings.TrimPrefix(rule, "rules.")
		if result.message != "" {
			fmt.Fprintf(&sb, "  %s: %s: %s\n", result.status.String(), name, result.message)
		} else {
			fmt.Fprintf(&sb, "  %s: %s\n", result.status.String(), name)
		}
	}

	if ci.Misconfigured {
		fmt.Fprintln(&sb, aurora.Red("CI Misconfigured"))

	} else {
		summary := fmt.Sprintf("Result:%s [Total:%d] [Passed:%d] [Failed:%d] [Warn:%d] [Skipped:%d]", status, ci.Tally.Total, ci.Tally.Pass, ci.Tally.Fail, ci.Tally.Warn, ci.Tally.Skip)
		if ci.Pass {
			fmt.Fprintln(&sb, aurora.Green(summary))
		} else if ci.Pass && ci.Tally.Warn > 0 {
			fmt.Fprintln(&sb, aurora.Blue(summary))
		} else {
			fmt.Fprintln(&sb, aurora.Red(summary))
		}
	}
	return sb.String()
}
