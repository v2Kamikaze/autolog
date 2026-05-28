package vehicles

import (
	"github.com/v2code/autolog/internal/persistence/entities"
	"github.com/v2code/autolog/internal/ui"
)

func Home(vehicles []*entities.Vehicle) ui.Renderer {
	return ui.View{
		TemplateName: "base_page",
		Data: ui.Data{
			"Vehicles": vehicles,
		},
	}
}

func VehicleCreated(vehicle *entities.Vehicle) ui.Renderer {
	return ui.View{
		TemplateName: "add_vehicle_list_item_response",
		Data: ui.Data{
			"Vehicle": vehicle,
		},
	}
}

func VehicleUpdated(vehicle *entities.Vehicle) ui.Renderer {
	return ui.View{
		TemplateName: "edit_vehicle_list_item_response",
		Data: ui.Data{
			"Vehicle": vehicle,
		},
	}
}

func VehicleDeleted(vehicle *entities.Vehicle) ui.Renderer {
	return ui.View{
		TemplateName: "delete_vehicle_list_item_response",
		Data: ui.Data{
			"Vehicle": vehicle,
		},
	}
}

func LogEntryAdded(vehicle *entities.Vehicle, entry *entities.LogEntry) ui.Renderer {
	return ui.View{
		TemplateName: "add_log_entry_list_item_response",
		Data: ui.Data{
			"LogEntry":      entry,
			"TotalCost":     vehicle.TotalCost(),
			"LogEntryCount": vehicle.LogEntryCount(),
		},
	}
}

func LogEntryUpdated(vehicle *entities.Vehicle, entry *entities.LogEntry) ui.Renderer {
	return ui.View{
		TemplateName: "edit_log_entry_list_item_response",
		Data: ui.Data{
			"LogEntry":  entry,
			"TotalCost": vehicle.TotalCost(),
		},
	}
}

func LogEntryDeleted(vehicle *entities.Vehicle, entry *entities.LogEntry) ui.Renderer {
	return ui.View{
		TemplateName: "delete_log_entry_list_item_response",
		Data: ui.Data{
			"LogEntry":      entry,
			"TotalCost":     vehicle.TotalCost(),
			"LogEntryCount": vehicle.LogEntryCount(),
		},
	}
}
