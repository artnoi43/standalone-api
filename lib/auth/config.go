package auth

type Config struct {
	JWTSecret string `mapstructure:"jwt_secret"`
}
