package controller

import (
	"project/xihe-statistics/app"
	"project/xihe-statistics/domain"
)

type AccessRequest struct {
}

func (req *AccessRequest) toCmd(ip string, url string) (cmd app.AccessCmd, err error) {

	if cmd.IP, err = domain.NewIPAddress(ip); err != nil {
		return
	}

	if cmd.URL, err = domain.NewURL(url); err != nil {
		return
	}

	return
}
