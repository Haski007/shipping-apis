package api

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"
)

type SecondApi struct {
	url         string
	method      string
	contentType ContentType
}

func (rcv SecondApi) GetURL() string {
	return rcv.url
}

type secondApiRequest struct {
	SourceAddress string    `json:"consignee"`
	DestAddress   string    `json:"consignor"`
	BoxDimensions []float64 `json:"cartons"`
}

type secondApiResponse struct {
	Amount interface{} `json:"amount" xml:"amount"`
}

func NewSecondApi() Resource {
	return &SecondApi{
		url:         "http://localhost:2222/",
		method:      "POST",
		contentType: ApplicationJson,
	}
}

func (rcv *SecondApi) GetAmount(data *Input, client http.Client) (float64, error) {
	body, err := json.Marshal(rcv.encodeSecondApiRequest(data))
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

	var responseData secondApiResponse
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

	var amount float64
	switch t := responseData.Amount.(type) {
	case float64:
		amount = responseData.Amount.(float64)
	case string:
		amount, err = strconv.ParseFloat(responseData.Amount.(string), 64)
		if err != nil {
			return 0, fmt.Errorf("can't parse float | rsp:%s | err: %s", responseData.Amount, err)
		} else if amount == 0 {
			return 0, fmt.Errorf("got amount: 0, so will not count as a real offer")
		}
	default:
		return 0, fmt.Errorf("not supported value type: %T", t)
	}

	return amount, nil
}

func (rcv SecondApi) encodeSecondApiRequest(data *Input) secondApiRequest {
	return secondApiRequest{
		SourceAddress: data.SourceAddress,
		DestAddress:   data.DestAddress,
		BoxDimensions: data.BoxDimensions,
	}
}
