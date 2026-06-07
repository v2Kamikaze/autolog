package http

type CreateVehicleForm struct {
	Brand        string `schema:"brand"`
	Model        string `schema:"model"`
	Year         int    `schema:"year"`
	Version      string `schema:"version"`
	Engine       string `schema:"engine"`
	Transmission string `schema:"transmission"`
	Fuel         string `schema:"fuel"`
	KM           int    `schema:"km"`
}

type EditVehicleForm struct {
	Brand        string `schema:"brand"`
	Model        string `schema:"model"`
	Year         int    `schema:"year"`
	Version      string `schema:"version"`
	Engine       string `schema:"engine"`
	Transmission string `schema:"transmission"`
	Fuel         string `schema:"fuel"`
	KM           int    `schema:"km"`
}

type AddLogEntryForm struct {
	Name      string `schema:"name"`
	Notes     string `schema:"notes"`
	EntryType string `schema:"entry_type"`
	Date      string `schema:"date"`
	KM        int    `schema:"km"`
	Cost      int    `schema:"cost"`
}

type EditLogEntryForm struct {
	Name      string `schema:"name"`
	Notes     string `schema:"notes"`
	EntryType string `schema:"entry_type"`
	Date      string `schema:"date"`
	KM        int    `schema:"km"`
	Cost      int    `schema:"cost"`
}
