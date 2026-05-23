package entities

import "time"

type LogEntryCategory string

const (
	LogEntryCategoryMaintenance LogEntryCategory = "maintenance"
	LogEntryCategoryExpense     LogEntryCategory = "expense"
)

func (c LogEntryCategory) Valid() bool {
	return c == LogEntryCategoryMaintenance || c == LogEntryCategoryExpense
}

func (c LogEntryCategory) Label() string {
	switch c {
	case LogEntryCategoryMaintenance:
		return "Manutenção"
	case LogEntryCategoryExpense:
		return "Despesa"
	default:
		return string(c)
	}
}

type LogEntry struct {
	ID    int
	CarID int
	Name  string
	Type  LogEntryCategory
	Date  time.Time
	KM    int
	Cost  int
}
