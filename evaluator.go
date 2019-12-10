package querybuilder

type Evaluator struct {
	Ruleset map[string]interface{}
}

func (e *Evaluator) Match(dataset map[string]interface{}) bool {
	rg := RuleGroup{Condition: e.Ruleset["condition"], Rules: e.Ruleset["rules"]}
	return rg.Evaluate(dataset)
}

func New(ruleset map[string]interface{}) *Evaluator {
	return &Evaluator{Ruleset: ruleset}
}
