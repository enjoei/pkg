# Query Builder
This package is a Golang Rule Evaluator for [jQuery QueryBuilder](https://querybuilder.js.org/index.html)

You will use the Evaluator to check if a dataset matches the JSON rules produced by the jQuery QueryBuilder plugin.

Dataset is a map of the type `map[string]interface{}`

You can access nested fields with a dot notation. e.g. fields.Question to access {'fields' => {'Question' => 'Answer'}}

## Example usage

```golang

var rulesetJSON = `{
  "condition": "AND",
  "rules": [
    {"id": "name", "field": "name", "type": "string", "input": "text", "operator": "begins_with", "value": "shirt"},
    {"id": "category", "field": "category", "type": "string", "input": "text", "operator": "equal", "value": "woman"},
    {"condition": "OR","rules": [
      {"id": "price", "field": "price", "type": "double", "input": "text", "operator": "greater","value": 10.50},
      {"id": "published_at", "field": "published_at", "type": "date", "input": "text", "operator": "less", "value": "2019-12-25"}
      ]}
    ]}`

var datasetJSON = `{"name": "short sleeve shirt", "category": "woman", "price":  33.0, "published_at": "2020-1-1"}`

func main(){
  var ruleset map[string]interface{}
  var dataset map[string]interface{}
  json.Unmarshal([]byte(rulesetJSON), &ruleset)
  json.Unmarshal([]byte(datasetJSON), &dataset)
  
  qb := querybuilder.New(ruleset)
  if qb.Match(dataset) {
    fmt.Println("√ dataset is valid!")
  }
}

```

## Operators
All the default operators have been implemented, see in `/operators` folder

You can easily add custom operators, see:

```golang
import "github.com/enjoei/pkg/querybuilder/operator"

func init() {
	AddOperator(EqualsFive)
}

// Equal
var EqualsFive = &Operator{
	Name: "equals_five",
	Evaluate: func(input, value interface{}) bool {
		return input(int) == 5
	},
}
```
The Evaluate function shoult receive a input and value params, even if a value is not required

## Contributing
- Fork it
- Create your feature branch (`git checkout -b my-new-feature`)
- Commit your changes (`git commit -am 'Add some feature'`)
- Push to the branch (`git push origin my-new-feature`)
- Create new Pull Request
