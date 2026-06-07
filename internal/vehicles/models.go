package vehicles

import (
	"strconv"
	"time"
)

type LogEntryCategory string

const (
	LogEntryCategoryMaintenance LogEntryCategory = "maintenance"
	LogEntryCategoryExpense     LogEntryCategory = "expense"
	LogEntryCategoryFuel        LogEntryCategory = "fuel"
	LogEntryCategoryInsurance   LogEntryCategory = "insurance"
	LogEntryCategoryTax         LogEntryCategory = "tax"
	LogEntryCategoryFine        LogEntryCategory = "fine"
	LogEntryCategoryParking     LogEntryCategory = "parking"
	LogEntryCategoryCleaning    LogEntryCategory = "cleaning"
	LogEntryCategoryToll        LogEntryCategory = "toll"
	LogEntryCategoryAccessory   LogEntryCategory = "accessory"
	LogEntryCategoryOther       LogEntryCategory = "other"
)

func (c LogEntryCategory) Valid() bool {
	switch c {
	case LogEntryCategoryMaintenance,
		LogEntryCategoryExpense,
		LogEntryCategoryFuel,
		LogEntryCategoryInsurance,
		LogEntryCategoryTax,
		LogEntryCategoryFine,
		LogEntryCategoryParking,
		LogEntryCategoryCleaning,
		LogEntryCategoryToll,
		LogEntryCategoryAccessory,
		LogEntryCategoryOther:
		return true
	default:
		return false
	}
}

func (c LogEntryCategory) Label() string {
	switch c {
	case LogEntryCategoryMaintenance:
		return "Manutenção"
	case LogEntryCategoryExpense:
		return "Despesa"
	case LogEntryCategoryFuel:
		return "Combustível"
	case LogEntryCategoryInsurance:
		return "Seguro"
	case LogEntryCategoryTax:
		return "Imposto"
	case LogEntryCategoryFine:
		return "Multa"
	case LogEntryCategoryParking:
		return "Estacionamento"
	case LogEntryCategoryCleaning:
		return "Lavagem"
	case LogEntryCategoryToll:
		return "Pedágio"
	case LogEntryCategoryAccessory:
		return "Acessório"
	case LogEntryCategoryOther:
		return "Outro"
	default:
		return string(c)
	}
}

type TransmissionType string

const (
	TransmissionManual    TransmissionType = "manual"
	TransmissionAutomatic TransmissionType = "automatic"
	TransmissionCVT       TransmissionType = "cvt"
	TransmissionAutomated TransmissionType = "automated"
)

func (t TransmissionType) Label() string {
	switch t {
	case TransmissionManual:
		return "Manual"
	case TransmissionAutomatic:
		return "Automático"
	case TransmissionCVT:
		return "CVT"
	case TransmissionAutomated:
		return "Automatizado"
	default:
		return string(t)
	}
}

type FuelType string

const (
	FuelGasoline FuelType = "gasoline"
	FuelEthanol  FuelType = "ethanol"
	FuelFlex     FuelType = "flex"
	FuelDiesel   FuelType = "diesel"
	FuelHybrid   FuelType = "hybrid"
	FuelElectric FuelType = "electric"
)

func (f FuelType) Label() string {
	switch f {
	case FuelGasoline:
		return "Gasolina"
	case FuelEthanol:
		return "Etanol"
	case FuelFlex:
		return "Flex"
	case FuelDiesel:
		return "Diesel"
	case FuelHybrid:
		return "Híbrido"
	case FuelElectric:
		return "Elétrico"
	default:
		return string(f)
	}
}

type LogEntry struct {
	ID        int
	VehicleID int

	Name string
	Type LogEntryCategory

	Date time.Time
	KM   int

	Cost int

	Notes string
}

type Vehicle struct {
	ID int

	Brand string
	Model string
	Year  int

	Version string
	Engine  string

	Transmission TransmissionType
	Fuel         FuelType

	KM int

	LogEntries []*LogEntry
}

func (v *Vehicle) DisplayName() string {
	name := ""

	if v.Brand != "" {
		name += v.Brand
	}

	if v.Model != "" {
		if name != "" {
			name += " "
		}
		name += v.Model
	}

	if v.Version != "" {
		name += " " + v.Version
	}

	if v.Engine != "" {
		name += " " + v.Engine
	}

	if v.Year > 0 {
		name += " (" + strconv.Itoa(v.Year) + ")"
	}

	return name
}

func (v *Vehicle) LogEntryCount() int {
	return len(v.LogEntries)
}

func (v *Vehicle) TotalCost() int {
	total := 0
	for _, entry := range v.LogEntries {
		total += entry.Cost
	}
	return total
}

func (v *Vehicle) MaintenanceCount() int {
	count := 0
	for _, entry := range v.LogEntries {
		if entry.Type == LogEntryCategoryMaintenance {
			count++
		}
	}
	return count
}

func (v *Vehicle) ExpenseCount() int {
	count := 0
	for _, entry := range v.LogEntries {
		if entry.Type == LogEntryCategoryExpense {
			count++
		}
	}
	return count
}
