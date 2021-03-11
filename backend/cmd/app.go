package cmd

import "go.uber.org/fx"

func GetApp() *fx.App {
	opts := make([]fx.Option, 0)
	opts = GetProviderOptions()
	opts = append(opts, GetInvokersOptions())
	return fx.New(
		opts...,
	)
}
