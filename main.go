package main

import (
	"embed"
	"html/template"
	"log"
	nethttp "net/http"

	"github.com/v2code/autolog/internal/functions"
	"github.com/v2code/autolog/internal/http"
	"github.com/v2code/autolog/internal/persistence"
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
			"internal/templates/components/ui/logo.html",
			"internal/templates/components/modals/add_vehicle_modal.html",
			"internal/templates/components/modals/add_log_entry_modal.html",
			"internal/templates/components/modals/edit_vehicle_modal.html",
			"internal/templates/components/modals/edit_log_entry_modal.html",
			"internal/templates/components/modals/delete_vehicle_modal.html",
			"internal/templates/components/modals/delete_log_entry_modal.html",
		}

		files = append(files, components...)

		return template.Must(template.New("base_page.html").Funcs(functions.FuncMap).ParseFS(templateFiles, files...))
	}

	templates := map[string]*template.Template{
		"home": load(
			"internal/templates/pages/base_page.html",
			"internal/templates/pages/home_page.html",
			"internal/templates/partials/navigation_bar.html",
			"internal/templates/partials/vehicle_list.html",
			"internal/templates/partials/vehicle_log_entry_count_text.html",
			"internal/templates/partials/vehicle_list_item.html",
			"internal/templates/partials/log_entry_list.html",
			"internal/templates/partials/log_entry_list_item.html",
			"internal/templates/responses/add_vehicle_list_item_response.html",
			"internal/templates/responses/edit_vehicle_list_item_response.html",
			"internal/templates/responses/add_log_entry_list_item_response.html",
			"internal/templates/responses/edit_log_entry_list_item_response.html",
			"internal/templates/responses/delete_vehicle_list_item_response.html",
			"internal/templates/responses/delete_log_entry_list_item_response.html",
		),
		"vehicles": load(
			"internal/templates/pages/base_page.html",
			"internal/templates/pages/vehicles_page.html",
		),
	}

	vehicleService := vehicles.NewService(persistence.NewInMemoryVehiclePersistence())
	engine := ui.NewEngine(templates["home"])
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

	log.Fatal(nethttp.ListenAndServe(":8080", mux))
}
