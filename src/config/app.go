package config

import "os"

type app struct {
	Secret string
}

var App = &app{}

func setupAppConfig() {
	App.Secret = os.Getenv("APP_SECRET")
}
