package  main

import (
	"alfred/cmd"
	"alfred/config"
	"alfred/logger"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		 config.ConfigProvider,
		 fx.Invoke(
		 	 logger.InitLogger,
             cmd.RunServer,
		 	),
		).Run()

}
