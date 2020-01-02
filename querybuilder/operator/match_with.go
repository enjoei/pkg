package operator

import (
	"reflect"
	"regexp"
	"strings"
)

func init() {
	AddOperator(MatchWith)
}

var MatchWith = &Operator{
	Name: "match_with",
	Evaluate: func(input, value interface{}) bool {
		rv := reflect.ValueOf(value)
		if rv.Kind() != reflect.String {
			return false
		}

		v := value.(string)
		in := input.(string)

		if strings.HasPrefix(v, "/") && strings.HasSuffix(v, "/") {
			v = v[1 : len(v)-1]                   // remove slashes
			vr, err := regexp.Compile("(?i)" + v) // add ignore case option to regex and compile it
			if err != nil {
				return false
			}

			if vr.MatchString(in) {
				return true
			}
		}

		return false
	},
}
