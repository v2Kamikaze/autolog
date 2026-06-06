package vehicles

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
	ID        int
	VehicleID int
	Name      string
	Type      LogEntryCategory
	Date      time.Time
	KM        int
	Cost      int
}

type Vehicle struct {
	ID         int
	Title      string
	Type       string
	KM         int
	LogEntries []*LogEntry
}

func (v *Vehicle) LogEntryCount() int {
	if v.LogEntries == nil {
		return 0
	}
	return len(v.LogEntries)
}

func (v *Vehicle) TotalCost() int {
	total := 0
	for _, entry := range v.LogEntries {
		total += entry.Cost
	}
	return total
}
