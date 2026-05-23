package requests

type EditLogEntryFormRequest struct {
	Name      string `schema:"name"`
	EntryType string `schema:"entry_type"`
	Date      string `schema:"date"`
	KM        int    `schema:"km"`
	Cost      int    `schema:"cost"`
}
