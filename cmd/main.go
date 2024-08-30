package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rmarmolejo90/hvac/internal/app/initializers"
)

func main() {
	// Initialize the application components
	initializers.Init()

	// Initialize handlers through dependency injection
	handlers := initializers.InitHandlers()

	// Start the server
	r := gin.Default()

	// Define routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server is Running!!",
		})
	})

	r.POST("/customers", handlers.CustomerHandler.CreateCustomer)
	r.GET("/customers/:id", handlers.CustomerHandler.GetCustomer)
	r.PUT("/customers/:id", handlers.CustomerHandler.UpdateCustomer)
	r.DELETE("/customers/:id", handlers.CustomerHandler.DeleteCustomer)
	r.GET("/customers", handlers.CustomerHandler.ListCustomers)

	r.POST("/locations", handlers.LocationHandler.CreateLocation)
	r.GET("/locations/:id", handlers.LocationHandler.GetLocation)
	r.PUT("/locations/:id", handlers.LocationHandler.UpdateLocation)
	r.DELETE("/locations/:id", handlers.LocationHandler.DeleteLocation)
	r.GET("/customers/:customer_id/locations", handlers.LocationHandler.ListLocationsByCustomerID)

	r.POST("/technicians", handlers.TechnicianHandler.CreateTechnician)
	r.GET("/technicians/:id", handlers.TechnicianHandler.GetTechnician)
	r.PUT("/technicians/:id", handlers.TechnicianHandler.UpdateTechnician)
	r.DELETE("/technicians/:id", handlers.TechnicianHandler.DeleteTechnician)
	r.GET("/technicians", handlers.TechnicianHandler.ListTechnicians)

	r.POST("/jobs", handlers.JobHandler.CreateJob)
	r.GET("/jobs/:id", handlers.JobHandler.GetJob)
	r.PUT("/jobs/:id", handlers.JobHandler.UpdateJob)
	r.DELETE("/jobs/:id", handlers.JobHandler.DeleteJob)
	r.GET("/jobs", handlers.JobHandler.ListJobs)

	r.POST("/job-status-history", handlers.JobStatusHistoryHandler.CreateJobStatusHistory)
	r.GET("/jobs/:job_id/status-history", handlers.JobStatusHistoryHandler.ListJobStatusHistoryByJobID)

	r.POST("/services", handlers.ServiceHandler.CreateService)
	r.GET("/services/:id", handlers.ServiceHandler.GetService)
	r.PUT("/services/:id", handlers.ServiceHandler.UpdateService)
	r.DELETE("/services/:id", handlers.ServiceHandler.DeleteService)
	r.GET("/services", handlers.ServiceHandler.ListServices)

	r.POST("/invoices", handlers.InvoiceHandler.CreateInvoice)
	r.GET("/invoices/:id", handlers.InvoiceHandler.GetInvoice)
	r.PUT("/invoices/:id", handlers.InvoiceHandler.UpdateInvoice)
	r.DELETE("/invoices/:id", handlers.InvoiceHandler.DeleteInvoice)
	r.GET("/invoices", handlers.InvoiceHandler.ListInvoices)

	r.POST("/payments", handlers.PaymentHandler.CreatePayment)
	r.GET("/payments/:id", handlers.PaymentHandler.GetPayment)
	r.PUT("/payments/:id", handlers.PaymentHandler.UpdatePayment)
	r.DELETE("/payments/:id", handlers.PaymentHandler.DeletePayment)
	r.GET("/invoices/:invoice_id/payments", handlers.PaymentHandler.ListPaymentsByInvoiceID)

	r.POST("/quotes", handlers.QuoteHandler.CreateQuote)
	r.GET("/quotes/:id", handlers.QuoteHandler.GetQuote)
	r.PUT("/quotes/:id", handlers.QuoteHandler.UpdateQuote)
	r.DELETE("/quotes/:id", handlers.QuoteHandler.DeleteQuote)
	r.GET("/quotes", handlers.QuoteHandler.ListQuotes)

	r.POST("/equipment", handlers.EquipmentHandler.CreateEquipment)
	r.GET("/equipment/:id", handlers.EquipmentHandler.GetEquipment)
	r.PUT("/equipment/:id", handlers.EquipmentHandler.UpdateEquipment)
	r.DELETE("/equipment/:id", handlers.EquipmentHandler.DeleteEquipment)
	r.GET("/locations/:location_id/equipment", handlers.EquipmentHandler.ListEquipmentByLocationID)

	r.POST("/consumables", handlers.ConsumablesHandler.CreateConsumable)
	r.GET("/consumables/:id", handlers.ConsumablesHandler.GetConsumable)
	r.PUT("/consumables/:id", handlers.ConsumablesHandler.UpdateConsumable)
	r.DELETE("/consumables/:id", handlers.ConsumablesHandler.DeleteConsumable)
	r.GET("/equipment/:equipment_id/consumables", handlers.ConsumablesHandler.ListConsumablesByEquipmentID)

	r.POST("/stock", handlers.StockHandler.CreateStock)
	r.GET("/stock/:id", handlers.StockHandler.GetStock)
	r.PUT("/stock/:id", handlers.StockHandler.UpdateStock)
	r.DELETE("/stock/:id", handlers.StockHandler.DeleteStock)
	r.GET("/technicians/:technician_id/stock", handlers.StockHandler.ListStockByTechnicianID)

	r.POST("/tags", handlers.TagHandler.CreateTag)
	r.GET("/tags/:id", handlers.TagHandler.GetTag)
	r.PUT("/tags/:id", handlers.TagHandler.UpdateTag)
	r.DELETE("/tags/:id", handlers.TagHandler.DeleteTag)
	r.GET("/tags", handlers.TagHandler.ListTags)

	r.POST("/notes", handlers.NoteHandler.CreateNote)
	r.GET("/notes/:id", handlers.NoteHandler.GetNote)
	r.PUT("/notes/:id", handlers.NoteHandler.UpdateNote)
	r.DELETE("/notes/:id", handlers.NoteHandler.DeleteNote)
	r.GET("/customers/:customer_id/notes", handlers.NoteHandler.ListNotesByCustomerID)
	r.GET("/locations/:location_id/notes", handlers.NoteHandler.ListNotesByLocationID)
	r.GET("/jobs/:job_id/notes", handlers.NoteHandler.ListNotesByJobID)

	r.POST("/hourly-rates", handlers.HourlyRate
