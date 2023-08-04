package app

import "project/xihe-statistics/domain"

type AccessCmd struct {
	URL domain.URL
	IP  domain.IPAddress
}

type AccessCountDTO struct {
	IP    string `json:"ip"`
	Count string `json:"count"`
}
