package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// AppConfig holds application configuration
type AppConfig struct {
	ID          string
	Environment string
	Version     string
	Logger      *log.Logger
}

// APIResponse is a standardized response format
type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type App struct {
	Config *AppConfig
}

// NewAppConfig creates a new application configuration
func NewAppConfig(id, env, version string) *AppConfig {
	return &AppConfig{
		ID:          id,
		Environment: env,
		Version:     version,
		Logger:      log.Default(),
	}
}

// NewApp creates a new application instance
func NewApp(config *AppConfig) *App {
	return &App{
		Config: config,
	}
}

// Handler processes API Gateway requests
func (app *App) Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Log incoming request
	app.Config.Logger.Printf("Received request: %s %s", request.HTTPMethod, request.Path)

	responseData := APIResponse{
		Status:  "success",
		Message: "Deployed with Lambda... API is alive!",
		Data: map[string]string{
			"app_id":      app.Config.ID,
			"environment": app.Config.Environment,
			"version":     app.Config.Version,
		},
	}

	// Marshal response to JSON
	body, err := json.Marshal(responseData)
	if err != nil {
		app.Config.Logger.Printf("Error marshaling response: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       `{"status":"error","message":"Internal server error"}`,
			Headers:    getDefaultHeaders(),
		}, nil
	}

	// Return successful response
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
		Headers:    getDefaultHeaders(),
	}, nil
}

// getDefaultHeaders returns standard response headers including CORS
func getDefaultHeaders() map[string]string {
	return map[string]string{
		"Content-Type":                 "application/json",
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": "GET, POST, OPTIONS",
		"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key",
	}
}

func main() {
	// Initialize application with configuration
	appConfig := NewAppConfig(
		"123",
		"production",
		"1.0.0",
	)

	app := NewApp(appConfig)
	app.Config.Logger.Printf("Starting application with ID: %s", appConfig.ID)

	// Start Lambda handler
	/*
			AWS Lambda Needs a Function to Execute
				AWS Lambda expects you to provide a handler function that processes incoming events
				(e.g., an HTTP request from API Gateway, an S3 file upload, or a message from an SQS queue).

				2. lambda.Start() Registers Your Function
				When you write lambda.Start(app.Handler), you are telling AWS Lambda:

				"Hey AWS, whenever an event occurs,
				please call my Handler function to process it."

				This function should match AWS Lambda's expected format.
				func Handler(ctx context.Context, event MyEvent) (MyResponse, error) {
		    // Business logic
		}


	*/

	lambda.Start(app.Handler)
}
