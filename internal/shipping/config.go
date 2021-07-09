package shipping

import "github.com/Haski007/shipping-apis/pkg/file"

type Config struct {
	InputDataFile string `json:"input_data_file"`
}

func (cfg *Config) ParseFile(path string) error {

	return file.ReadJson(path, cfg)
}
