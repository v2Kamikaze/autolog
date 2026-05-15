package entities

type Car struct {
	ID           int
	Title        string
	KM           int
	Maintenances []*Maintenance
}

func (c *Car) TotalCost() int {
	total := 0
	for _, maintenance := range c.Maintenances {
		total += maintenance.Cost
	}
	return total
}
