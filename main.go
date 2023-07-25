package main

import (
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"log"
)

func main() {
	var restConf rest.RestConf
	conf.MustLoad("api/user/etc/user-api.yaml", &restConf)
	s, err := rest.NewServer(restConf)
	if err != nil {
		log.Fatal(err)
	}
}
