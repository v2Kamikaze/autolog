package entities

type Car struct {
	ID           int
	Title        string
	Type         string
	KM           int
	Maintenances []*Maintenance
}

func (c *Car) MaintenanceCount() int {
	if c.Maintenances == nil {
		return 0
	}
	return len(c.Maintenances)
}

func (c *Car) TotalCost() int {
	total := 0
	for _, maintenance := range c.Maintenances {
		total += maintenance.Cost
	}
	return total
}
