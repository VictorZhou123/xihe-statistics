package domain

import (
	"errors"
	"project/xihe-statistics/utils"
)

// Site
type Site interface {
	Site() string
}

func NewSite(v string) (Site, error) {
	if v == "" {
		return nil, errors.New("invalid site")
	}

	return site(v), nil
}

type site string

func (r site) Site() string {
	return string(r)
}

// path
type Path interface {
	Path() string
}

func NewPath(v string) (Path, error) {
	if v == "" {
		return nil, errors.New("invalid path")
	}

	return path(v), nil
}

type path string

func (r path) Path() string {
	return string(r)
}

// URL
type URL interface {
	URL() string
}

func NewURL(v string) (URL, error) {
	if !utils.IsValidURL(v) {
		return nil, errors.New("invalid url")
	}

	return url(v), nil
}

type url string

func (r url) URL() string {
	return string(r)
}

// IP Address
type IPAddress interface {
	IPAddress() string
}

func NewIPAddress(v string) (IPAddress, error) {
	if !utils.IsValidIPAddress(v) {
		return nil, errors.New("invalid ip address")
	}

	return ipaddress(v), nil
}

type ipaddress string

func (r ipaddress) IPAddress() string {
	return string(r)
}
