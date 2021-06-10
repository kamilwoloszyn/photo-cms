package configs

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

type AuthGatewayConfig struct {
	Host string
	Port string
}

type AuthGateway struct {
	Url string
}

func (config *AuthGatewayConfig) Load() error {
	keyHost, existHost := os.LookupEnv("KRATOS_HOST")
	keyPort, existPort := os.LookupEnv("KRATOS_PORT")

	if !existHost || !existPort {
		if err := godotenv.Load("../.env"); err != nil {
			return errors.Wrap(err, "Loading env variables")
		}

		keyHost, existHost = os.LookupEnv("KRATOS_HOST")
		keyPort, existPort = os.LookupEnv("KRATOS_PORT")

		if !existHost || !existPort {
			return errors.New("Required vars[ auth gateway ] are empty.")
		}
	}
	config.Host = keyHost
	config.Port = keyPort
	return nil
}

func (config *AuthGatewayConfig) TestConnection() (AuthGateway, error) {
	var (
		body []byte
	)
	if config.IsIncomplete() {
		return AuthGateway{}, errors.New("Bad config")
	}

	apiUrl := "http://" + config.Host + ":" + config.Port

	headers := map[string][]string{
		"Accept": []string{"application/json"},
	}

	req, err := http.NewRequest("GET", apiUrl, bytes.NewBuffer(body))
	if err != nil {
		return AuthGateway{}, errors.Wrap(err, fmt.Sprintf("Sending request to %v", apiUrl))
	}
	req.Header = headers
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return AuthGateway{}, errors.Wrap(err, "Processing request")
	}
	if resp.StatusCode != 200 {
		return AuthGateway{}, errors.New("Your auth server probably not working properly.")
	}
	return AuthGateway{Url: apiUrl}, nil
}

func (config *AuthGatewayConfig) IsIncomplete() bool {
	return len(config.Host) == 0 || len(config.Port) == 0
}
