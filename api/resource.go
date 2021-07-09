package api

import "net/http"

type Resource interface {
	GetAmount(data interface{}, client *http.Client) (float64, error)
}
