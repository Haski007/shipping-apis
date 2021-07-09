package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestThirdApi_GetAmount(t *testing.T) {
	serverReturnValue := 502.0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(
			`<ResponseC>
  <amount>%.2f</amount>
</ResponseC>`, serverReturnValue)))
		return
	}))
	defer server.Close()

	api := ThirdApi{
		url:         server.URL,
		method:      "POST",
		contentType: ApplicationXML,
	}
	inputData := Input{
		SourceAddress: "Ireland",
		DestAddress:   "Kiev",
		BoxDimensions: []float64{200.5, 100, 12},
	}
	client := http.Client{}

	amount, err := api.GetAmount(&inputData, client)
	if err != nil {
		t.Error(err)
	}

	if amount == 0 {
		t.Fail()
	}
}
