package configs

type (
	App struct {
		Name    string `env-required:"true" yaml:"name" env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

  HTTP struct {
    Host string `env-required:"true" yaml:"host" env:"HTTP_HOST"`
    PORT int `env-required:"true" yaml:"port" env:"HTTP_PORT"`
  }
)
