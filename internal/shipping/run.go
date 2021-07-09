package shipping

import (
	"fmt"
	"math"
	"net/http"
	"sync"

	"github.com/Haski007/shipping-apis/api"
	"github.com/Haski007/shipping-apis/pkg/file"
	"github.com/sirupsen/logrus"
)

type Offer struct {
	Amount float64 `json:"amount"`
	APIUrl string  `json:"api_url"`
}

// ConfigFilePath add this file to .gitignore if needed
const ConfigFilePath = "conf.json"

func Run() error {
	var (
		cfg Config

		// ---> Data channels
		out = make(chan Offer)
	)
	if err := cfg.ParseFile(ConfigFilePath); err != nil {
		return fmt.Errorf("[cfg.ParseFile] err: %s", err)
	}

	sh := NewShipping()

	var inputData api.Input
	if err := file.ReadJson(cfg.InputDataFile, &inputData); err != nil {
		return fmt.Errorf("[file.ReadJson] read input data file %s | err: %s", cfg.InputDataFile, err)
	}

	wg := sync.WaitGroup{}
	for _, resource := range sh.Resources {
		wg.Add(1)
		go func(res api.Resource, client http.Client) {
			defer wg.Done()

			amount, err := res.GetAmount(&inputData, client)
			if err != nil {
				logrus.Errorf("[res.GetAmount] api: %s, err: %s", res.GetURL(), err)
				return
			}

			out <- Offer{
				Amount: amount,
				APIUrl: res.GetURL(),
			}
		}(resource, sh.Client)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	bestDeal := Offer{
		Amount: math.MaxFloat64,
		APIUrl: "No offers",
	}
	for result := range out {
		if result.Amount < bestDeal.Amount {
			bestDeal = result
		}
	}

	if bestDeal.Amount == math.MaxFloat64 {
		return fmt.Errorf("there are not fetch info from any API")
	}

	fmt.Printf("Best deals is:\n"+
		"Api url: %s\n"+
		"Amount: %.2f\n",
		bestDeal.APIUrl,
		bestDeal.Amount)

	return nil
}
