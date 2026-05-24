package requests

type EditVehicleFormRequest struct {
	Title   string `schema:"title"`
	VehicleType string `schema:"vehicle_type"`
	KM      int    `schema:"km"`
}
