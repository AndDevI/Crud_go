package user

import (
	"errors"
	"net/mail"
	"strings"
	"time"
)

var (
	ErrInvalidID       = errors.New("invalid user id")
	ErrInvalidName     = errors.New("user name is required")
	ErrInvalidEmail    = errors.New("user email is invalid")
	ErrInvalidPassword = errors.New("password must have at least 6 characters")
	ErrInvalidGroupID  = errors.New("group_id must be greater than zero")
	ErrInvalidActive   = errors.New("active must be 0 or 1")
)

type User struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Telephone *string    `json:"telephone,omitempty"`
	Password  string     `json:"-"`
	Image     *string    `json:"image,omitempty"`
	Active    int        `json:"active"`
	GroupID   *int64     `json:"group_id,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

func NewUser(name, email, password string, telephone, image *string, groupID *int64, active *int) (User, error) {
	normalizedName, err := NormalizeName(name)
	if err != nil {
		return User{}, err
	}

	normalizedEmail, err := NormalizeEmail(email)
	if err != nil {
		return User{}, err
	}

	if err := ValidatePassword(password); err != nil {
		return User{}, err
	}

	if err := ValidateGroupID(groupID); err != nil {
		return User{}, err
	}

	resolvedActive := 1
	if active != nil {
		if err := ValidateActive(*active); err != nil {
			return User{}, err
		}
		resolvedActive = *active
	}

	return User{
		Name:      normalizedName,
		Email:     normalizedEmail,
		Password:  password,
		Telephone: NormalizeOptionalString(telephone),
		Image:     NormalizeOptionalString(image),
		Active:    resolvedActive,
		GroupID:   groupID,
	}, nil
}

func NormalizeName(name string) (string, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return "", ErrInvalidName
	}
	return name, nil
}

func NormalizeEmail(email string) (string, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	if _, err := mail.ParseAddress(email); err != nil {
		return "", ErrInvalidEmail
	}
	return email, nil
}

func ValidatePassword(password string) error {
	if len(strings.TrimSpace(password)) < 6 {
		return ErrInvalidPassword
	}
	return nil
}

func ValidateGroupID(groupID *int64) error {
	if groupID != nil && *groupID <= 0 {
		return ErrInvalidGroupID
	}
	return nil
}

func ValidateActive(active int) error {
	if active != 0 && active != 1 {
		return ErrInvalidActive
	}
	return nil
}

func NormalizeOptionalString(value *string) *string {
	if value == nil {
		return nil
	}
	trimmed := strings.TrimSpace(*value)
	if trimmed == "" {
		return nil
	}
	return &trimmed
}
