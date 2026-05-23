package entities

type Car struct {
	ID         int
	Title      string
	Type       string
	KM         int
	LogEntries []*LogEntry
}

func (c *Car) LogEntryCount() int {
	if c.LogEntries == nil {
		return 0
	}
	return len(c.LogEntries)
}

func (c *Car) TotalCost() int {
	total := 0
	for _, entry := range c.LogEntries {
		total += entry.Cost
	}
	return total
}
