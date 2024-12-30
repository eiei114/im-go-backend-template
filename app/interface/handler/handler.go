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

// @Summary ユーザーの作成
// @Description ユーザーを作成します
// @Tags user
// @Accept json
// @Produce json
// @Param request body request.UserCreateRequest true "UserCreateRequest"
// @Success 200 {object} response.UserCreateResponse
// @Router /user/create [post]
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

// @Summary ユーザーの取得
// @Description ユーザーを取得します
// @Tags user
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.UserGetResponse
// @Router /user/get [post]
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

// @Summary ユーザーのカウントを追加
// @Description ユーザーのカウントを追加します
// @Tags user
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body request.UserCountAddRequest true "UserCountAddRequest"
// @Success 200 {object} response.UserCountAddResponse
// @Router /user/count [post]
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

		// Prepare the response using UserCountAddResponse struct
		count := requestData.Count

		log.Println("count", count)
		log.Println("user.Count", user.Count)

		user.Count = user.Count + count
		log.Println("user.Count", user.Count)
		_ = u.userService.UpdateUser(ctx, user)

		responseData := &response.UserCountAddResponse{
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

// @Summary ユーザーの削除
// @Description ユーザーを削除します
// @Tags user
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.UserDestroyResponse
// @Router /user/destroy [post]
func (u *UserHandler) DestroyHandle() bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		ctx := req.Context()
		id := auth.GetUserIDFromContext(ctx)

		// Delete the user using userService
		if _, err := u.userService.Delete(ctx, id); err != nil {
			http.Error(w, "Failed to delete user", http.StatusInternalServerError)
			return err
		}

		responseData := &response.UserDestroyResponse{
			Message: "User successfully deleted",
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
