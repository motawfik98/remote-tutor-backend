package assignments

import (
	authController "backend/controllers/auth"
	submissionsFiles "backend/controllers/files/assignments"
	submissionsDBInteractions "backend/database/assignments"
	submissionsModel "backend/models/assignments"
	"backend/utils"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func GetSubmissionByUserAndAssignment(c echo.Context) error {
	assignmentID := utils.ConvertToUInt(c.QueryParam("assignmentID"))
	userID := authController.FetchLoggedInUserID(c)
	submission := submissionsDBInteractions.GetSubmissionByUserAndAssignment(userID, assignmentID)
	return c.JSON(http.StatusOK, echo.Map{
		"submission": submission,
	})
}

func CreateOrUpdateSubmission(c echo.Context) error {
	method := c.Request().Method
	submission := new(submissionsModel.AssignmentSubmission)
	submission.UserID = authController.FetchLoggedInUserID(c)
	submission.AssignmentID = utils.ConvertToUInt(c.FormValue("assignmentID"))
	submission.UploadedAt = time.Now()
	if method == "POST" {
		submissionsDBInteractions.CreateSubmission(submission)
	}
	submissionFilePath, submissionErr := submissionsFiles.UploadUserSubmissionFile(c, submission)
	if submissionErr != nil {
		if method == "POST" {
			submissionsDBInteractions.DeleteSubmission(submission)
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred when trying to upload the submission. Please try again later",
		})
	}
	if submissionFilePath != "" {
		submission.File = submissionFilePath
		submission.UploadedAt = time.Now()
	}
	submissionsDBInteractions.UpdateSubmission(submission)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "AssignmentSubmission Saved Successfully",
	})
}

func GetSubmissionsByAssignmentForAllUsers(c echo.Context) error {
	assignmentID := utils.ConvertToUInt(c.QueryParam("assignmentID"))
	fullNameSearch := c.QueryParam("searchBy")
	submissions, totalSubmissions := submissionsDBInteractions.GetSubmissionsByAssignmentForAllUsers(c, assignmentID, fullNameSearch)
	return c.JSON(http.StatusOK, echo.Map{
		"submissions": submissions,
		"totalSubmissions": totalSubmissions,
	})
}