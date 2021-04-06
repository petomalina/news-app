package news

// Config is the news server configuration that can be passed
// in either directly or via an env+default configuration.
// This struct uses https://github.com/sethvargo/go-envconfig
// to process the incoming env variables.
type Config struct {
	Port      string            `env:"PORT,default=8080"`
	Providers map[string]string `env:"PROVIDERS"`
}
