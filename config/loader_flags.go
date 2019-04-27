package config

import (
	"strings"
	"strconv"
)

// LoadConfigFromFlags fills a configuration object (passed as parameter) with
// values from command-line flags.
func LoadConfigFromFlags(config *Config, flags *StartupFlags) error {
	ushb, err := bucketFromString(flags.UpstreamSecondsHistBucket)
	if err != nil {
		return err
	}

	rshb, err := bucketFromString(flags.ResponseSecondsHistBucket)
	if err != nil {
		return err
	}

	config.Listen = ListenConfig{
		Port:    flags.ListenPort,
		Address: "0.0.0.0",
	}
	config.Namespaces = []NamespaceConfig{
		{
			Format:      flags.Format,
			SourceFiles: flags.Filenames,
			Name:        flags.Namespace,
			UpstreamSecondsHistBucket: ushb, 
			ResponseSecondsHistBucket: rshb,
		},
	}

	return nil
}

func bucketFromString(value string) ([]float64, error) {
	tokens := strings.Split(value, ",")
	bucket := make([]float64, 1)

	for _, numStr := range tokens {
		if numStr != "" {
			num, err := strconv.ParseFloat(numStr, 64)
			
			if err != nil {
				return nil, err
			}
			
			bucket = append(bucket, num)
		}
	}

	return bucket, nil
}
