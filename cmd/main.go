package main

import (
	"fmt"

	// Import Swagger docs

	"github.com/gin-gonic/gin"
	"github.com/rogeriofontes/cert-generator/api"
	"github.com/rogeriofontes/cert-generator/config"
	"github.com/rogeriofontes/cert-generator/internal/app"
	"github.com/rogeriofontes/cert-generator/internal/infra/db"
	"github.com/rogeriofontes/cert-generator/internal/infra/email"
	"github.com/rogeriofontes/cert-generator/internal/infra/pdf"
	"github.com/rogeriofontes/cert-generator/middleware"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/swag/example/basic/docs"
)

// @title API de Certificados
// @version 1.0
// @description API para geração e validação de certificados.
// @termsOfService http://meusistema.com/termos/

// @contact.name Suporte da API
// @contact.url http://meusistema.com/suporte
// @contact.email suporte@meusistema.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:9393
// @BasePath /
// @schemes http
func main() {
	// 🔹 Carrega configurações do .env
	//database.InitDB()
	// Inicializa o banco de dados
	cfg := config.LoadConfig()
	config.InitDB()

	communityRepo := &db.CommunityRepo{DB: config.DB}     // 🔹 Repositório de comunidades
	eventRepo := &db.EventRepo{DB: config.DB}             // 🔹 Repositório de eventos
	participantRepo := &db.ParticipantRepo{DB: config.DB} // 🔹 Repositório de participantes

	pdfGen := &pdf.PDFService{BackgroundPath: cfg.BackgroundPath, OutputDir: cfg.OutputDir, ParticipantRepo: participantRepo}
	emailSvc := &email.EmailServiceImpl{SMTPUser: cfg.SMTPUser, SMTPPass: cfg.SMTPPass}

	communityService := &app.CommunityService{CommunityRepo: communityRepo}
	eventService := &app.EventService{EventRepo: eventRepo}
	participantService := &app.ParticipantService{ParticipantRepo: participantRepo}

	// Initialize userHandler
	userRepo := &db.UserRepo{DB: config.DB}
	userService := &app.UserService{UserRepo: userRepo}
	userHandler := &api.UserHandler{UserService: userService}

	// 🔹 Instancia a App Service
	certificateService := &app.CertificateService{
		EventRepo:       eventRepo,
		ParticipantRepo: participantRepo,
		PdfGen:          pdfGen,
		EmailSvc:        emailSvc,
	}

	// 🔹 Instancia os handlers separados
	h := api.Handler{
		CertificateService: certificateService,
	}
	e := api.EventHandler{
		EventService: eventService,
	}
	p := api.ParticipantHandler{
		ParticipantService: participantService,
	}
	c := api.CommunityHandler{
		CommunityService: communityService,
	}

	r := gin.Default()

	// Configuração do Swagger
	//docs.SwaggerInfo.BasePath = "/"
	// Adicionar Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Endpoints

	// Rota pública de login
	//r.POST("/login", api.Login)

	// Rota pública de registro
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/register", userHandler.Register)
		userRoutes.POST("/login", userHandler.Login)
	}

	// ✅ Rotas protegidas por JWT (aplicando o middleware)
	protected := r.Group("/")
	protected.Use(middleware.JWTMiddleware()) // 🔒 Exige JWT para acessar essas rotas

	// ✅ Rotas para comunidades (Handler Original)
	communitiesRoutes := protected.Group("/communities")
	{
		communitiesRoutes.POST("", c.CreateCommunity)
		communitiesRoutes.GET("", c.BuscarComunidades)
	}

	// ✅ Rotas para eventos (Handler Separado)
	eventRoutes := protected.Group("/events")
	{
		eventRoutes.POST("", e.CreateEvent)
		eventRoutes.GET("", e.GetEvents)

		eventRoutes.GET("/:eventID/participants", p.GetParticipantesByEvent)
		eventRoutes.GET("/participants/:id", p.GetParticipantByEvent)

	}

	participantRoutes := protected.Group("/participants")
	{
		participantRoutes.POST("", p.CreateParticipant)
		participantRoutes.GET("", p.GetParticipants)
		//participantRoutes.GET("/validate", p.ValidateParticipantCertificate)
	}

	participantRoutesNoProtected := r.Group("/certificate-participants")
	{
		participantRoutesNoProtected.GET("/validate", p.ValidateParticipantCertificate)
	}

	// ✅ Rotas para certificados (Handler Original)
	certificateRoutes := protected.Group("/certificates")
	{
		certificateRoutes.POST("/event/:eventID", h.GenerateCertificatesByEvent)
		certificateRoutes.POST("/user/:userID", h.GenerateCertificateForUser)
		certificateRoutes.POST("/pending", h.GeneratePendingCertificates)
	}

	// 🔹 Altere a porta para uma disponível (exemplo: 9090)
	port := "9393"
	fmt.Printf("🚀 Servidor rodando na porta %s\n", port)
	r.Run(":" + port)
}
