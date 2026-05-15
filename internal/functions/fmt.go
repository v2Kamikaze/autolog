package functions

import (
	"fmt"
	"strings"
	"time"
)

func FmtDate(date time.Time) string {
	return date.Format("02/01/2006")
}

func ToDate(date time.Time) string {
	return date.Format("2006-01-02")
}

func FmtCurrency(cents int) string {
	return strings.ReplaceAll(fmt.Sprintf("R$ %0.2f", float64(cents)/100), ".", ",")
}
