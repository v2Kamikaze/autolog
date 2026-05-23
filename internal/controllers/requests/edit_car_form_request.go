package requests

type EditCarFormRequest struct {
	Title   string `schema:"title"`
	CarType string `schema:"car_type"`
	KM      int    `schema:"km"`
}
