package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"

	"github.com/v2code/autolog/internal/controllers"
	"github.com/v2code/autolog/internal/functions"
	"github.com/v2code/autolog/internal/persistence"
	"github.com/v2code/autolog/internal/usecase/carusecases"
)

//go:embed internal/templates/**/*.html
var templateFiles embed.FS

//go:embed static/**
var staticFiles embed.FS

func main() {
	load := func(files ...string) *template.Template {
		components := []string{
			"internal/templates/components/base_page.html",
			"internal/templates/components/icons.html",
			"internal/templates/components/add_car_modal.html",
			"internal/templates/components/add_maintenance_modal.html",
			"internal/templates/components/edit_car_modal.html",
			"internal/templates/components/edit_maintenance_modal.html",
			"internal/templates/components/delete_car_modal.html",
			"internal/templates/components/delete_maintenance_modal.html",
		}

		files = append(files, components...)

		return template.Must(template.New("base_page.html").Funcs(functions.FuncMap).ParseFS(templateFiles, files...))
	}

	templates := map[string]*template.Template{
		"home": load(
			"internal/templates/pages/home_page.html",
			"internal/templates/partials/car_list.html",
			"internal/templates/partials/car_list_item.html",
			"internal/templates/partials/maintenance_list.html",
			"internal/templates/partials/maintenance_list_item.html",
			"internal/templates/responses/add_car_list_item_response.html",
			"internal/templates/responses/edit_car_list_item_response.html",
			"internal/templates/responses/add_maintenance_list_item_response.html",
			"internal/templates/responses/edit_maintenance_list_item_response.html",
			"internal/templates/responses/delete_car_list_item_response.html",
			"internal/templates/responses/delete_maintenance_list_item_response.html",
		),
		"cars": load(
			"internal/templates/pages/cars_page.html",
		),
	}

	carPersistence := persistence.NewInMemoryCarPersistence()
	listCarsUseCase := carusecases.NewListCarsUseCase(carPersistence)
	createCarUseCase := carusecases.NewCreateCarUseCase(carPersistence)
	addMaintenanceUseCase := carusecases.NewAddMaintenanceUseCase(carPersistence)
	editMaintenanceUseCase := carusecases.NewEditMaintenanceUseCase(carPersistence)
	editCarUseCase := carusecases.NewEditCarUseCase(carPersistence)
	deleteCarUseCase := carusecases.NewDeleteCarUseCase(carPersistence)
	deleteMaintenanceUseCase := carusecases.NewDeleteMaintenanceUseCase(carPersistence)
	homeHandler := controllers.NewHomeHandler(
		listCarsUseCase,
		createCarUseCase,
		addMaintenanceUseCase,
		editMaintenanceUseCase,
		editCarUseCase,
		deleteCarUseCase,
		deleteMaintenanceUseCase,
		templates["home"],
	)

	mux := http.NewServeMux()
	mux.Handle("GET /static/", http.FileServer(http.FS(staticFiles)))
	mux.HandleFunc("GET /", homeHandler.Home)
	mux.HandleFunc("POST /cars", homeHandler.CreateCar)
	mux.HandleFunc("POST /cars/{id}/maintenances", homeHandler.AddMaintenance)
	mux.HandleFunc("PUT /cars/{carId}/maintenances/{maintenanceId}", homeHandler.EditMaintenance)
	mux.HandleFunc("PUT /cars/{id}", homeHandler.EditCar)
	mux.HandleFunc("DELETE /cars/{id}", homeHandler.DeleteCar)
	mux.HandleFunc("DELETE /cars/{carId}/maintenances/{maintenanceId}", homeHandler.DeleteMaintenance)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
