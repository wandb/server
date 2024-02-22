package deployer

import (
	"io"
	"net/http"
	"os"
)

const DeployerAPI = "https://deploy.wandb.ai/api/v1/operator/channel"

func GetURL() string {
	if v := os.Getenv("DEPLOYER_CHANNEL_URL"); v != "" {
		return v
	}
	return DeployerAPI
}

func GetChannelSpec(license string) ([]byte, error) {
	url := GetURL()
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if license != "" {
		req.SetBasicAuth("license", license)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
