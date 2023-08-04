package repositories

import (
	"project/xihe-statistics/domain"
	"project/xihe-statistics/domain/repository"
)

type AccessMapper interface {
	Upsert(AccessInputDO) error
	Get(AccessInputDO) (int64, error)
}

func NewAccessRepository(mapper AccessMapper) repository.Access {
	return access{mapper: mapper}
}

type access struct {
	mapper AccessMapper
}

func (impl access) Get(
	ac domain.Access,
) (int64, error) {
	return impl.mapper.Get(AccessInputDO{
		IP:  ac.IPAddress.IPAddress(),
		URL: ac.URL.URL(),
	})
}

func (impl access) Upsert(
	ac domain.Access,
) error {
	return impl.mapper.Upsert(AccessInputDO{
		IP:  ac.IPAddress.IPAddress(),
		URL: ac.URL.URL(),
	})
}
