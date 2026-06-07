package functions

import (
	"encoding/json"
	"fmt"
	"html/template"
	"strings"
	"time"
)

func Dict(values ...any) map[string]any {
	if len(values)%2 != 0 {
		panic("invalid dict call: must have even number of arguments")
	}
	dict := make(map[string]any, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			panic("dict keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict
}

func FmtDate(date time.Time) string {
	return date.Format("02/01/2006")
}

func ToDate(date time.Time) string {
	return date.Format("2006-01-02")
}

func FmtCurrency(cents int) string {
	return strings.ReplaceAll(fmt.Sprintf("R$ %0.2f", float64(cents)/100), ".", ",")
}

func JsQuote(s string) string {
	b, _ := json.Marshal(s)
	return string(b)
}

var FuncMap = template.FuncMap{
	"Dict":        Dict,
	"FmtDate":     FmtDate,
	"FmtCurrency": FmtCurrency,
	"ToDate":      ToDate,
	"JsQuote":     JsQuote,
}
