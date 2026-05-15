package functions

import (
	"encoding/json"
	"html/template"
)

func ToJSON(data any) template.JS {
	jsonb, _ := json.Marshal(data)
	return template.JS(jsonb)
}
