package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/eiei114/go-backend-template/application/auth"
	"github.com/eiei114/go-backend-template/application/service"
	"github.com/eiei114/go-backend-template/interface/request"
	"github.com/eiei114/go-backend-template/interface/response"
	"github.com/uptrace/bunrouter"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (u *UserHandler) UserCreateHandle() bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		var requestData request.UserCreateRequest
		if err := json.NewDecoder(req.Body).Decode(&requestData); err != nil {
			http.Error(w, "Failed to parse request", http.StatusBadRequest)
			return err
		}

		ctx := req.Context()
		authToken, err := u.userService.Add(ctx, requestData.Name)
		if err != nil {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return err
		}

		responseData := &response.UserCreateResponse{Token: authToken}
		responseBytes, err := json.Marshal(responseData)
		if err != nil {
			http.Error(w, "Failed to generate response", http.StatusInternalServerError)
			return err
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseBytes)
		return nil
	}
}

func (u *UserHandler) UserGetHandle() bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {

		ctx := req.Context()

		id := auth.GetUserIDFromContext(ctx)
		// Retrieve user by auth token
		user, err := u.userService.GetUserByUserId(ctx, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}

		// Prepare the response using UserGetResponse struct
		responseData := &response.UserGetResponse{
			Id:    user.Id,
			Name:  user.Name,
			Count: user.Count,
		}

		respBytes, err := json.Marshal(responseData)
		if err != nil {
			http.Error(w, "Failed to generate response", http.StatusInternalServerError)
			return err
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(respBytes)

		return nil
	}
}

func (u *UserHandler) CountAddHandle() bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		var requestData request.UserCountAddRequest
		if err := json.NewDecoder(req.Body).Decode(&requestData); err != nil {
			http.Error(w, "Failed to parse request", http.StatusBadRequest)
			return err
		}
		ctx := req.Context()
		id := auth.GetUserIDFromContext(ctx)

		log.Println("id", id)

		user, err := u.userService.GetUserByUserId(ctx, id)

		log.Println("user", user.Name)
		if err != nil {
			http.Error(w, "Failed to Count Update user", http.StatusInternalServerError)
			return err
		}

		count := requestData.Count

		log.Println("count", count)
		log.Println("user.Count", user.Count)

		user.Count = user.Count + count
		log.Println("user.Count", user.Count)
		_ = u.userService.UpdateUser(ctx, user)

		return nil
	}
}

func (u *UserHandler) DestroyHandle() bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		ctx := req.Context()
		id := auth.GetUserIDFromContext(ctx)

		// Delete the user using userService
		if _, err := u.userService.Delete(ctx, id); err != nil {
			http.Error(w, "Failed to delete user", http.StatusInternalServerError)
			return err
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "User successfully deleted"}`))
		return nil
	}
}