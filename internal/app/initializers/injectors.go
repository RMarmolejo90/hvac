package initializers

import (
	"github.com/rmarmolejo90/hvac/internal/adapters/database/postgres"
	"github.com/rmarmolejo90/hvac/internal/adapters/http"
	"github.com/rmarmolejo90/hvac/internal/app/services"
	"gorm.io/gorm"
)

// Handlers struct holds all the HTTP handlers.
type Handlers struct {
	CustomerHandler         *http.CustomerHandler
	LocationHandler         *http.LocationHandler
	TechnicianHandler       *http.TechnicianHandler
	JobHandler              *http.JobHandler
	JobStatusHistoryHandler *http.JobStatusHistoryHandler
	ServiceHandler          *http.ServiceHandler
	InvoiceHandler          *http.InvoiceHandler
	PaymentHandler          *http.PaymentHandler
	QuoteHandler            *http.QuoteHandler
	EquipmentHandler        *http.EquipmentHandler
	ConsumablesHandler      *http.ConsumablesHandler
	StockHandler            *http.StockHandler
	TagHandler              *http.TagHandler
	NoteHandler             *http.NoteHandler
	HourlyRateHandler       *http.HourlyRateHandler
	ScheduleHandler         *http.ScheduleHandler
}

// InitHandlers initializes all the handlers with their respective services.
func InitHandlers() *Handlers {
	// Initialize Repositories
	customerRepo := postgres.NewCustomerRepository()
	locationRepo := postgres.NewLocationRepository()
	technicianRepo := postgres.NewTechnicianRepository()
	jobRepo := postgres.NewJobRepository()
	jobStatusHistoryRepo := postgres.NewJobStatusHistoryRepository()
	serviceRepo := postgres.NewServiceRepository()
	invoiceRepo := postgres.NewInvoiceRepository()
	paymentRepo := postgres.NewPaymentRepository()
	quoteRepo := postgres.NewQuoteRepository()
	equipmentRepo := postgres.NewEquipmentRepository()
	consumablesRepo := postgres.NewConsumablesRepository()
	stockRepo := postgres.NewStockRepository()
	tagRepo := postgres.NewTagRepository()
	noteRepo := postgres.NewNoteRepository()
	hourlyRateRepo := postgres.NewHourlyRateRepository()
	scheduleRepo := postgres.NewScheduleRepository()

	// Initialize Services
	customerService := services.NewCustomerService(customerRepo)
	locationService := services.NewLocationService(locationRepo)
	technicianService := services.NewTechnicianService(technicianRepo)
	jobService := services.NewJobService(jobRepo)
	jobStatusHistoryService := services.NewJobStatusHistoryService(jobStatusHistoryRepo)
	serviceService := services.NewServiceService(serviceRepo)
	invoiceService := services.NewInvoiceService(invoiceRepo)
	paymentService := services.NewPaymentService(paymentRepo)
	quoteService := services.NewQuoteService(quoteRepo)
	equipmentService := services.NewEquipmentService(equipmentRepo)
	consumablesService := services.NewConsumablesService(consumablesRepo)
	stockService := services.NewStockService(stockRepo)
	tagService := services.NewTagService(tagRepo)
	noteService := services.NewNoteService(noteRepo)
	hourlyRateService := services.NewHourlyRateService(hourlyRateRepo)
	scheduleService := services.NewScheduleService(scheduleRepo)

	// Initialize Handlers
	return &Handlers{
		CustomerHandler:         http.NewCustomerHandler(customerService),
		LocationHandler:         http.NewLocationHandler(locationService),
		TechnicianHandler:       http.NewTechnicianHandler(technicianService),
		JobHandler:              http.NewJobHandler(jobService),
		JobStatusHistoryHandler: http.NewJobStatusHistoryHandler(jobStatusHistoryService),
		ServiceHandler:          http.NewServiceHandler(serviceService),
		InvoiceHandler:          http.NewInvoiceHandler(invoiceService),
		PaymentHandler:          http.NewPaymentHandler(paymentService),
		QuoteHandler:            http.NewQuoteHandler(quoteService),
		EquipmentHandler:        http.NewEquipmentHandler(equipmentService),
		ConsumablesHandler:      http.NewConsumablesHandler(consumablesService),
		StockHandler:            http.NewStockHandler(stockService),
		TagHandler:              http.NewTagHandler(tagService),
		NoteHandler:             http.NewNoteHandler(noteService),
		HourlyRateHandler:       http.NewHourlyRateHandler(hourlyRateService),
		ScheduleHandler:         http.NewScheduleHandler(scheduleService),
	}
}
