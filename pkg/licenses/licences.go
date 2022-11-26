package licenses

import (
	"encoding/json"
	"github.com/amirfakhrullah/license-gen/pkg/utils"
	"io"
	"net/http"
)

type License struct {
	Key     string `json:"key"`
	Name    string `json:"name"`
	SPDX_ID string `json:"spdx_id"`
	URL     string `json:"url"`
	NODE_ID string `json:"node_id"`
}

func fetchLicenses() []License {
	url := "https://api.github.com/licenses"
	resp, httpErr := http.Get(url)
	utils.HandlePanic(httpErr)

	defer resp.Body.Close()
	body, ioErr := io.ReadAll(resp.Body)
	utils.HandlePanic(ioErr)

	var licenses []License
	unmarshalErr := json.Unmarshal(body, &licenses)
	utils.HandlePanic(unmarshalErr)

	return licenses
}
