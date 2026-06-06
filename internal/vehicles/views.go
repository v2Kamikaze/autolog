package vehicles

import (
	"github.com/v2code/autolog/internal/ui"
)

func HomeView(vehicles []*Vehicle) ui.View {
	return ui.View{
		TemplateName: "base_page",
		Data: ui.Data{
			"Vehicles": vehicles,
		},
	}
}

func VehicleCreatedView(vehicle *Vehicle) ui.View {
	return ui.View{
		TemplateName: "add_vehicle_stream",
		Data: ui.Data{
			"Vehicle": vehicle,
		},
	}
}

func VehicleUpdatedView(vehicle *Vehicle) ui.View {
	return ui.View{
		TemplateName: "edit_vehicle_stream",
		Data: ui.Data{
			"Vehicle": vehicle,
		},
	}
}

func VehicleDeletedView(vehicle *Vehicle) ui.View {
	return ui.View{
		TemplateName: "delete_vehicle_stream",
		Data: ui.Data{
			"Vehicle": vehicle,
		},
	}
}

func LogEntryAddedView(vehicle *Vehicle, entry *LogEntry) ui.View {
	return ui.View{
		TemplateName: "add_log_stream",
		Data: ui.Data{
			"LogEntry": entry,
			"Vehicle":  vehicle,
		},
	}
}

func LogEntryUpdatedView(vehicle *Vehicle, entry *LogEntry) ui.View {
	return ui.View{
		TemplateName: "edit_log_stream",
		Data: ui.Data{
			"LogEntry": entry,
			"Vehicle":  vehicle,
		},
	}
}

func LogEntryDeletedView(vehicle *Vehicle, entry *LogEntry) ui.View {
	return ui.View{
		TemplateName: "delete_log_stream",
		Data: ui.Data{
			"LogEntry": entry,
			"Vehicle":  vehicle,
		},
	}
}
