package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type FirstApi struct {
	URL    string
	method string
}

type FirstApiRequest struct {
	SourceAddress string    `json:"contact address"`
	DestAddress   string    `json:"warehouse address"`
	BoxDimensions []float64 `json:"package dimensions"`
}

type firstApiResponse struct {
	Total float64 `json:"total"`
}

func NewFirstApi() Resource {
	return &FirstApi{
		URL:    "http://localhost:1111/",
		method: "POST",
	}
}

func (rcv *FirstApi) GetAmount(data interface{}, client *http.Client) (float64, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return 0, fmt.Errorf("marshall body err: %s", err)
	}

	req, err := http.NewRequest(rcv.method, rcv.URL, bytes.NewBuffer(body))
	if err != nil {
		return 0, fmt.Errorf("new request err: %s", err)
	}

	req.Header.Add("Accept", "application/json")

	rsp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("do request err: %s", err)
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("response status code not OK: status: %s", rsp.Status)
	}

	var responseData firstApiResponse
	if err := json.NewDecoder(rsp.Body).Decode(&responseData); err != nil {
		return 0, fmt.Errorf("json decode response err: %s", err)
	}

	return responseData.Total, nil
}
