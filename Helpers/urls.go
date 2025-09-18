package helpers

import (
	"io"
	"net/http"

	logger "github.com/Chris-Kellett/workflow-manager/Logger"
)

func GetBytesFromURL(guildId string, url string) ([]byte, error) {
	logger.Debug(guildId, "Getting Data from URL: %s", url)
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(guildId, err)
		return nil, err
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			logger.Error(guildId, err)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error(guildId, err)
		return nil, err
	}

	return body, nil
}
