package access_token

import (
	"github.com/Narachii/bookstore_oauth_api/src/utils/errors"
	"strings"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}
type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *errors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
