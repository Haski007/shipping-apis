package api

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type ThirdApi struct {
	url         string
	method      string
	contentType ContentType
}

func (rcv ThirdApi) GetURL() string {
	return rcv.url
}

type thirdApiRequest struct {
	SourceAddress string    `json:"consignee"`
	DestAddress   string    `json:"consignor"`
	BoxDimensions []float64 `json:"cartons"`
}

type thirdApiResponse struct {
	Amount float64 `json:"amount" xml:"amount"`
}

func NewThirdApi() Resource {
	return &ThirdApi{
		url:         "http://localhost:3333/",
		method:      "POST",
		contentType: ApplicationXML,
	}
}

func (rcv *ThirdApi) GetAmount(data *Input, client http.Client) (float64, error) {
	body, err := json.Marshal(rcv.encodeThirdApiRequest(data))
	if err != nil {
		return 0, fmt.Errorf("marshall body err: %s", err)
	}

	req, err := http.NewRequest(rcv.method, rcv.url, bytes.NewBuffer(body))
	if err != nil {
		return 0, fmt.Errorf("new request err: %s", err)
	}

	req.Header.Add("Accept", rcv.contentType.String())

	rsp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("do request err: %s", err)
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("response status code not OK: status: %s", rsp.Status)
	}

	var responseData thirdApiResponse
	switch rcv.contentType {
	case ApplicationJson:
		if err := json.NewDecoder(rsp.Body).Decode(&responseData); err != nil {
			return 0, fmt.Errorf("json decode response err: %s", err)
		}
	case ApplicationXML:
		if err := xml.NewDecoder(rsp.Body).Decode(&responseData); err != nil {
			return 0, fmt.Errorf("xml decode response err: %s", err)
		}
	}

	if responseData.Amount == 0 {
		return 0, fmt.Errorf("got amount: 0, so will not count as a real offer")
	}

	return responseData.Amount, nil
}

func (rcv ThirdApi) encodeThirdApiRequest(data *Input) thirdApiRequest {
	return thirdApiRequest{
		SourceAddress: data.SourceAddress,
		DestAddress:   data.DestAddress,
		BoxDimensions: data.BoxDimensions,
	}
}
