package requests

type CreateCarFormRequest struct {
	Title    string `schema:"title"`
	CarType  string `schema:"car_type"`
	KM       int    `schema:"km"`
}
