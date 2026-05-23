package controllers

import (
	"errors"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/schema"
	"github.com/v2code/autolog/internal/controllers/requests"
	"github.com/v2code/autolog/internal/persistence"
	"github.com/v2code/autolog/internal/usecase/carusecases"
)

const TurboStreamContentType = "text/vnd.turbo-stream.html"

type HomeHandler struct {
	listCarsUseCase          *carusecases.ListCarsUseCase
	createCarUseCase         *carusecases.CreateCarUseCase
	addMaintenanceUseCase    *carusecases.AddMaintenanceUseCase
	editMaintenanceUseCase   *carusecases.EditMaintenanceUseCase
	editCarUseCase           *carusecases.EditCarUseCase
	deleteCarUseCase         *carusecases.DeleteCarUseCase
	deleteMaintenanceUseCase *carusecases.DeleteMaintenanceUseCase
	formDecoder              *schema.Decoder
	tmpl                     *template.Template
}

func NewHomeHandler(
	listCarsUseCase *carusecases.ListCarsUseCase,
	createCarUseCase *carusecases.CreateCarUseCase,
	addMaintenanceUseCase *carusecases.AddMaintenanceUseCase,
	editMaintenanceUseCase *carusecases.EditMaintenanceUseCase,
	editCarUseCase *carusecases.EditCarUseCase,
	deleteCarUseCase *carusecases.DeleteCarUseCase,
	deleteMaintenanceUseCase *carusecases.DeleteMaintenanceUseCase,
	tmpl *template.Template,
) *HomeHandler {
	formDecoder := schema.NewDecoder()
	formDecoder.IgnoreUnknownKeys(true)

	return &HomeHandler{
		listCarsUseCase:          listCarsUseCase,
		createCarUseCase:         createCarUseCase,
		addMaintenanceUseCase:    addMaintenanceUseCase,
		editMaintenanceUseCase:   editMaintenanceUseCase,
		editCarUseCase:           editCarUseCase,
		deleteCarUseCase:         deleteCarUseCase,
		deleteMaintenanceUseCase: deleteMaintenanceUseCase,
		formDecoder:              formDecoder,
		tmpl:                     tmpl,
	}
}

