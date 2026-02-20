package register

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Role     string `json:"role"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Repository struct {
	Db *gorm.DB
}

type Handler struct {
	Repo *Repository
}

func (h *Handler) Register(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	if req.Name == "" || req.Role == "" || req.Username == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "all fields required"})
	}

	// Cek username sudah ada
	var existing Account
	if err := h.Repo.Db.Where("username = ?", req.Username).First(&existing).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "username already exists"})
	}

	// Buat user
	user := User{Name: req.Name, Role: req.Role}
	if err := h.Repo.Db.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to create user"})
	}

	// Hash password
	hashed, err := HashPassword(req.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to hash password"})
	}

	// Buat account
	account := Account{
		Username: req.Username,
		Password: hashed,
		UserID:   user.ID,
	}
	if err := h.Repo.Db.Create(&account).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to create account"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "user registered"})
}

// HashPassword hashes the password using bcrypt
func HashPassword(password string) (string, error) {
	if password == "" {
		return "", fiber.NewError(fiber.StatusBadRequest, "password cannot be empty")
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}
