package requests

type EditCarFormRequest struct {
	Title string `schema:"title"`
	KM    int    `schema:"km"`
}
