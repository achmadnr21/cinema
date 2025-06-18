package usecase

import (
	"time"

	"github.com/achmadnr21/cinema/internal/domain/contract"
	"github.com/achmadnr21/cinema/internal/domain/dto"
	"github.com/achmadnr21/cinema/internal/utils"
	"github.com/google/uuid"
)

type AuthUsecase struct {
	UserRepo contract.UserInterface
}

func NewAuthUsecase(userRepo contract.UserInterface) *AuthUsecase {
	return &AuthUsecase{
		UserRepo: userRepo,
	}
}

func (uc *AuthUsecase) Register(user *dto.UserCreateRequest) (*dto.User, error) {

	newUser := &dto.User{
		ID:       uuid.New(),
		FullName: user.FullName,
		Email:    user.Email,
		Password: user.Password,
	}
	createdUser, err := uc.UserRepo.Create(newUser)
	if err != nil {
		return nil, &utils.ConflictError{Message: "Email already exists"}
	}
	return createdUser, nil
}

func (uc *AuthUsecase) Login(email, password string) (string, string, error) {
	// Find user by email
	user, err := uc.UserRepo.FindByEmail(email)
	if err != nil {
		return "", "", &utils.NotFoundError{Message: "User Not Found"}
	}

	// check password
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", "", &utils.UnauthorizedError{Message: "Invalid Password"}
	}
	// generate token
	token, err := utils.GenerateAccessToken(user.ID.String())
	if err != nil {
		return "", "", &utils.InternalServerError{Message: "Failed to generate token"}
	}
	// generate refresh token
	refreshToken, err := utils.GenerateRefreshToken(user.ID.String())
	if err != nil {
		return "", "", &utils.InternalServerError{Message: "Failed to generate refresh token"}
	}
	// Return the token and user information
	return token, refreshToken, nil
}

func (uc *AuthUsecase) RefreshToken(refreshToken string) (string, string, error) {
	// Parse Refresh Token
	claims, err := utils.ParseRefreshToken(refreshToken)
	if err != nil {
		return "", "", &utils.UnauthorizedError{Message: "Invalid Refresh Token"}
	}
	// expired check
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return "", "", &utils.UnauthorizedError{Message: "Refresh Token Expired"}
	}
	// Generate new access token
	token, err := utils.GenerateAccessToken(claims.UserId)
	if err != nil {
		return "", "", &utils.InternalServerError{Message: "Failed to generate new access token"}
	}
	// Generate new refresh token
	newRefreshToken, err := utils.GenerateRefreshToken(claims.UserId)
	if err != nil {
		return "", "", &utils.InternalServerError{Message: "Failed to generate new refresh token"}
	}
	return token, newRefreshToken, nil
}
