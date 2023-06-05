package download

import (
	"io"
	"net/http"
	"os"
)

func HTTPDownloadAndSave(url string, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0644)
	return err
}