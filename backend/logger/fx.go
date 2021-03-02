package logger

import "go.uber.org/fx"

var LoggerProvider = fx.Options(
	  fx.Invoke(
	  	 InitLogger,
	  	),
	)
