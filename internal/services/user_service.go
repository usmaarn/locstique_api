package services

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/usmaarn/locstique_api/internal/database"
	"github.com/usmaarn/locstique_api/internal/dto"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) CreateUser(requestDto dto.CreateUserDto) (database.User, error) {
	hashedPasswordByte, err := bcrypt.GenerateFromPassword([]byte(requestDto.Password), 12)
	if err != nil {
		return database.User{}, err
	}
	userParams := database.CreateUserParams{
		ID:       uuid.New(),
		Name:     fmt.Sprintf("%s %s", requestDto.FirstName, requestDto.LastName),
		Email:    requestDto.Email,
		Phone:    requestDto.Phone,
		Password: string(hashedPasswordByte),
		Type:     "user",
		Status:   "active",
	}
	return s.db.CreateUser(s.ctx, userParams)
}

func (s *Service) UserExistsByEmail(email string) bool {
	_, err := s.db.FindUserByEmailAddress(s.ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false
		}
		fmt.Println("Error fetching user: ", err)
	}
	return true
}

func (s *Service) GetUserByEmail(email string) (database.User, error) {
	return s.db.FindUserByEmailAddress(s.ctx, email)
}
