package shipping

import (
	"net/http"

	"github.com/Haski007/shipping-apis/api"
)

type Shipping struct {
	Client http.Client

	Resources []api.Resource
	//FirstAPI api.Resource
	//FirstAPI api.Resource
}

func NewShipping() *Shipping {

	return &Shipping{
		Client: http.Client{},
		Resources: []api.Resource{
			api.NewFirstApi(),
			api.NewSecondApi(),
			api.NewThirdApi(),
		},
	}
}
