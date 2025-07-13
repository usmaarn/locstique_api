package services

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/usmaarn/locstique_api/internal/database"
	"github.com/usmaarn/locstique_api/internal/dto"
)

func (s *Service) RegisterUser(request dto.CreateUserDto) (database.Token, error) {
	user, err := s.CreateUser(request)
	if err != nil {
		fmt.Println("Error creating user: ", err)
		return database.Token{}, err
	}

	token, err := s.CreateToken(user.ID)
	if err != nil {
		fmt.Println("Error creating user: ", err)
		return database.Token{}, err
	}
	return token, nil
}

func (s *Service) Login(request dto.LoginDto) (database.Token, error) {
	user, err := s.db.FindUserByEmailAddress(s.ctx, request.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return database.Token{}, errors.New("incorrect email or password")
		}
		fmt.Println("Error fetching user: ", err)
		return database.Token{}, errors.New("an error occurred")
	}
}
