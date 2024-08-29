package config

type App struct {
	Url       string
	Env       string
	JwtSecret string
}

var AppSetting *App
