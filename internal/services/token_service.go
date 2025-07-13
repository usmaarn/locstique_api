package services

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/google/uuid"
	"github.com/usmaarn/locstique_api/internal/database"
	"log"
	"strings"
	"time"
)

func (s *Service) CreateToken(userId uuid.UUID) (database.Token, error) {
	tokenParams := database.CreateTokenParams{
		ID:        GenerateRandomToken(),
		UserID:    userId,
		ExpiresAt: time.Now().UTC().Add(time.Hour * 24),
	}
	return s.db.CreateToken(s.ctx, tokenParams)
}

func GenerateRandomToken() string {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatal("Error generating byte: ", err)
	}
	randStr := hex.EncodeToString(bytes)
	return strings.ToUpper(randStr)
}
