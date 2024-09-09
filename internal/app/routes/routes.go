package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rmarmolejo90/hvac/internal/app/initializers"
)

// SetupRoutes initializes all the routes for the application
func SetupRoutes(r *gin.Engine, handlers *initializers.Handlers) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server is Running!!",
		})
	})

	// Customer routes
	r.POST("/customers", handlers.CustomerHandler.CreateCustomer)
	r.GET("/customers/:id", handlers.CustomerHandler.GetCustomer)
	r.PUT("/customers/:id", handlers.CustomerHandler.UpdateCustomer)
	r.DELETE("/customers/:id", handlers.CustomerHandler.DeleteCustomer)
	r.GET("/customers", handlers.CustomerHandler.ListCustomers)

	// ContactMethod routes
	r.POST("/contact-methods", handlers.ContactMethodHandler.CreateContactMethod)
	r.GET("/contact-methods/:id", handlers.ContactMethodHandler.GetContactMethod)
	r.PUT("/contact-methods/:id", handlers.ContactMethodHandler.UpdateContactMethod)
	r.DELETE("/contact-methods/:id", handlers.ContactMethodHandler.DeleteContactMethod)
	r.GET("/customers/:customer_id/contact-methods", handlers.ContactMethodHandler.ListContactMethodsByCustomerID)

	// Location routes
	r.POST("/locations", handlers.LocationHandler.CreateLocation)
	r.GET("/locations/:id", handlers.LocationHandler.GetLocation)
	r.PUT("/locations/:id", handlers.LocationHandler.UpdateLocation)
	r.DELETE("/locations/:id", handlers.LocationHandler.DeleteLocation)
	r.GET("/customers/:customer_id/locations", handlers.LocationHandler.ListLocationsByCustomerID)

	// Technician routes
	r.POST("/technicians", handlers.TechnicianHandler.CreateTechnician)
	r.GET("/technicians/:id", handlers.TechnicianHandler.GetTechnician)
	r.PUT("/technicians/:id", handlers.TechnicianHandler.UpdateTechnician)
	r.DELETE("/technicians/:id", handlers.TechnicianHandler.DeleteTechnician)
	r.GET("/technicians", handlers.TechnicianHandler.ListTechnicians)

	// Job routes
	r.POST("/jobs", handlers.JobHandler.CreateJob)
	r.GET("/jobs/:id", handlers.JobHandler.GetJob)
	r.PUT("/jobs/:id", handlers.JobHandler.UpdateJob)
	r.DELETE("/jobs/:id", handlers.JobHandler.DeleteJob)
	r.GET("/jobs", handlers.JobHandler.ListJobs)

	// Job Status History routes
	r.POST("/job-status-history", handlers.JobStatusHistoryHandler.CreateJobStatusHistory)
	r.GET("/jobs/:job_id/status-history", handlers.JobStatusHistoryHandler.ListJobStatusHistoryByJobID)

	// Service routes
	r.POST("/services", handlers.ServiceHandler.CreateService)
	r.GET("/services/:id", handlers.ServiceHandler.GetService)
	r.PUT("/services/:id", handlers.ServiceHandler.UpdateService)
	r.DELETE("/services/:id", handlers.ServiceHandler.DeleteService)
	r.GET("/services", handlers.ServiceHandler.ListServices)

	// Invoice routes
	r.POST("/invoices", handlers.InvoiceHandler.CreateInvoice)
	r.GET("/invoices/:id", handlers.InvoiceHandler.GetInvoice)
	r.PUT("/invoices/:id", handlers.InvoiceHandler.UpdateInvoice)
	r.DELETE("/invoices/:id", handlers.InvoiceHandler.DeleteInvoice)
	r.GET("/invoices", handlers.InvoiceHandler.ListInvoices)

	// Payment routes
	r.POST("/payments", handlers.PaymentHandler.CreatePayment)
	r.GET("/payments/:id", handlers.PaymentHandler.GetPayment)
	r.PUT("/payments/:id", handlers.PaymentHandler.UpdatePayment)
	r.DELETE("/payments/:id", handlers.PaymentHandler.DeletePayment)
	r.GET("/invoices/:invoice_id/payments", handlers.PaymentHandler.ListPaymentsByInvoiceID)

	// Quote routes
	r.POST("/quotes", handlers.QuoteHandler.CreateQuote)
	r.GET("/quotes/:id", handlers.QuoteHandler.GetQuote)
	r.PUT("/quotes/:id", handlers.QuoteHandler.UpdateQuote)
	r.DELETE("/quotes/:id", handlers.QuoteHandler.DeleteQuote)
	r.GET("/quotes", handlers.QuoteHandler.ListQuotes)

	// Equipment routes
	r.POST("/equipment", handlers.EquipmentHandler.CreateEquipment)
	r.GET("/equipment/:id", handlers.EquipmentHandler.GetEquipment)
	r.PUT("/equipment/:id", handlers.EquipmentHandler.UpdateEquipment)
	r.DELETE("/equipment/:id", handlers.EquipmentHandler.DeleteEquipment)
	r.GET("/locations/:location_id/equipment", handlers.EquipmentHandler.ListEquipmentByLocationID)

	// Consumables routes
	r.POST("/consumables", handlers.ConsumablesHandler.CreateConsumable)
	r.GET("/consumables/:id", handlers.ConsumablesHandler.GetConsumable)
	r.PUT("/consumables/:id", handlers.ConsumablesHandler.UpdateConsumable)
	r.DELETE("/consumables/:id", handlers.ConsumablesHandler.DeleteConsumable)
	r.GET("/equipment/:equipment_id/consumables", handlers.ConsumablesHandler.ListConsumablesByEquipmentID)

	// Stock routes
	r.POST("/stock", handlers.StockHandler.CreateStock)
	r.GET("/stock/:id", handlers.StockHandler.GetStock)
	r.PUT("/stock/:id", handlers.StockHandler.UpdateStock)
	r.DELETE("/stock/:id", handlers.StockHandler.DeleteStock)
	r.GET("/technicians/:technician_id/stock", handlers.StockHandler.ListStockByTechnicianID)

	// Tag routes
	r.POST("/tags", handlers.TagHandler.CreateTag)
	r.GET("/tags/:id", handlers.TagHandler.GetTag)
	r.PUT("/tags/:id", handlers.TagHandler.UpdateTag)
	r.DELETE("/tags/:id", handlers.TagHandler.DeleteTag)
	r.GET("/tags", handlers.TagHandler.ListTags)

	// Note routes
	r.POST("/notes", handlers.NoteHandler.CreateNote)
	r.GET("/notes/:id", handlers.NoteHandler.GetNote)
	r.PUT("/notes/:id", handlers.NoteHandler.UpdateNote)
	r.DELETE("/notes/:id", handlers.NoteHandler.DeleteNote)
	r.GET("/customers/:customer_id/notes", handlers.NoteHandler.ListNotesByCustomerID)
	r.GET("/locations/:location_id/notes", handlers.NoteHandler.ListNotesByLocationID)
	r.GET("/jobs/:job_id/notes", handlers.NoteHandler.ListNotesByJobID)

	// HourlyRate routes
	r.POST("/hourly-rates", handlers.HourlyRateHandler.CreateHourlyRate)
	r.GET("/hourly-rates/:id", handlers.HourlyRateHandler.GetHourlyRate)
	r.PUT("/hourly-rates/:id", handlers.HourlyRateHandler.UpdateHourlyRate)
	r.DELETE("/hourly-rates/:id", handlers.HourlyRateHandler.DeleteHourlyRate)
	r.GET("/hourly-rates", handlers.HourlyRateHandler.ListHourlyRates)

	// Schedule routes
	r.POST("/schedules", handlers.ScheduleHandler.CreateSchedule)
	r.GET("/schedules/:id", handlers.ScheduleHandler.GetSchedule)
	r.PUT("/schedules/:id", handlers.ScheduleHandler.UpdateSchedule)
	r.DELETE("/schedules/:id", handlers.ScheduleHandler.DeleteSchedule)
	r.GET("/technicians/:technician_id/schedules", handlers.ScheduleHandler.ListSchedulesByTechnicianID)

	// Event routes
	r.POST("/events", handlers.EventHandler.CreateEvent)
	r.GET("/events/:id", handlers.EventHandler.GetEvent)
	r.PUT("/events/:id", handlers.EventHandler.UpdateEvent)
	r.DELETE("/events/:id", handlers.EventHandler.DeleteEvent)
	r.GET("/schedules/:schedule_id/events", handlers.EventHandler.ListEventsByScheduleID)
}
