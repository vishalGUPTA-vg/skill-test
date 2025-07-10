package pdf


import (
   "bytes"
   "fmt"


   "go-service/client"


   "github.com/phpdave11/gofpdf"
)


type PDFGenerator struct{}


func NewPDFGenerator() *PDFGenerator {
   return &PDFGenerator{}
}


func (g *PDFGenerator) Generate(student *client.Student) ([]byte, error) {
   pdf := gofpdf.New("P", "mm", "A4", "")
   pdf.AddPage()
   pdf.SetFont("Arial", "B", 16)
   pdf.Cell(40, 10, "Student Report")
   pdf.Ln(20)
   pdf.SetFont("Arial", "", 12)
   pdf.Cell(40, 10, fmt.Sprintf("ID: %d", student.ID))
   pdf.Ln(10)
   pdf.Cell(40, 10, fmt.Sprintf("Name: %s", student.Name))
   pdf.Ln(10)
   pdf.Cell(40, 10, fmt.Sprintf("Email: %s", student.Email))
   pdf.Ln(10)
   pdf.Cell(40, 10, fmt.Sprintf("Class: %s", student.Class))
   pdf.Ln(10)
   pdf.Cell(40, 10, fmt.Sprintf("Section: %s", student.Section))
   pdf.Ln(10)
   pdf.Cell(40, 10, fmt.Sprintf("Roll: %s", student.Roll))
   pdf.Ln(10)
   pdf.Cell(40, 10, fmt.Sprintf("Father's Name: %s", student.FatherName))
   pdf.Ln(10)
   pdf.Cell(40, 10, fmt.Sprintf("Mother's Name: %s", student.MotherName))
   pdf.Ln(10)
   pdf.Cell(40, 10, fmt.Sprintf("Guardian's Name: %s", student.GuardianName))
   pdf.Ln(10)
   pdf.Cell(40, 10, fmt.Sprintf("Admission Date: %s", student.AdmissionDate))
   pdf.Ln(10)
   pdf.Cell(40, 10, fmt.Sprintf("Reporter Name: %s", student.ReporterName))
   // Add more fields as needed


   buf := new(bytes.Buffer)
   if err := pdf.Output(buf); err != nil {
       return nil, err
   }
   return buf.Bytes(), nil
}
