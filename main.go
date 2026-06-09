package main

import (
	"embed"
	"html/template"
	"log"
	nethttp "net/http"

	"github.com/v2code/autolog/internal/database"
	"github.com/v2code/autolog/internal/functions"
	"github.com/v2code/autolog/internal/http"
	"github.com/v2code/autolog/internal/ui"
	"github.com/v2code/autolog/internal/vehicles"
)

//go:embed internal/templates
var templateFiles embed.FS

//go:embed static/**
var staticFiles embed.FS

func main() {
	load := func(files ...string) *template.Template {
		components := []string{

			"internal/templates/components/ui/icons.html",
			"internal/templates/components/layout/navigation_bar.html",
			"internal/templates/components/modals/add_vehicle_modal.html",
			"internal/templates/components/modals/add_log_entry_modal.html",
			"internal/templates/components/modals/edit_vehicle_modal.html",
			"internal/templates/components/modals/edit_log_entry_modal.html",
			"internal/templates/components/modals/delete_vehicle_modal.html",
			"internal/templates/components/modals/delete_log_entry_modal.html",
			"internal/templates/components/lists/log_entry_amount_text.html",
			"internal/templates/components/lists/vehicles_list.html",
			"internal/templates/components/lists/vehicles_list_empty.html",
			"internal/templates/components/lists/vehicles_list_item.html",
			"internal/templates/components/lists/log_list.html",
			"internal/templates/components/lists/log_list_item.html",
			"internal/templates/components/lists/log_list_empty.html",
			"internal/templates/components/lists/log_entry_count_text.html",
			"internal/templates/components/streams/add_vehicle_stream.html",
			"internal/templates/components/streams/edit_vehicle_stream.html",
			"internal/templates/components/streams/delete_vehicle_stream.html",
			"internal/templates/components/streams/add_log_stream.html",
			"internal/templates/components/streams/edit_log_stream.html",
			"internal/templates/components/streams/delete_log_stream.html",
		}

		files = append(files, components...)

		return template.Must(template.New("base_page.html").Funcs(functions.FuncMap).ParseFS(templateFiles, files...))
	}

	homeTemplate := load(
		"internal/templates/components/layout/base_page.html",
		"internal/templates/pages/home_page.html",
	)

	conn := database.OpenDatabase()
	defer conn.Close()

	if err := database.Migrate(conn); err != nil {
		log.Fatal(err)
	}

	db := database.NewDatabase(conn)

	vehicleStore := vehicles.NewStore(db)

	vehicleService := vehicles.NewService(vehicleStore)
	engine := ui.NewEngine(homeTemplate)
	handler := http.NewVehicleHandler(vehicleService, engine)

	mux := nethttp.NewServeMux()
	mux.Handle("GET /static/", nethttp.FileServer(nethttp.FS(staticFiles)))
	mux.HandleFunc("GET /", handler.Home)
	mux.HandleFunc("POST /vehicles", handler.CreateVehicle)
	mux.HandleFunc("POST /vehicles/{id}/log-entries", handler.AddLogEntry)
	mux.HandleFunc("PUT /vehicles/{vehicleId}/log-entries/{logEntryId}", handler.EditLogEntry)
	mux.HandleFunc("PUT /vehicles/{id}", handler.EditVehicle)
	mux.HandleFunc("DELETE /vehicles/{id}", handler.DeleteVehicle)
	mux.HandleFunc("DELETE /vehicles/{vehicleId}/log-entries/{logEntryId}", handler.DeleteLogEntry)

	log.Fatal(nethttp.ListenAndServe("0.0.0.0:8080", http.GzipMiddleware(mux)))
}
