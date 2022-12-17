package structs

type Config struct {
	JWTSecret string `yaml:"JWTSecret"`
	DBPath    string `yaml:"DBPath"`
	Port      string `yaml:"Port"`
}
