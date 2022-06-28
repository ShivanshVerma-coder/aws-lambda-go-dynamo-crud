package handlers

import (
	"net/http"

	"github.com/ShivanshVerma-coder/aws-lambda-go/pkg/repositories"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var ErrorMethodNotAllowed = "method not allowed"

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

func GetUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	email := req.QueryStringParameters["email"]
	if len(email) > 0 {
		result, err := repositories.FetchUser(email, tableName, dynaClient)
		if err != nil {
			return ApiResponse(http.StatusBadRequest, ErrorBody{ErrorMsg: aws.String(err.Error())})
		}
		return ApiResponse(http.StatusOK, result)
	}
	return nil, nil
}

func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	result, err := repositories.CreateUser(req, tableName, dynaClient)
	if err != nil {
		return ApiResponse(http.StatusBadRequest, ErrorBody{
			ErrorMsg: aws.String(err.Error()),
		})
	}
	return ApiResponse(http.StatusCreated, result)
}

func UpdateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	result, err := repositories.UpdateUser(req, tableName, dynaClient)
	if err != nil {
		return ApiResponse(http.StatusBadRequest, ErrorBody{
			ErrorMsg: aws.String(err.Error()),
		})
	}
	return ApiResponse(http.StatusOK, result)
}

func DeleteUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	err := repositories.DeleteUser(req, tableName, dynaClient)
	if err != nil {
		return ApiResponse(http.StatusBadRequest, ErrorBody{
			ErrorMsg: aws.String(err.Error()),
		})
	}
	return ApiResponse(http.StatusOK, "Deleted Successfully")
}

func UnhandledMethod() (*events.APIGatewayProxyResponse, error) {
	return ApiResponse(http.StatusBadGateway, ErrorBody{
		ErrorMsg: aws.String(ErrorMethodNotAllowed),
	})
}
