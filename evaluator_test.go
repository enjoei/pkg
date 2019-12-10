package querybuilder

import (
	"encoding/json"
	"testing"
)

var rulestr = `{
  "condition": "AND",
  "rules": [{
    "id": "Decimal_Question",
    "field": "Decimal_Question",
    "type": "double",
    "input": "text",
    "operator": "equal",
    "value": "1.2"
  }, {
    "condition": "AND",
    "rules": [{
      "id": "Date_Question",
      "field": "Date_Question",
      "type": "date",
      "input": "text",
      "operator": "greater",
      "value": "2016-07-19"
    }, {
      "id": "Yes_No_Question",
      "field": "Yes_No_Question",
      "type": "boolean",
      "input": "select",
      "operator": "equal",
      "value": "true"
    }]
  }]
}`

func TestNew(t *testing.T) {
	//
	//
	// err := json.Unmarshal([]byte(rulestr), &rg)
	// t.Log(err)
	// t.Log(rg)
}
