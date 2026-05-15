package requests

type AddMaintenanceFormRequest struct {
	Name string `schema:"name"`
	Date string `schema:"date"`
	KM   int    `schema:"km"`
	Cost int    `schema:"cost"`
}
