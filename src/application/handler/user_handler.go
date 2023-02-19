package handler

import (
  "encoding/json"
  "example_onion/application/service"
  "github.com/gorilla/mux"
  "net/http"
  "strconv"
)

type UserHandler struct {
  UserService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
  return &UserHandler{UserService: userService}
}

func (h *UserHandler) Register(r *mux.Router, userService *service.UserService) {
  h.UserService = userService
  r.HandleFunc("/users/{id}", h.GetUser).Methods("GET")
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id, err := strconv.Atoi(vars["id"])
  if err != nil {
      http.Error(w, "Invalid user ID", http.StatusBadRequest)
      return
  }

  user, err := h.UserService.FindUserByID(id)
  if err != nil {
      http.Error(w, "User not found", http.StatusNotFound)
      return
  }

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(user)
}
