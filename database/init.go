package database

import (
	announcementsModel "backend/models/announcements"
	quizzesModel "backend/models/quizzes"
	usersModel "backend/models/users"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // to import the gorm database wrapper
)

var (
	databaseConnection, err = gorm.Open("mysql", "root:password@/tutoring?charset=utf8&parseTime=True&loc=Local")
)

// MigrateTables makes sure that the tables are migrated at the start of the application
func MigrateTables() {
	if err == nil {
		databaseConnection.LogMode(true)
		databaseConnection.AutoMigrate(&usersModel.User{})
		databaseConnection.AutoMigrate(&announcementsModel.Announcement{})

		databaseConnection.AutoMigrate(&quizzesModel.Quiz{})
		databaseConnection.AutoMigrate(&quizzesModel.MCQ{}).AddForeignKey("quiz_id", "quizzes(id)", "CASCADE", "CASCADE")
		databaseConnection.AutoMigrate(&quizzesModel.LongAnswer{}).AddForeignKey("quiz_id", "quizzes(id)", "CASCADE", "CASCADE")
		databaseConnection.AutoMigrate(&quizzesModel.Choice{}).AddForeignKey("mcq_id", "mcqs(id)", "CASCADE", "CASCADE")
		databaseConnection.AutoMigrate(&quizzesModel.MCQSubmission{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").AddForeignKey("question_id", "mcqs(id)", "CASCADE", "CASCADE")
		databaseConnection.AutoMigrate(&quizzesModel.LongAnswerSubmission{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").AddForeignKey("question_id", "long_answers(id)", "CASCADE", "CASCADE")
		databaseConnection.AutoMigrate(&quizzesModel.QuizGrade{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").AddForeignKey("quiz_id", "quizzes(id)", "CASCADE", "CASCADE")
	}
}

// GetDBConnection returns the DB connection
func GetDBConnection() *gorm.DB {
	return databaseConnection
}
