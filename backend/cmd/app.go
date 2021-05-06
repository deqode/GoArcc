package cmd

import "go.uber.org/fx"

//GetApp : Get App will return the fx app.
// Fx App contains invokers , providers , lifecycles etc.
// When we start the application using fx app then used  providers  will initialises first.
// After that invoker will invoked automatically.
// Note: Invokers will be executed in the same order.
func GetApp() *fx.App {
	opts := make([]fx.Option, 0)
	opts = GetProviderOptions()
	opts = append(opts, GetInvokersOptions())
	return fx.New(
		opts...,
	)
}
