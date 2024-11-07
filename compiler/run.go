package compiler

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/anuragrao04/pesuio-final-project/models"
	"github.com/gin-gonic/gin"
)

func executePython(code string) (string, error) {
	// Create the command
	cmd := exec.Command("python3", "-c", code)
	cmd := exec.Command("C:\\Users\\anina\\AppData\\Local\\Programs\\Python\\Python313\\python.exe", "-c", code)

	// Capture the output
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	// Run the command
	err := cmd.Run()
	if err != nil {
		return stderr.String(), fmt.Errorf("error executing Python code: %v", err)
	}

	return out.String(), nil
}

func Run(c *gin.Context) {
	var request models.RunRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid JSON format",
		})
	}