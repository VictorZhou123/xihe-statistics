package app

import (
	"errors"
	"project/xihe-statistics/domain"
	"project/xihe-statistics/domain/repository"
	"project/xihe-statistics/infrastructure/repositories"
)

const (
	maxAccessAllowed = 5
)

type AccessService interface {
	GetAccess(*AccessCmd) (AccessCountDTO, error)
	AddAccess(*AccessCmd) (string, error)
}

func NewAccessService(
	r repository.Access,
) AccessService {
	return &accessService{
		r: r,
	}
}

type accessService struct {
	r repository.Access
}

func (s *accessService) AddAccess(cmd *AccessCmd) (code string, err error) {
	// get ip access counts
	a := domain.Access{
		URL:       cmd.URL,
		IPAddress: cmd.IP,
	}

	count, err := s.r.Get(a)
	if err != nil {
		return
	}

	// check count if over max allowed
	if count >= maxAccessAllowed {
		code = ErrorAccessOverAllowed
		err = errors.New("access over allowed")

		return
	}

	// add access
	if err = s.r.Upsert(a); err != nil {
		if repositories.IsErrorConcurrentUpdating(err) {
			code = ErrorAccessConcurrentUpdating
			return
		}

		return
	}

	return
}

func (s *accessService) GetAccess(cmd *AccessCmd) (AccessCountDTO, error) {
	return AccessCountDTO{}, nil
}
