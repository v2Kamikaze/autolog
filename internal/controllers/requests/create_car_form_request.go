package requests

type CreateCarFormRequest struct {
	Title string `schema:"title"`
	KM    int    `schema:"km"`
}
