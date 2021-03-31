package handler

import (
	"encoding/json"
	"errors"
	"net/http"
)

type AuthenticateRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       uint   `json:"age"`
}

type AuthenticateResponse struct {
	ID      string `json:"id,omitempty"`
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

func authenticateEndpoint(userService UserService, logger Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := AuthenticateResponse{}

		request, err := validateAuthenticateRequest(r)
		if err != nil {
			response.Message = err.Error()
			sendResponse(w, response, http.StatusBadRequest, logger)
			return
		}

		usr, err := userService.NewUser(request.FirstName, request.LastName, request.Age)
		if err != nil {
			response.Message = err.Error()
			sendResponse(w, response, http.StatusBadRequest, logger)
			return
		}

		response.ID = usr.ID
		response.Success = true

		sendResponse(w, response, 200, logger)
	})
}

func validateAuthenticateRequest(r *http.Request) (*AuthenticateRequest, error) {
	request := AuthenticateRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return &request, err
	}

	if request.FirstName == "" || len(request.FirstName) < 3 {
		return &request, errors.New("Invalid request param first_name")
	}

	if request.LastName == "" || len(request.LastName) < 3 {
		return &request, errors.New("Invalid request param last_name")
	}

	if request.Age < 0 || request.Age > 99 {
		return &request, errors.New("Invalid request param age")
	}

	return &request, nil
}
