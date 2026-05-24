package entities

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
