package web

import (
	"airline/backend/internal/infrastructure/web/handler"
	"airline/backend/internal/infrastructure/web/middleware"

	"github.com/gin-gonic/gin"

	_ "airline/backend/docs" // Import generated docs

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// NewRouter creates and configures the Gin router.
func NewRouter(h *handler.VoucherHandler) *gin.Engine {
	r := gin.Default()

	// Use CORS middleware (for development)
	// In production, Docker Nginx proxy will handle this
	r.Use(middleware.CORSMiddleware())

	// API routes
	api := r.Group("/api")
	{
		api.POST("/check", h.CheckAssignment)
		api.POST("/generate", h.GenerateAssignment)
	}

	// Swagger documentation route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
