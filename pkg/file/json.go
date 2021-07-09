package file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ReadJson(path string, dest interface{}) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read file err: %s", err)
	}

	if err := json.Unmarshal(data, &dest); err != nil {
		return fmt.Errorf("json unmarshall file [%s] data err: %s", path, err)
	}

	return err
}
