package services

// HttpClient interface to be implemented by various HTTP clients
type HttpClient interface {
	Get(url string) ([]byte, error)
}
