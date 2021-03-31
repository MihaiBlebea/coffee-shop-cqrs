package handler

import (
	"net/http"
)

type OrdersRequest struct {
	CoffeeID string `json:"coffee_id"`
	UserID   string `json:"user_id"`
}

type OrdersResponse struct {
	TransactionID string `json:"transaction_id,omitempty"`
	Success       bool   `json:"success"`
	Message       string `json:"message,omitempty"`
}

func ordersEndpoint(logger Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := OrdersResponse{}

		sendResponse(w, response, 200, logger)
	})
}
