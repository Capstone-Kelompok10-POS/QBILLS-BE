package routes

import (
	"os"
	"qbills/handler"
	"qbills/repository"
	"qbills/services"
	"qbills/utils/helpers/middleware"

	echoJwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func MembershipCardRoutes(e *echo.Echo, db *gorm.DB) {
	membershipCardRepository := repository.NewMembershipCardRepository(db)
	membershipCardService := services.NewMembershipCardService(membershipCardRepository)
	MembershipCardHandler := handler.NewMembershipCardHandler(membershipCardService)

	membershipCardGroup := e.Group("api/v1/membership/card")

	membershipCardGroup.Use(echoJwt.JWT([]byte(os.Getenv("SECRET_KEY"))))

	membershipCardGroup.GET("/:id", MembershipCardHandler.PrintMembershipCardHandler, middleware.AuthMiddleware("Cashier"))
	// membershipCardGroup.GET("/barcode/:", handler.ShowBarcodeHandler, middleware.AuthMiddleware("Cashier"))
}
