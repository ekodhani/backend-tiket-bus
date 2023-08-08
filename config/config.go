package config

func Port() string {
	return ":8080"
}

func AllowAkses() []string {
	return []string{
		"http://localhost:3002", "http://localhost/", "http://192.168.1.36:3002", "http://192.168.1.36"}
}
