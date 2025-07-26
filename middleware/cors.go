package middleware

var allowedOrigins = []string{
	"http://localhost:3000",
	"http://localhost:5173",
	"http://localhost:5174",
	"http://Mobius0263.github.io",
	"https://Mobius0263.github.io",
	"http://127.0.0.1:8080",
	"https://pbo-meet-up.vercel.app",
	"https://meet-up.up.railway.app",
}

func GetAllowedOrigins() []string {
	return allowedOrigins
}
