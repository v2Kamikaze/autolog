package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"

	"github.com/v2code/autolog/internal/controllers"
	"github.com/v2code/autolog/internal/functions"
	"github.com/v2code/autolog/internal/persistence"
	"github.com/v2code/autolog/internal/usecase/vehicleusecases"
)

//go:embed internal/templates
var templateFiles embed.FS

//go:embed static/**
var staticFiles embed.FS

func main() {
	load := func(files ...string) *template.Template {
		components := []string{

			"internal/templates/components/ui/icons.html",
			"internal/templates/components/ui/kanji_bg.html",
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

	vehiclePersistence := persistence.NewInMemoryVehiclePersistence()
	listVehiclesUseCase := vehicleusecases.NewListVehiclesUseCase(vehiclePersistence)
	createVehicleUseCase := vehicleusecases.NewCreateVehicleUseCase(vehiclePersistence)
	addLogEntryUseCase := vehicleusecases.NewAddLogEntryUseCase(vehiclePersistence)
	editLogEntryUseCase := vehicleusecases.NewEditLogEntryUseCase(vehiclePersistence)
	editVehicleUseCase := vehicleusecases.NewEditVehicleUseCase(vehiclePersistence)
	deleteVehicleUseCase := vehicleusecases.NewDeleteVehicleUseCase(vehiclePersistence)
	deleteLogEntryUseCase := vehicleusecases.NewDeleteLogEntryUseCase(vehiclePersistence)
	homeHandler := controllers.NewHomeHandler(
		listVehiclesUseCase,
		createVehicleUseCase,
		addLogEntryUseCase,
		editLogEntryUseCase,
		editVehicleUseCase,
		deleteVehicleUseCase,
		deleteLogEntryUseCase,
		templates["home"],
	)

	mux := http.NewServeMux()
	mux.Handle("GET /static/", http.FileServer(http.FS(staticFiles)))
	mux.HandleFunc("GET /", homeHandler.Home)
	mux.HandleFunc("POST /vehicles", homeHandler.CreateVehicle)
	mux.HandleFunc("POST /vehicles/{id}/log-entries", homeHandler.AddLogEntry)
	mux.HandleFunc("PUT /vehicles/{vehicleId}/log-entries/{logEntryId}", homeHandler.EditLogEntry)
	mux.HandleFunc("PUT /vehicles/{id}", homeHandler.EditVehicle)
	mux.HandleFunc("DELETE /vehicles/{id}", homeHandler.DeleteVehicle)
	mux.HandleFunc("DELETE /vehicles/{vehicleId}/log-entries/{logEntryId}", homeHandler.DeleteLogEntry)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
