package config

type App struct {
	Name      string
	Url       string
	Env       string
	JwtSecret string
}

var AppSetting *App
