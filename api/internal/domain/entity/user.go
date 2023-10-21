package entity

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/klassmann/cpfcnpj"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	FirstName string    `json:"first_name" bson:"first_name" validate:"required"`
	LastName  string    `json:"last_name"  bson:"last_name"  validate:"required"`
	Nickname  string    `json:"nickname"   bson:"nickname"   validate:"required"`
	Email     string    `json:"email"      bson:"email"      validate:"required,email"`
	Password  string    `json:"password"   bson:"password"`
	CPF       string    `json:"cpf"        bson:"cpf"        validate:"required,min=11,max=11"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

func NewUser(firstName, lastName, nickname, email, password, cpf string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		FirstName: firstName,
		LastName:  lastName,
		Nickname:  nickname,
		Email:     email,
		Password:  string(hash),
		CPF:       cpf,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = user.isValid()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (user *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

func (user *User) isValid() error {
	validate := validator.New()

	err := validate.Struct(user)

	if err != nil {
		return err
	}

	if !cpfcnpj.ValidateCPF(user.CPF) {
		return errors.New("invalid CPF")
	}

	return nil
}

func (user *User) Modify(firstName, lastName, nickname string) error {
	user.FirstName = firstName
	user.LastName = lastName
	user.Nickname = nickname
	user.UpdatedAt = time.Now()

	err := user.isValid()
	if err != nil {
		return err
	}

	return nil
}
