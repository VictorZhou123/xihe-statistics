package repository

import "project/xihe-statistics/domain"

type Access interface {
	Get(domain.Access) (int64, error)
	Upsert(domain.Access) error
}
