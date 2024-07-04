package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Lambda handler for AWS Lambda environment
func lambdaHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return routeRequest(request)
}

// HTTP handler for local testing
func httpHandler(w http.ResponseWriter, r *http.Request) {
	request := events.APIGatewayProxyRequest{
		HTTPMethod: r.Method,
		Path:       r.URL.Path,
	}
	response, err := routeRequest(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "\"%s\"", err.Error())
		return
	}
	w.WriteHeader(response.StatusCode)
	fmt.Fprintf(w, response.Body)
}

// Route the request based on the method and path
func routeRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch request.HTTPMethod {
	case http.MethodGet:
		if strings.HasPrefix(request.Path, "/v1/getinfo/") {
			return getInfo(request)
		}
	case http.MethodPost:
		if strings.HasPrefix(request.Path, "/v1/set/") {
			return setInfo(request)
		}
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusNotFound,
		Body:       "\"Not Found\"",
	}, nil
}

// Handler for GET /v1/getinfo/:uuid
func getInfo(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	uuid := strings.TrimPrefix(request.Path, "/v1/getinfo/")
	response := fmt.Sprintf("Get info for UUID: %s", uuid)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       fmt.Sprintf("\"%s\"", response),
	}, nil
}

// Handler for POST /v1/set/:uuid/of/:id
func setInfo(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	parts := strings.Split(request.Path, "/")
	if len(parts) != 6 {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "\"Invalid path parameters\"",
		}, nil
	}
	uuid := parts[3]
	id := parts[5]
	response := fmt.Sprintf("Set info for UUID: %s and ID: %s", uuid, id)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       fmt.Sprintf("\"%s\"", response),
	}, nil
}

func main() {
	if len(os.Getenv("_LAMBDA_SERVER_PORT")) > 0 {
		// Running in AWS Lambda environment
		lambda.Start(lambdaHandler)
	} else {
		// Running in local environment for testing
		http.HandleFunc("/", httpHandler)
		fmt.Println("Starting local server on :9000")
		log.Fatal(http.ListenAndServe(":9000", nil))
	}
}
