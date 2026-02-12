package store

import "context"

type UserStore interface{
	SelectUser(ctx context.Context, login string) (*User, error)
	InsertUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, id int) error
}

// // internal/service/auth_service.go
// type AuthService interface {
//     Register(ctx context.Context, user *User) error
//     Login(ctx context.Context, login, password string) (*User, error)
// }

// // internal/handlers/handler.go
// type Handler struct {
//     authService AuthService
//     userStore   UserStore
// }

// func NewHandler(authService AuthService, userStore UserStore) *Handler {
//     return &Handler{
//         authService: authService,
//         userStore:   userStore,
//     }
// }

// func (h *Handler) Registration(w http.ResponseWriter, r *http.Request) {
//     // использовать h.authService вместо прямого доступа
// }