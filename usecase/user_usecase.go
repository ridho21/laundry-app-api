package usecase

import (
	"errors"
	"go-enigma-laundry/model"
	"go-enigma-laundry/model/dto/request"
	"go-enigma-laundry/model/dto/response"
	"go-enigma-laundry/repository"
	"go-enigma-laundry/utils/common"
	"go-enigma-laundry/utils/encryption"

	"github.com/google/uuid"
)

type UserUsecase interface {
	FindUserById(id string) (model.User, error)
	CreateUser(payload model.User) (response.RegisterResponseDto, error)
	LoginUser(paylod request.LoginRequestDto) (response.LoginResponseDto,error)
}

type userUsecase struct {
	repo repository.UserRepository
	jwtToken common.JwtToken
}

func (u *userUsecase)FindUserById(id string) (model.User, error){
	return model.User{},nil
}

func (u *userUsecase)CreateUser(payload model.User) (response.RegisterResponseDto, error){
	payload.Id = uuid.NewString()
	hasPass, err := encryption.HashPassword(payload.Password)
	payload.Password = hasPass
	if err != nil {
		return response.RegisterResponseDto{},err
	}
	user,err := u.repo.Create(payload)
	if err != nil {
		return response.RegisterResponseDto{},err
	}
	return response.RegisterResponseDto{
		Id: user.Id,
		FullName: user.FullName,
		Email: user.Email,
		Username: user.Username,
		Role: user.Role,
	},nil
}

func (u *userUsecase)LoginUser(paylod request.LoginRequestDto) (response.LoginResponseDto,error) {
	// Get user by username
	user, err := u.repo.GetByUsername(paylod.Username)
	if err != nil {
		return response.LoginResponseDto{},err
	}
	// Validasi Password
	isValid := encryption.CheckPassword(paylod.Password,user.Password)
	if !isValid {
		return response.LoginResponseDto{},errors.New("Password is invalid")
	}

	// Generate Token
	// loginExpDuration := time.Duration(10) * time.Minute
	// expiredAt := time.Now().Add(loginExpDuration).Unix()
	accessToken, err := u.jwtToken.GenerateTokenJwt(user)
	if err != nil {
		return response.LoginResponseDto{},err
	}
	return response.LoginResponseDto{
		AccessToken: accessToken,
		UserId: user.Id,
	},nil
}

func NewUserUsecase(
	repo repository.UserRepository,
	jwtToken common.JwtToken,
	) UserUsecase {
	return &userUsecase {
		repo : repo,
		jwtToken: jwtToken,
	}
}