package functions

import "html/template"

var FuncMap = template.FuncMap{
	"Dict":        Dict,
	"FmtDate":     FmtDate,
	"FmtCurrency": FmtCurrency,
	"ToDate":      ToDate,
}
