package authFamilyService

import (
	"context"
	"kissandeat/internal/repository"
	"kissandeat/internal/structs"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repository.FamilyAuthRepository
}

func NewAuthService(repo repository.FamilyAuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

// Family management
func (s *AuthService) RegisterFamily(ctx context.Context, family *structs.Family) error {
	return s.repo.RegisterFamily(ctx, family)
}

func (s *AuthService) UpdateFamily(ctx context.Context, family *structs.Family) error {
	return s.repo.UpdateFamily(ctx, family)
}

func (s *AuthService) DeleteFamily(ctx context.Context, familyID int) error {
	return s.repo.DeleteFamily(ctx, familyID)
}

func (s *AuthService) GetFamily(ctx context.Context, familyID int) (*structs.Family, error) {
	return s.repo.GetFamily(ctx, familyID)
}
func (s *AuthService) ListFamilies(ctx context.Context) ([]*structs.Family, error) {
	return s.repo.ListFamilies(ctx)
}

// func (s *AuthService) ListMembers(ctx context.Context) ([]*structs.User, error) {
// 	return s.repo.ListMembers(ctx)
// }

// Member management
func (s *AuthService) CreateMember(ctx context.Context, member *structs.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(member.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	member.Password = string(hashedPassword)
	return s.repo.CreateMember(ctx, member)
}

func (s *AuthService) UpdateMember(ctx context.Context, member *structs.User) error {
	if member.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(member.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		member.Password = string(hashedPassword)
	}
	return s.repo.UpdateMember(ctx, member)
}

func (s *AuthService) DeleteMember(ctx context.Context, memberID int) error {
	return s.repo.DeleteMember(ctx, memberID)
}

func (s *AuthService) GetMember(ctx context.Context, memberID int) (*structs.User, error) {
	return s.repo.GetMember(ctx, memberID)
}

// Authentication and authorization
func (s *AuthService) LoginMember(ctx context.Context, email, password string) (string, error) {
	return s.repo.LoginMember(ctx, email, password)
}

// func (s *AuthService) ValidateToken(ctx context.Context, token string) (*structs.User, error) {
// 	return s.repo.ValidateToken(ctx, token)
// }

func (s *AuthService) LogoutMember(ctx context.Context, token string) error {
	return s.repo.LogoutMember(ctx, token)
}
