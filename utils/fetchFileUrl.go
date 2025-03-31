package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func FetchFileUrl(path string) (io.ReadCloser, error) {

	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {

		resp, err := http.Get(path)

		if err != nil {
			return nil, err
		}

		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			return nil, fmt.Errorf(`error: A bad connection %d`, resp.StatusCode)
		}

		return resp.Body, nil
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("neither a file or a url")
	}

	return os.Open(path)
}
