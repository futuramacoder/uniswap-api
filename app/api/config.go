package api

type Config struct {
	LogLevel         string
	Port             int
	Domain           string
	MetricsNamespace string
	CorsOrigins      []string
	CorsMethods      []string
}
