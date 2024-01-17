package api

const (
	API_URL = "http://checkpoint.phishtank.com/checkurl"
)

type client struct {
	ApiKey string
	UserAgent string
}

func NewClient(apiKey string, userAgent string) client {
	return client {
		ApiKey: apiKey,
		UserAgent: userAgent,
	}
}

func CheckURL(url string) {
	
}
