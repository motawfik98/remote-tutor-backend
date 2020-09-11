package quizzes

import "mime/multipart"

// Question struct to store the question data
type Question struct {
	ID        uint           `json:"ID"`
	Text      string         `json:"text"`
	TotalMark int            `json:"totalMark"`
	Quiz      Quiz           `json:"quiz" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	QuizID    uint           `json:"quizID"`
	ImagePath string         `json:"imagePath"`
	Image     multipart.File `json:"image" gorm:"-"`
}

// MCQ struct to store the MCQ question type data
type MCQ struct {
	Question      `json:"question"`
	CorrectAnswer uint     `json:"correctAnswer"`
	Choices       []Choice `json:"choices" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// LongAnswer struct to store the LongAnswer question type data
type LongAnswer struct {
	Question      `json:"question"`
	CorrectAnswer string `json:"correctAnswer"`
}
