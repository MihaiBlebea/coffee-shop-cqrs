package handler

import (
	"encoding/json"
	"errors"
	"net/http"
)

type OrderRequest struct {
	CoffeeID string `json:"coffee_id"`
	UserID   string `json:"user_id"`
}

type OrderResponse struct {
	TransactionID string `json:"transaction_id,omitempty"`
	Success       bool   `json:"success"`
	Message       string `json:"message,omitempty"`
}

func orderEndpoint(userService UserService, logger Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := OrderResponse{}

		request, err := validateOrderRequest(r)
		if err != nil {
			response.Message = err.Error()
			sendResponse(w, response, http.StatusBadRequest, logger)
			return
		}

		transactionID, err := userService.OrderCoffee(request.UserID, request.CoffeeID)
		if err != nil {
			response.Message = err.Error()
			sendResponse(w, response, http.StatusBadRequest, logger)
			return
		}

		response.TransactionID = transactionID
		response.Success = true

		sendResponse(w, response, 200, logger)
	})
}

func validateOrderRequest(r *http.Request) (*OrderRequest, error) {
	request := OrderRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return &request, err
	}

	if request.CoffeeID == "" || len(request.CoffeeID) < 3 {
		return &request, errors.New("Invalid request param coffee_id")
	}

	if request.UserID == "" || len(request.UserID) < 3 {
		return &request, errors.New("Invalid request param user_id")
	}

	return &request, nil
}
