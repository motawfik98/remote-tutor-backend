package payments

import (
	paymentsController "backend/controllers/payments"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// InitializeRoutes initializes all the required routes to deal with payments
func InitializeRoutes(e *echo.Echo, adminRouter *echo.Group) {
	payments := e.Group("/payments", middleware.JWT([]byte(os.Getenv("JWT_TOKEN"))))
	payments.GET("/month", paymentsController.GetPaymentsByUserAndMonth)
	payments.GET("/week", paymentsController.GetPaymentsByUserAndWeek)

	adminPayments := adminRouter.Group("/payments")
	adminPayments.POST("", paymentsController.CreatePayment)
	adminPayments.DELETE("", paymentsController.DeletePayment)
}
