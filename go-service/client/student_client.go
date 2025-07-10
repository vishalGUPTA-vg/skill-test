package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StudentClient struct {
	BackendBaseURL string
}

func NewStudentClient(baseURL string) *StudentClient {
	return &StudentClient{BackendBaseURL: baseURL}
}

func (c *StudentClient) FetchStudentData(ctx *gin.Context, studentID string) (*Student, error) {
	url := fmt.Sprintf("%s/api/v1/students/%s", c.BackendBaseURL, studentID)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Extract cookies from the Gin context and attach them to the new request
	for _, cookie := range ctx.Request.Cookies() {
		req.AddCookie(cookie)
	}
	fmt.Println("x-csrf-token", ctx.GetHeader("x-csrf-token"))
	req.Header.Set("x-csrf-token", ctx.GetHeader("x-csrf-token"))

	// Use an HTTP client to send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		res := make(map[string]any)
		if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return nil, err
		}

		fmt.Println("invalid code ", res)
		return nil, fmt.Errorf("backend returned status %d", resp.StatusCode)
	}

	var student Resp
	if err := json.NewDecoder(resp.Body).Decode(&student); err != nil {
		return nil, err
	}

	fmt.Println("student found ", student)
	return &student.Student, nil
}
