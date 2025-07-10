package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-service/service"
)

func accountRoute(v1 *gin.RouterGroup) {
	v1.GET("/students/:id/report", HandleStudentReport)
}

func HandleStudentReport(c *gin.Context) {
	studentID := c.Param("id")
	// serv := service.NewStudentService()
	pdfBytes, err :=service.NewStudentService().GenerateStudentReport(c,studentID)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to generate report: %s", err.Error())
		return
	}

	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=student_report.pdf")
	c.Data(http.StatusOK, "application/pdf", pdfBytes)
}
