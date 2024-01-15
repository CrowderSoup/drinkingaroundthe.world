package main

import (
	"go.uber.org/fx"
)

func main() {
	bundle := fx.Options()
	app := fx.New(
		bundle,
	)

	app.Run()

	<-app.Done()
}
