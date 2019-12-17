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

	switch rg.Condition.(string) {
	case AND:
		for _, v := range rules {
			checker := rg.GetChecker(v.(map[string]interface{}))
			if ok := checker.Evaluate(dataset); !ok {
				return false
			}
		}
		return true

	case OR:
		for _, v := range rules {
			checker := rg.GetChecker(v.(map[string]interface{}))
			if ok := checker.Evaluate(dataset); ok {
				return true
			}
		}
		return false
	}

	return false
}

func (rg *RuleGroup) GetChecker(rule map[string]interface{}) Checker {
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