func (h *HomeHandler) Home(w http.ResponseWriter, r *http.Request) {
	output, err := h.listCarsUseCase.Execute(r.Context())
	if err != nil {
		http.Error(w, "failed to list cars", http.StatusInternalServerError)
		return
	}

	if err := h.tmpl.ExecuteTemplate(w, "base_page", output); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *HomeHandler) CreateCar(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "failed to parse form", http.StatusBadRequest)
		return
	}

	form := requests.CreateCarFormRequest{}
	if err := h.formDecoder.Decode(&form, r.PostForm); err != nil {
		http.Error(w, "invalid form data", http.StatusBadRequest)
		return
	}

	output, err := h.createCarUseCase.Execute(r.Context(), carusecases.CreateCarInput{
		Title: form.Title,
		Type:  form.CarType,
		KM:    form.KM,
	})
	if err != nil {
		http.Error(w, "failed to create car", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", TurboStreamContentType)

	if err := h.tmpl.ExecuteTemplate(w, "add_car_list_item_response", output); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *HomeHandler) AddMaintenance(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "failed to parse form", http.StatusBadRequest)
		return
	}

	carID := r.PathValue("id")

	carIDInt, err := strconv.Atoi(carID)
	if err != nil {
		http.Error(w, "invalid car id", http.StatusBadRequest)
		return
	}

	form := requests.AddMaintenanceFormRequest{}
	if err := h.formDecoder.Decode(&form, r.PostForm); err != nil {
		http.Error(w, "invalid form data", http.StatusBadRequest)
		return
	}

	dateTime, err := time.Parse("2006-01-02", form.Date)
	if err != nil {
		http.Error(w, "invalid date", http.StatusBadRequest)
		return
	}

	output, err := h.addMaintenanceUseCase.Execute(r.Context(), carusecases.AddMaintenanceInput{
		CarID: carIDInt,
		Name:  form.Name,
		Date:  dateTime,
		KM:    form.KM,
		Cost:  form.Cost,
	})
	if err != nil {
		if errors.Is(err, persistence.ErrCarNotFound) {
			http.Error(w, "car not found", http.StatusNotFound)
			return
		}
		http.Error(w, "failed to add maintenance", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", TurboStreamContentType)

	if err := h.tmpl.ExecuteTemplate(w, "add_maintenance_list_item_response", output); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *HomeHandler) EditMaintenance(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "failed to parse form", http.StatusBadRequest)
		return
	}

	carID := r.PathValue("carId")
	maintenanceID := r.PathValue("maintenanceId")

	carIDInt, err := strconv.Atoi(carID)
	if err != nil {
		http.Error(w, "invalid car id", http.StatusBadRequest)
		return
	}

	maintenanceIDInt, err := strconv.Atoi(maintenanceID)
	if err != nil {
		http.Error(w, "invalid maintenance id", http.StatusBadRequest)
		return
	}

	form := requests.EditMaintenanceFormRequest{}
	if err := h.formDecoder.Decode(&form, r.PostForm); err != nil {
		http.Error(w, "invalid form data", http.StatusBadRequest)
		return
	}

	dateTime, err := time.Parse("2006-01-02", form.Date)
	if err != nil {
		http.Error(w, "invalid date", http.StatusBadRequest)
		return
	}

	output, err := h.editMaintenanceUseCase.Execute(r.Context(), carusecases.EditMaintenanceInput{
		CarID:         carIDInt,
		MaintenanceID: maintenanceIDInt,
		Name:          form.Name,
		Date:          dateTime,
		KM:            form.KM,
		Cost:          form.Cost,
	})
	if err != nil {
		if errors.Is(err, persistence.ErrCarNotFound) {
			http.Error(w, "car not found", http.StatusNotFound)
			return
		}
		if errors.Is(err, persistence.ErrMaintenanceNotFound) {
			http.Error(w, "maintenance not found", http.StatusNotFound)
			return
		}
		http.Error(w, "failed to edit maintenance", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", TurboStreamContentType)

	if err := h.tmpl.ExecuteTemplate(w, "edit_maintenance_list_item_response", output); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *HomeHandler) EditCar(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "failed to parse form", http.StatusBadRequest)
		return
	}

	form := requests.EditCarFormRequest{}
	if err := h.formDecoder.Decode(&form, r.PostForm); err != nil {
		http.Error(w, "invalid form data", http.StatusBadRequest)
		return
	}

	carID := r.PathValue("id")

	carIDInt, err := strconv.Atoi(carID)
	if err != nil {
		http.Error(w, "invalid car id", http.StatusBadRequest)
		return
	}

	resp, err := h.editCarUseCase.Execute(r.Context(), carusecases.EditCarInput{
		ID:    carIDInt,
		Title: form.Title,
		Type:  form.CarType,
		KM:    form.KM,
	})
	if err != nil {
		if errors.Is(err, persistence.ErrCarNotFound) {
			http.Error(w, "car not found", http.StatusNotFound)
			return
		}
		http.Error(w, "failed to edit car", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", TurboStreamContentType)

	if err := h.tmpl.ExecuteTemplate(w, "edit_car_list_item_response", resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *HomeHandler) DeleteCar(w http.ResponseWriter, r *http.Request) {
	carID := r.PathValue("id")

	carIDInt, err := strconv.Atoi(carID)
	if err != nil {
		http.Error(w, "invalid car id", http.StatusBadRequest)
		return
	}

	resp, err := h.deleteCarUseCase.Execute(r.Context(), carusecases.DeleteCarInput{
		ID: carIDInt,
	})

	if err != nil {
		http.Error(w, "failed to delete car", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", TurboStreamContentType)

	if err := h.tmpl.ExecuteTemplate(w, "delete_car_list_item_response", resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *HomeHandler) DeleteMaintenance(w http.ResponseWriter, r *http.Request) {
	carID := r.PathValue("carId")
	maintenanceID := r.PathValue("maintenanceId")

	carIDInt, err := strconv.Atoi(carID)
	if err != nil {
		http.Error(w, "invalid car id", http.StatusBadRequest)
		return
	}

	maintenanceIDInt, err := strconv.Atoi(maintenanceID)
	if err != nil {
		http.Error(w, "invalid maintenance id", http.StatusBadRequest)
		return
	}

	resp, err := h.deleteMaintenanceUseCase.Execute(r.Context(), carusecases.DeleteMaintenanceInput{
		CarID:         carIDInt,
		MaintenanceID: maintenanceIDInt,
	})

	if err != nil {
		http.Error(w, "failed to delete maintenance", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", TurboStreamContentType)

	if err := h.tmpl.ExecuteTemplate(w, "delete_maintenance_list_item_response", resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
