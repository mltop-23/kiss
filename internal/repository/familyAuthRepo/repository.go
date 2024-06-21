package FamilyAuthRepo

import (
	"context"
	"database/sql"
	"errors"
	"kissandeat/internal/structs"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type FamilyAuthRepository interface {
	// Family management
	RegisterFamily(ctx context.Context, family *structs.Family) error
	UpdateFamily(ctx context.Context, family *structs.Family) error
	DeleteFamily(ctx context.Context, familyID int) error
	GetFamily(ctx context.Context, familyID int) (*structs.Family, error)

	// Member management
	CreateMember(ctx context.Context, member *structs.User) error
	UpdateMember(ctx context.Context, member *structs.User) error
	DeleteMember(ctx context.Context, memberID int) error
	GetMember(ctx context.Context, memberID int) (*structs.User, error)

	// Authentication and authorization
	LoginMember(ctx context.Context, email, password string) (string, error)
	ValidateToken(ctx context.Context, token string) (*structs.User, error)
	LogoutMember(ctx context.Context, token string) error
}

type FamilyAuthPostgres struct {
	db *sql.DB
}

func NewFamilyAuthPostgres(db *sql.DB) *FamilyAuthPostgres {
	return &FamilyAuthPostgres{db: db}
}

// Family management
func (r *FamilyAuthPostgres) RegisterFamily(ctx context.Context, family *structs.Family) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO families (name) VALUES ($1)", family.ID)
	return err
}

func (r *FamilyAuthPostgres) UpdateFamily(ctx context.Context, family *structs.Family) error {
	_, err := r.db.ExecContext(ctx, "UPDATE families SET name = $1 WHERE id = $2", family.ID, family.ID)
	return err
}

func (r *FamilyAuthPostgres) DeleteFamily(ctx context.Context, familyID int) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM families WHERE id = $1", familyID)
	return err
}

func (r *FamilyAuthPostgres) GetFamily(ctx context.Context, familyID int) (*structs.Family, error) {
	var family structs.Family
	err := r.db.QueryRowContext(ctx, "SELECT id, name FROM families WHERE id = $1", familyID).Scan(&family.ID, &family.ID)
	return &family, err
}

// Member management
func (r *FamilyAuthPostgres) CreateMember(ctx context.Context, member *structs.User) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO users ( username, password, email, firstName, lastName, gender, role) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		member.Username, member.Password, member.Email, member.FirstName, member.LastName, member.Gender, member.Role)
	return err
}

func (r *FamilyAuthPostgres) UpdateMember(ctx context.Context, member *structs.User) error {
	_, err := r.db.ExecContext(ctx, "UPDATE users SET username = $1, password = $2, email = $3, first_name = $4, last_name = $5, gender = $6, role = $7 WHERE id = $8",
		member.Username, member.Password, member.Email, member.FirstName, member.LastName, member.Gender, member.Role, member.ID)
	return err
}

func (r *FamilyAuthPostgres) DeleteMember(ctx context.Context, memberID int) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM users WHERE id = $1", memberID)
	return err
}

func (r *FamilyAuthPostgres) GetMember(ctx context.Context, memberID int) (*structs.User, error) {
	var member structs.User
	err := r.db.QueryRowContext(ctx, "SELECT id, username, password, email, first_name, last_name, gender, role FROM users WHERE id = $1", memberID).Scan(
		&member.ID, &member.Username, &member.Password, &member.Email, &member.FirstName, &member.LastName, &member.Gender, &member.Role)
	return &member, err
}

// Authentication and authorization
func (r *FamilyAuthPostgres) LoginMember(ctx context.Context, email, password string) (string, error) {
	var user structs.User
	err := r.db.QueryRowContext(ctx, "SELECT id, password FROM users WHERE email = $1", email).Scan(&user.ID, &user.Password)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *FamilyAuthPostgres) ValidateToken(ctx context.Context, token string) (*structs.User, error) {
	claims, err := ParseToken(token)
	if err != nil {
		return nil, err
	}

	var user structs.User
	err = r.db.QueryRowContext(ctx, "SELECT id, username, email, first_name, last_name, gender, role FROM users WHERE id = $1", claims.UserID).Scan(
		&user.ID, &user.Username, &user.Email, &user.FirstName, &user.LastName, &user.Gender, &user.Role)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *FamilyAuthPostgres) LogoutMember(ctx context.Context, token string) error {
	// В данном примере logout не требует действий на сервере,
	// поскольку JWT токен является самодостаточным.
	// Если требуется ревокация токенов, можно использовать черный список или другой механизм.
	return nil
}

// JWT helpers
func GenerateToken(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte("your-secret-key"))
}

func ParseToken(tokenStr string) (*structs.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &structs.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("your-secret-key"), nil
	})

	if claims, ok := token.Claims.(*structs.Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
