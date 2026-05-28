package http

type CreateVehicleForm struct {
	Title       string `schema:"title"`
	VehicleType string `schema:"vehicle_type"`
	KM          int    `schema:"km"`
}

type EditVehicleForm struct {
	Title       string `schema:"title"`
	VehicleType string `schema:"vehicle_type"`
	KM          int    `schema:"km"`
}

type AddLogEntryForm struct {
	Name      string `schema:"name"`
	EntryType string `schema:"entry_type"`
	Date      string `schema:"date"`
	KM        int    `schema:"km"`
	Cost      int    `schema:"cost"`
}

type EditLogEntryForm struct {
	Name      string `schema:"name"`
	EntryType string `schema:"entry_type"`
	Date      string `schema:"date"`
	KM        int    `schema:"km"`
	Cost      int    `schema:"cost"`
}
