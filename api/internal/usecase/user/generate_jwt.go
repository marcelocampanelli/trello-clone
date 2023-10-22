package user

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/marcelocampanelli/trello-clone/internal/domain/gateway"
	"github.com/marcelocampanelli/trello-clone/pkg/utils"
	"time"
)

type CreateJWTInputDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateJWTOutputDTO struct {
	Token string `json:"token"`
}

type CreateJWTUseCase struct {
	UserGateway gateway.UserGateway
}

func NewCreateJWTUseCase(userGateway gateway.UserGateway) *CreateJWTUseCase {
	return &CreateJWTUseCase{UserGateway: userGateway}
}

type CreateJWTUseCaseInterface interface {
	Execute(input *CreateJWTInputDTO) (*CreateJWTOutputDTO, *utils.CustomError)
}

func (useCase *CreateJWTUseCase) Execute(input *CreateJWTInputDTO) (*CreateJWTOutputDTO, *utils.CustomError) {

	jwtExpiresIn := 3000

	user, err := useCase.UserGateway.FindByEmail(input.Email)
	if err != nil {
		return nil, utils.NewCustomRerror(400, "user not found")
	}

	if !user.ValidatePassword(input.Password) {
		return nil, utils.NewCustomRerror(401, "invalid password")
	}

	_, tokenString, _ := jwtauth.New("HS256", []byte("secret"), nil).Encode(map[string]interface{}{
		"user_id": user.ID, "exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})

	return &CreateJWTOutputDTO{
		Token: tokenString,
	}, nil
}
