package services

import (
    "auth-service/config"
    "auth-service/models"
    "auth-service/repository"
    "crypto/rand"
    "encoding/base64"
    "encoding/json"
    "errors"
    "time"
    "log"
    "fmt"

    "golang.org/x/crypto/bcrypt"
    "github.com/golang-jwt/jwt/v5"
    "gorm.io/gorm"
)

// AuthService defines the methods for user authentication
type AuthService interface {
    Register(user *models.User) error
    Login(identifier, password string) (accessToken string, refreshToken string, err error)
    Refresh(refreshToken string) (newAccessToken string, newRefreshToken string, err error)
    Logout(refreshToken string) error
    FindByEmailOrMobile(identifier, mobile string) (models.User, error)

}

// Concrete implementation of AuthService
type authService struct {
    repo repository.UserRepository
    cfg  config.Config
}

// NewAuthService initializes DB, auto-migrates User, and returns service instance
func NewAuthService(cfg config.Config) AuthService {
	db := cfg.DB
	if db == nil {
		panic("❌ Database connection is not initialized in config")
	}

	// Auto-migrate the User model
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("❌ Auto migration failed: %v", err)
	}

	log.Println("✅ User model migrated successfully")

	// Initialize and return the AuthService
	return &authService{
		repo: repository.NewUserRepository(db),
		cfg:  cfg,
	}
}

// Register registers a new user
func (s *authService) Register(user *models.User) error {
    existingUser, _ := s.repo.FindByEmailOrMobile(user.Email, user.Mobile)
    if existingUser.ID != 0 {
        return errors.New("email or mobile already registered")
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)
    return s.repo.CreateUser(user)
}

// Login validates credentials and returns access & refresh tokens
func (s *authService) Login(identifier, password string) (string, string, error) {
    // Fetch user by email or mobile
    user, err := s.repo.FindByEmailOrMobile(identifier, identifier)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
           log.Printf("Login failed: no user found for identifier %s", identifier)
           return "", "", errors.New("invalid credentials")
        }
        log.Printf("Login error (DB): %v", err)
        return "", "", err
    }
    if user.ID == 0 {
        log.Printf("Login failed: user not found for identifier %s, err: %v", identifier, err)
        return "","", errors.New("invalid credentials")
    }

    // Log both passwords for debugging
    fmt.Println("🔐 Hashed Password:", string(user.Password))
    fmt.Println("Stored Hash: %q\n", user.Password)
    fmt.Println("Input Password: %q\n", password)

    // Verify password
//     err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(strings.TrimSpace(password)));
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if  err != nil {
        fmt.Println("❌ bcrypt comparison failed:", err)
        return "","", errors.New("invalid credentials")
    }else {
     	fmt.Println("✅ Password matched!")
    }

    // Parse roles from JSON
    roles := []string{}
    if err := json.Unmarshal(user.Roles, &roles); err != nil {
        log.Println("Login failed: invalid role format for user %s", identifier)
        return "","", errors.New("invalid user role format")
    }

    // Generate tokens
    accessToken, err := s.createAccessToken(user.ID, user.Email, user.Mobile, roles)
    if err != nil {
        return "", "", err
    }

    refreshToken, err := s.createRefreshToken()
    if err != nil {
        return "", "", err
    }

    // Save refresh token
    expiresAt := time.Now().Add(7 * 24 * time.Hour)
    err = s.repo.StoreRefreshToken(user.ID, refreshToken, expiresAt)
    if err != nil {
        return "", "", err
    }

    return accessToken, refreshToken, nil
}


// Refresh refreshes JWT tokens
func (s *authService) Refresh(refreshToken string) (string, string, error) {
    rt, err := s.repo.FindRefreshToken(refreshToken)
    if err != nil || rt.ExpiresAt.Before(time.Now()) {
        return "", "", errors.New("invalid or expired refresh token")
    }

    user, err := s.repo.FindByID(rt.UserID)
    if err != nil || user.ID == 0 {
        return "", "", errors.New("user not found")
    }

    roles := []string{}
    if err := json.Unmarshal(user.Roles, &roles); err != nil {
        return "", "", errors.New("invalid roles format")
    }

    accessToken, err := s.createAccessToken(user.ID, user.Email,user.Mobile, roles)
    if err != nil {
        return "", "", err
    }

    newRefreshToken, err := s.createRefreshToken()
    if err != nil {
        return "", "", err
    }

    // Delete old refresh token and save new one
    _ = s.repo.DeleteRefreshToken(refreshToken)
    expiresAt := time.Now().Add(7 * 24 * time.Hour)
    err = s.repo.StoreRefreshToken(user.ID, newRefreshToken, expiresAt)
    if err != nil {
        return "", "", err
    }

    return accessToken, newRefreshToken, nil
}

// Logout deletes refresh token
func (s *authService) Logout(refreshToken string) error {
    return s.repo.DeleteRefreshToken(refreshToken)
}

// Find user by email or mobile
func (s *authService) FindByEmailOrMobile(email string, mobile string) (models.User, error) {
	return s.repo.FindByEmailOrMobile(email, mobile)
}

// createAccessToken creates a JWT token valid for 15 minutes
func (s *authService) createAccessToken(userID uint, email string, mobile string, roles []string) (string, error) {
    claims := jwt.MapClaims{
        "id":    userID,
        "email": email,
        "mobile": mobile,
        "roles": roles,
        "exp":   time.Now().Add(15 * time.Minute).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(s.cfg.JWTSecret))
}

// createRefreshToken generates a secure random refresh token
func (s *authService) createRefreshToken() (string, error) {
    b := make([]byte, 32)
    _, err := rand.Read(b)
    if err != nil {
        return "", err
    }
    return base64.URLEncoding.EncodeToString(b), nil
}