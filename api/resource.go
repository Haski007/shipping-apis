package api

import "net/http"

type Input struct {
	SourceAddress string    `json:"source_address"`
	DestAddress   string    `json:"dest_address"`
	BoxDimensions []float64 `json:"box_dimensions"`
}

type Resource interface {
	GetURL() string
	GetAmount(data *Input, client http.Client) (float64, error)
}

type ContentType string

const (
	ApplicationJson ContentType = "application/json"
	ApplicationXML  ContentType = "application/xml"
)

func (typ ContentType) String() string {
	return string(typ)
}
