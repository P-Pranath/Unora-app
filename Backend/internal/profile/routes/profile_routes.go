// internal/profile/routes/profile_routes.go
package routes

import (
	"github.com/gin-gonic/gin"

	ent "github.com/UnoraApp/be/ent/generated"
	authservices "github.com/UnoraApp/be/internal/auth/services"
	"github.com/UnoraApp/be/internal/profile/handlers"
	"github.com/UnoraApp/be/internal/profile/services"
	"github.com/UnoraApp/be/pkg/storage"
)

// RegisterProfileRoutes registers all profile-related routes
func RegisterProfileRoutes(
	router *gin.RouterGroup,
	entClient *ent.Client,
	storageClient storage.Client,
	authService *authservices.AuthService,
	authMiddleware gin.HandlerFunc,
) {
	// Create services
	profileService := services.NewProfileService(entClient, storageClient)
	photoService := services.NewPhotoService(entClient, storageClient)

	// Create handler
	profileHandler := handlers.NewProfileHandler(profileService, photoService)

	// Public routes
	// Hobby options can be viewed without auth
	router.GET("/hobby-options", profileHandler.GetHobbyOptions)

	// Protected routes (require authentication)
	protected := router.Group("")
	protected.Use(authMiddleware)
	{
		// Profile CRUD
		protected.POST("/profile", profileHandler.CreateProfile)
		protected.GET("/profile", profileHandler.GetProfile)
		protected.PUT("/profile", profileHandler.UpdateProfile)
		protected.DELETE("/profile", profileHandler.DeleteProfile)
		protected.GET("/profile/onboarding-status", profileHandler.GetOnboardingStatus)

		// Photos
		protected.POST("/photos/upload-url", profileHandler.GetUploadURL)
		protected.POST("/photos", profileHandler.RegisterPhoto)
		protected.GET("/photos", profileHandler.GetUserPhotos)
		protected.PUT("/photos/:id/order", profileHandler.UpdatePhotoOrder)
		protected.DELETE("/photos/:id", profileHandler.DeletePhoto)

		// Hobbies
		protected.POST("/hobbies", profileHandler.AddHobby)
		protected.GET("/hobbies", profileHandler.GetUserHobbies)
		protected.DELETE("/hobbies/:id", profileHandler.DeleteHobby)
	}
}
