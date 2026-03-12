package auth

import (
	"example.com/trial-go/internal/middleware"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Repo *Repository
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Role     string `json:"role"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserListResponse struct {
	Data []User `json:"data"`
	Meta Meta   `json:"meta"`
}

type Meta struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

func (s *Service) Login(req LoginRequest) (map[string]interface{}, error) {
	if req.Username == "" || req.Password == "" {
		return nil, ErrBadRequest("username and password required")
	}
	account, err := s.Repo.FindAccountByUsername(req.Username)
	if err != nil {
		return nil, ErrUnauthorized("invalid username or password")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(req.Password)); err != nil {
		return nil, ErrUnauthorized("invalid username or password")
	}
	token, err := middleware.GenerateToken(account.User.ID, account.User.Role)
	if err != nil {
		return nil, ErrInternal("failed to generate token")
	}
	return map[string]interface{}{
		"message": "login successful",
		"token":   token,
		"user":    account.User,
		"role":    account.User.Role,
	}, nil
}

func (s *Service) Register(req RegisterRequest) (map[string]interface{}, error) {
	if req.Name == "" || req.Role == "" || req.Username == "" || req.Password == "" {
		return nil, ErrBadRequest("all fields required")
	}
	if s.Repo.IsUsernameExists(req.Username) {
		return nil, ErrConflict("username already exists")
	}
	user := User{Name: req.Name, Role: req.Role}
	if err := s.Repo.CreateUser(&user); err != nil {
		return nil, ErrInternal("failed to create user")
	}
	hashed, err := HashPassword(req.Password)
	if err != nil {
		return nil, ErrInternal("failed to hash password")
	}
	account := Account{
		Username: req.Username,
		Password: hashed,
		UserID:   user.ID,
	}
	if err := s.Repo.CreateAccount(&account); err != nil {
		return nil, ErrInternal("failed to create account")
	}
	return map[string]interface{}{"message": "user registered"}, nil
}

func (s *Service) GetUsers(page int, limit int) (*UserListResponse, error) {

	users, total, err := s.Repo.GetUsers(page, limit)
	if err != nil {
		return nil, err
	}

	totalPages := int(total) / limit
	if int(total)%limit != 0 {
		totalPages++
	}

	return &UserListResponse{
		Data: users,
		Meta: Meta{
			Page:       page,
			Limit:      limit,
			Total:      total,
			TotalPages: totalPages,
		},
	}, nil
}

func (s *Service) GetCurrentUser(userID uint) (*User, error) {

	user, err := s.Repo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) DeleteUser(id uint) error {

	err := s.Repo.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}

func HashPassword(password string) (string, error) {
	if password == "" {
		return "", ErrBadRequest("password cannot be empty")
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// Error helpers
func ErrBadRequest(msg string) error   { return &ServiceError{Code: 400, Msg: msg} }
func ErrUnauthorized(msg string) error { return &ServiceError{Code: 401, Msg: msg} }
func ErrConflict(msg string) error     { return &ServiceError{Code: 409, Msg: msg} }
func ErrInternal(msg string) error     { return &ServiceError{Code: 500, Msg: msg} }

// ServiceError for custom error handling
type ServiceError struct {
	Code int
	Msg  string
}

func (e *ServiceError) Error() string { return e.Msg }
