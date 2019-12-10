package querybuilder

const (
	AND = "AND"
	OR  = "OR"
)

type Checker interface {
	Evaluate(dataset map[string]interface{}) bool
}

type RuleGroup struct {
	Condition string
	Rules     []map[string]interface{}
}

func (rg *RuleGroup) Evaluate(dataset map[string]interface{}) bool {
	switch rg.Condition {
	case AND:

	case OR:

	}
}

func (rg *RuleGroup) GetChecker() Checker {

}
