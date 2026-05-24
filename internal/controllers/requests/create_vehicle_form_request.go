package requests

type CreateVehicleFormRequest struct {
	Title    string `schema:"title"`
	VehicleType  string `schema:"vehicle_type"`
	KM       int    `schema:"km"`
}
