package props

import "os"

var (
	Port, ApiKey, Secret string
)

func Setup() {
	ApiKey = os.Getenv("API_KEY_CONVERTER")
	Port = os.Getenv("PORT")
	Secret = os.Getenv("SECRET")
}
