package auth

import (
	"hsl/config"
	"hsl/internal/models"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	db  *sqlx.DB
	cfg *config.Config
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func New(db *sqlx.DB) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) ValidateCredentials(username, pwd string) error {
	var storedHash string

	err := s.db.Get(&storedHash, "SELECT password_hash FROM admin_user WHERE username = ?", username)
	if err != nil {
		log.Errorf("Unable to find user due: %s", err)
		return models.ErrUserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(pwd)); err != nil {
		log.Errorf("Unable to validate credentials provided due: %s", err)
		return models.ErrInvalidCredentials
	}

	_, err = s.db.Exec("UPDATE admin_user SET last_login = ? WHERE username = ?", time.Now(), username)

	return err
}

func (s *AuthService) GenerateToken(username string) (string, error) {
	expTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtKey := []byte(os.Getenv("JWT_KEY"))

	return token.SignedString(jwtKey)
}

func (s *AuthService) ValidateToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	jwtKey := []byte(s.cfg.JwtKey)

	token, err := jwt.ParseWithClaims(
		tokenStr,
		claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		log.Errorf("Unable to parse claims due: %s", err)
		return nil, err
	}

	if !token.Valid {
		log.Error("Unable to validate token.")
		return nil, models.ErrInvalidToken
	}

	return claims, nil
}
