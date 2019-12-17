package querybuilder

const (
	AND = "AND"
	OR  = "OR"
)

type Checker interface {
	Evaluate(dataset map[string]interface{}) bool
}

type RuleGroup struct {
	Condition interface{}
	Rules     interface{}
}

func (rg *RuleGroup) Evaluate(dataset map[string]interface{}) bool {
	rules := rg.Rules.([]interface{})
	rulesSize := len(rules)
	res := make(chan bool)

	switch rg.Condition.(string) {
	case AND:
		rg.evaluateRules(res, rules, dataset)

		for i := 0; i < rulesSize; i++ {
			if !<-res {
				return false
			}
		}
		return true
	case OR:
		rg.evaluateRules(res, rules, dataset)

		for i := 0; i < rulesSize; i++ {
			if <-res {
				return true
			}
		}
		return false
	default:
		return false
	}
}

func (rg *RuleGroup) getChecker(rule map[string]interface{}) Checker {
	if _, ok := rule["rules"]; ok {
		return &RuleGroup{Condition: rule["condition"], Rules: rule["rules"]}
	}

	return &Rule{
		ID:       rule["id"].(string),
		Field:    rule["field"].(string),
		Type:     rule["type"].(string),
		Input:    rule["input"].(string),
		Operator: rule["operator"].(string),
		Value:    rule["value"],
	}
}

func (rg *RuleGroup) evaluateRules(res chan<- bool, rules []interface{}, dataset map[string]interface{}) {
	for _, v := range rules {
		go func(r map[string]interface{}) {
			checker := rg.getChecker(r)
			res <- checker.Evaluate(dataset)
		}(v.(map[string]interface{}))
	}
}
