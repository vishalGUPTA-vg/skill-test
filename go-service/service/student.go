package service

import (
	"fmt"
	"go-service/client"
	"go-service/pdf"
	"os"

	"github.com/gin-gonic/gin"
)

type StudentService struct {
	Client       *client.StudentClient
	PDFGenerator *pdf.PDFGenerator
}

func NewStudentService() *StudentService {
	return &StudentService{Client: client.NewStudentClient(os.Getenv("BACKENT_BASE_URL")), PDFGenerator: pdf.NewPDFGenerator()}
}

func (s *StudentService) GenerateStudentReport(ctx *gin.Context, studentID string) ([]byte, error) {
	student, err := s.Client.FetchStudentData(ctx, studentID)
	if err != nil {
		fmt.Println("error in fetch ", err)
		return nil, err
	}
	return s.PDFGenerator.Generate(student)
}
