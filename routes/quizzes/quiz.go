package quizzes

import (
	quizzesController "backend/controllers/quizzes"

	"github.com/labstack/echo"
)

// InitializeQuizRoutes initializes all quiz routes
func InitializeQuizRoutes(quizzes *echo.Group, adminQuizzes *echo.Group) {
	quizzes.GET("/past", quizzesController.GetPastQuizzes)
	quizzes.GET("/future", quizzesController.GetFutureQuizzes)
	quizzes.GET("/current", quizzesController.GetCurrentQuizzes)
	quizzes.GET("/quiz", quizzesController.GetQuizByID)
	quizzes.GET("/month", quizzesController.GetQuizzesByMonthAndYear)

	adminQuizzes.POST("", quizzesController.CreateQuiz)
	adminQuizzes.DELETE("", quizzesController.DeleteQuiz)
	adminQuizzes.PUT("", quizzesController.UpdateQuiz)
}
