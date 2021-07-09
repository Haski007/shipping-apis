package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSecondApi_GetAmount(t *testing.T) {
	serverReturnValue := 502.0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`{"amount" : %.2f}`, serverReturnValue)))
		return
	}))
	defer server.Close()

	api := SecondApi{
		url:         server.URL,
		method:      "POST",
		contentType: ApplicationJson,
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

func TestSecondApi_GetAmountWrongValueType(t *testing.T) {
	var (
		value       = false
		expectedErr = fmt.Errorf("not supported value type: %T", value)
	)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`{"amount" : %v}`, value)))
		return
	}))
	defer server.Close()

	api := SecondApi{
		url:         server.URL,
		method:      "POST",
		contentType: ApplicationJson,
	}
	inputData := Input{
		SourceAddress: "Ireland",
		DestAddress:   "Kiev",
		BoxDimensions: []float64{200.5, 100, 12},
	}
	client := http.Client{}

	amount, err := api.GetAmount(&inputData, client)
	if err.Error() != expectedErr.Error() {
		t.Fatalf("\nexpected err: %s\ngot err: %s", expectedErr, err)
	}

	if amount != 0 {
		t.Fail()
	}
}
