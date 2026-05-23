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
	"github.com/v2code/autolog/internal/persistence/entities"
	"github.com/v2code/autolog/internal/usecase/carusecases"
)

const TurboStreamContentType = "text/vnd.turbo-stream.html"

type HomeHandler struct {
	listCarsUseCase       *carusecases.ListCarsUseCase
	createCarUseCase      *carusecases.CreateCarUseCase
	addLogEntryUseCase    *carusecases.AddLogEntryUseCase
	editLogEntryUseCase   *carusecases.EditLogEntryUseCase
	editCarUseCase        *carusecases.EditCarUseCase
	deleteCarUseCase      *carusecases.DeleteCarUseCase
	deleteLogEntryUseCase *carusecases.DeleteLogEntryUseCase
	formDecoder           *schema.Decoder
	tmpl                  *template.Template
}

func NewHomeHandler(
	listCarsUseCase *carusecases.ListCarsUseCase,
	createCarUseCase *carusecases.CreateCarUseCase,
	addLogEntryUseCase *carusecases.AddLogEntryUseCase,
	editLogEntryUseCase *carusecases.EditLogEntryUseCase,
	editCarUseCase *carusecases.EditCarUseCase,
	deleteCarUseCase *carusecases.DeleteCarUseCase,
	deleteLogEntryUseCase *carusecases.DeleteLogEntryUseCase,
	tmpl *template.Template,
) *HomeHandler {
	formDecoder := schema.NewDecoder()
	formDecoder.IgnoreUnknownKeys(true)

	return &HomeHandler{
		listCarsUseCase:       listCarsUseCase,
		createCarUseCase:      createCarUseCase,
		addLogEntryUseCase:    addLogEntryUseCase,
		editLogEntryUseCase:   editLogEntryUseCase,
		editCarUseCase:        editCarUseCase,
		deleteCarUseCase:      deleteCarUseCase,
		deleteLogEntryUseCase: deleteLogEntryUseCase,
		formDecoder:           formDecoder,
		tmpl:                  tmpl,
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

func (h *HomeHandler) AddLogEntry(w http.ResponseWriter, r *http.Request) {
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

	form := requests.AddLogEntryFormRequest{}
	if err := h.formDecoder.Decode(&form, r.PostForm); err != nil {
		http.Error(w, "invalid form data", http.StatusBadRequest)
		return
	}

	dateTime, err := time.Parse("2006-01-02", form.Date)
	if err != nil {
		http.Error(w, "invalid date", http.StatusBadRequest)
		return
	}

	output, err := h.addLogEntryUseCase.Execute(r.Context(), carusecases.AddLogEntryInput{
		CarID: carIDInt,
		Name:  form.Name,
		Type:  entities.LogEntryCategory(form.EntryType),
		Date:  dateTime,
		KM:    form.KM,
		Cost:  form.Cost,
	})
	if err != nil {
		if errors.Is(err, persistence.ErrCarNotFound) {
			http.Error(w, "car not found", http.StatusNotFound)
			return
		}
		if errors.Is(err, persistence.ErrInvalidLogEntryType) {
			http.Error(w, "invalid entry type", http.StatusBadRequest)
			return
		}
		http.Error(w, "failed to add log entry", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", TurboStreamContentType)

	if err := h.tmpl.ExecuteTemplate(w, "add_log_entry_list_item_response", output); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *HomeHandler) EditLogEntry(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "failed to parse form", http.StatusBadRequest)
		return
	}

	carID := r.PathValue("carId")
	logEntryID := r.PathValue("logEntryId")

	carIDInt, err := strconv.Atoi(carID)
	if err != nil {
		http.Error(w, "invalid car id", http.StatusBadRequest)
		return
	}

	logEntryIDInt, err := strconv.Atoi(logEntryID)
	if err != nil {
		http.Error(w, "invalid log entry id", http.StatusBadRequest)
		return
	}

	form := requests.EditLogEntryFormRequest{}
	if err := h.formDecoder.Decode(&form, r.PostForm); err != nil {
		http.Error(w, "invalid form data", http.StatusBadRequest)
		return
	}

	dateTime, err := time.Parse("2006-01-02", form.Date)
	if err != nil {
		http.Error(w, "invalid date", http.StatusBadRequest)
		return
	}

	output, err := h.editLogEntryUseCase.Execute(r.Context(), carusecases.EditLogEntryInput{
		CarID:      carIDInt,
		LogEntryID: logEntryIDInt,
		Name:       form.Name,
		Type:       entities.LogEntryCategory(form.EntryType),
		Date:       dateTime,
		KM:         form.KM,
		Cost:       form.Cost,
	})
	if err != nil {
		if errors.Is(err, persistence.ErrCarNotFound) {
			http.Error(w, "car not found", http.StatusNotFound)
			return
		}
		if errors.Is(err, persistence.ErrLogEntryNotFound) {
			http.Error(w, "log entry not found", http.StatusNotFound)
			return
		}
		if errors.Is(err, persistence.ErrInvalidLogEntryType) {
			http.Error(w, "invalid entry type", http.StatusBadRequest)
			return
		}
		http.Error(w, "failed to edit log entry", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", TurboStreamContentType)

	if err := h.tmpl.ExecuteTemplate(w, "edit_log_entry_list_item_response", output); err != nil {
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

func (h *HomeHandler) DeleteLogEntry(w http.ResponseWriter, r *http.Request) {
	carID := r.PathValue("carId")
	logEntryID := r.PathValue("logEntryId")

	carIDInt, err := strconv.Atoi(carID)
	if err != nil {
		http.Error(w, "invalid car id", http.StatusBadRequest)
		return
	}

	logEntryIDInt, err := strconv.Atoi(logEntryID)
	if err != nil {
		http.Error(w, "invalid log entry id", http.StatusBadRequest)
		return
	}

	resp, err := h.deleteLogEntryUseCase.Execute(r.Context(), carusecases.DeleteLogEntryInput{
		CarID:      carIDInt,
		LogEntryID: logEntryIDInt,
	})

	if err != nil {
		http.Error(w, "failed to delete log entry", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", TurboStreamContentType)

	if err := h.tmpl.ExecuteTemplate(w, "delete_log_entry_list_item_response", resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
