package cfg

// Redis holds config for redis connection
type Redis struct {
	Addr     string
	Password string
	DB       int
}
