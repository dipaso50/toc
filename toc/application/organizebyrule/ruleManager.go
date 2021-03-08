package organizebyrule

import "tocV2/toc/domain"

type RuleManager struct {
	allRules []domain.Rule
}

func NewRuleManager() *RuleManager {
	r := RuleManager{}
	r.allRules = make([]domain.Rule, 0)
	return &r
}

func (rm *RuleManager) AddRule(fileMatchRegularExpresion, destination string) {
	rm.allRules = append(rm.allRules, domain.NewRule(fileMatchRegularExpresion, destination))
}

func (rm *RuleManager) FileMatch(filename string) (bool, string) {
	for _, r := range rm.allRules {
		if r.RegexpFilter.MatchString(filename) {
			return true, r.DestinationFolder
		}
	}
	return false, ""
}
