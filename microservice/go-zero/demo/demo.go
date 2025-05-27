package main

import (
	"flag"
	"fmt"

	"demo/internal/config"
	"demo/internal/handler"
	"demo/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

// 定义配置文件的目录
var configFile = flag.String("f", "etc/demo-api.yaml", "the config file")

func main() {
	flag.Parse() // 解析flag中的参数

	// 根据配置文件目录解析配置文件中的内容
	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 根据配置启动服务器
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	/*
	进行依赖注入 配置、数据库连接、Redis连接
	*/
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
