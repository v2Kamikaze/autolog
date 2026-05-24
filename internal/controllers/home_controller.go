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
	"github.com/v2code/autolog/internal/usecase/vehicleusecases"
)

const TurboStreamContentType = "text/vnd.turbo-stream.html"

type HomeHandler struct {
	listVehiclesUseCase       *vehicleusecases.ListVehiclesUseCase
	createVehicleUseCase      *vehicleusecases.CreateVehicleUseCase
	addLogEntryUseCase    *vehicleusecases.AddLogEntryUseCase
	editLogEntryUseCase   *vehicleusecases.EditLogEntryUseCase
	editVehicleUseCase        *vehicleusecases.EditVehicleUseCase
	deleteVehicleUseCase      *vehicleusecases.DeleteVehicleUseCase
	deleteLogEntryUseCase *vehicleusecases.DeleteLogEntryUseCase
	formDecoder           *schema.Decoder
	tmpl                  *template.Template
}

func NewHomeHandler(
	listVehiclesUseCase *vehicleusecases.ListVehiclesUseCase,
	createVehicleUseCase *vehicleusecases.CreateVehicleUseCase,
	addLogEntryUseCase *vehicleusecases.AddLogEntryUseCase,
	editLogEntryUseCase *vehicleusecases.EditLogEntryUseCase,
	editVehicleUseCase *vehicleusecases.EditVehicleUseCase,
	deleteVehicleUseCase *vehicleusecases.DeleteVehicleUseCase,
	deleteLogEntryUseCase *vehicleusecases.DeleteLogEntryUseCase,
	tmpl *template.Template,
) *HomeHandler {
	formDecoder := schema.NewDecoder()
	formDecoder.IgnoreUnknownKeys(true)

	return &HomeHandler{
		listVehiclesUseCase:       listVehiclesUseCase,
		createVehicleUseCase:      createVehicleUseCase,
		addLogEntryUseCase:    addLogEntryUseCase,
		editLogEntryUseCase:   editLogEntryUseCase,
		editVehicleUseCase:        editVehicleUseCase,
		deleteVehicleUseCase:      deleteVehicleUseCase,
		deleteLogEntryUseCase: deleteLogEntryUseCase,
		formDecoder:           formDecoder,
		tmpl:                  tmpl,
	}
}

func (h *HomeHandler) Home(w http.ResponseWriter, r *http.Request) {
	output, err := h.listVehiclesUseCase.Execute(r.Context())
	if err != nil {
		http.Error(w, "failed to list vehicles", http.StatusInternalServerError)
		return
	}

	if err := h.tmpl.ExecuteTemplate(w, "base_page", output); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *HomeHandler) CreateVehicle(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "failed to parse form", http.StatusBadRequest)
		return
	}

	form := requests.CreateVehicleFormRequest{}
	if err := h.formDecoder.Decode(&form, r.PostForm); err != nil {
		http.Error(w, "invalid form data", http.StatusBadRequest)
		return
	}

	output, err := h.createVehicleUseCase.Execute(r.Context(), vehicleusecases.CreateVehicleInput{
		Title: form.Title,
		Type:  form.VehicleType,
		KM:    form.KM,
	})
	if err != nil {
		http.Error(w, "failed to create vehicle", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", TurboStreamContentType)

	if err := h.tmpl.ExecuteTemplate(w, "add_vehicle_list_item_response", output); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *HomeHandler) AddLogEntry(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "failed to parse form", http.StatusBadRequest)
		return
	}

	vehicleID := r.PathValue("id")

	vehicleIDInt, err := strconv.Atoi(vehicleID)
	if err != nil {
		http.Error(w, "invalid vehicle id", http.StatusBadRequest)
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

	output, err := h.addLogEntryUseCase.Execute(r.Context(), vehicleusecases.AddLogEntryInput{
		VehicleID: vehicleIDInt,
		Name:  form.Name,
		Type:  entities.LogEntryCategory(form.EntryType),
		Date:  dateTime,
		KM:    form.KM,
		Cost:  form.Cost,
	})
	if err != nil {
		if errors.Is(err, persistence.ErrVehicleNotFound) {
			http.Error(w, "vehicle not found", http.StatusNotFound)
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

	vehicleID := r.PathValue("vehicleId")
	logEntryID := r.PathValue("logEntryId")

	vehicleIDInt, err := strconv.Atoi(vehicleID)
	if err != nil {
		http.Error(w, "invalid vehicle id", http.StatusBadRequest)
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

	output, err := h.editLogEntryUseCase.Execute(r.Context(), vehicleusecases.EditLogEntryInput{
		VehicleID:      vehicleIDInt,
		LogEntryID: logEntryIDInt,
		Name:       form.Name,
		Type:       entities.LogEntryCategory(form.EntryType),
		Date:       dateTime,
		KM:         form.KM,
		Cost:       form.Cost,
	})
	if err != nil {
		if errors.Is(err, persistence.ErrVehicleNotFound) {
			http.Error(w, "vehicle not found", http.StatusNotFound)
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

func (h *HomeHandler) EditVehicle(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "failed to parse form", http.StatusBadRequest)
		return
	}

	form := requests.EditVehicleFormRequest{}
	if err := h.formDecoder.Decode(&form, r.PostForm); err != nil {
		http.Error(w, "invalid form data", http.StatusBadRequest)
		return
	}

	vehicleID := r.PathValue("id")

	vehicleIDInt, err := strconv.Atoi(vehicleID)
	if err != nil {
		http.Error(w, "invalid vehicle id", http.StatusBadRequest)
		return
	}

	resp, err := h.editVehicleUseCase.Execute(r.Context(), vehicleusecases.EditVehicleInput{
		ID:    vehicleIDInt,
		Title: form.Title,
		Type:  form.VehicleType,
		KM:    form.KM,
	})
	if err != nil {
		if errors.Is(err, persistence.ErrVehicleNotFound) {
			http.Error(w, "vehicle not found", http.StatusNotFound)
			return
		}
		http.Error(w, "failed to edit vehicle", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", TurboStreamContentType)

	if err := h.tmpl.ExecuteTemplate(w, "edit_vehicle_list_item_response", resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *HomeHandler) DeleteVehicle(w http.ResponseWriter, r *http.Request) {
	vehicleID := r.PathValue("id")

	vehicleIDInt, err := strconv.Atoi(vehicleID)
	if err != nil {
		http.Error(w, "invalid vehicle id", http.StatusBadRequest)
		return
	}

	resp, err := h.deleteVehicleUseCase.Execute(r.Context(), vehicleusecases.DeleteVehicleInput{
		ID: vehicleIDInt,
	})

	if err != nil {
		http.Error(w, "failed to delete vehicle", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", TurboStreamContentType)

	if err := h.tmpl.ExecuteTemplate(w, "delete_vehicle_list_item_response", resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *HomeHandler) DeleteLogEntry(w http.ResponseWriter, r *http.Request) {
	vehicleID := r.PathValue("vehicleId")
	logEntryID := r.PathValue("logEntryId")

	vehicleIDInt, err := strconv.Atoi(vehicleID)
	if err != nil {
		http.Error(w, "invalid vehicle id", http.StatusBadRequest)
		return
	}

	logEntryIDInt, err := strconv.Atoi(logEntryID)
	if err != nil {
		http.Error(w, "invalid log entry id", http.StatusBadRequest)
		return
	}

	resp, err := h.deleteLogEntryUseCase.Execute(r.Context(), vehicleusecases.DeleteLogEntryInput{
		VehicleID:      vehicleIDInt,
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
