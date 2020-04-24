package service

import (
	"github.com/micro/go-micro/v2"
)

func New() (micro.Service, error) {
	//b := rabbitmq.NewBroker(
	//	rabbitmq.ExchangeName("crm_exchange"),
	//)
	//
	//if err := b.Init(); err != nil {
	//	return nil, err
	//}
	//
	//if err := b.Connect(); err != nil {
	//	return nil, err
	//}

	server := micro.NewService(
		micro.Name("qilin.crm.service"),
		micro.Version("latest"),
		micro.Address(":8081"),
		//micro.Broker(b),
	)

	return server, nil
}
