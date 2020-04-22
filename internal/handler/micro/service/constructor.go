package service

import (
	"github.com/micro/go-micro/v2"
)

func New() micro.Service {
	server := micro.NewService(
		micro.Name("qilin.crm.service"),
		micro.Version("latest"),
		micro.Address(":8081"),
	)

	return server
}
