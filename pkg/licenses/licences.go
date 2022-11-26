package licenses

import (
	"encoding/json"
	"github.com/amirfakhrullah/license-gen/pkg/helpers"
	"io"
	"net/http"
)

type License struct {
	Key     string `json:"key"`
	Name    string `json:"name"`
	Spdx_Id string `json:"spdx_id"`
	Url     string `json:"url"`
	Node_Id string `json:"node_id"`
}

type FullLicense struct {
	Key            string   `json:"key"`
	Name           string   `json:"name"`
	Spdx_Id        string   `json:"spdx_id"`
	Url            string   `json:"url"`
	Node_Id        string   `json:"node_id"`
	Html_Url       string   `json:"html_url"`
	Description    string   `json:"description"`
	Implementation string   `json:"implementation"`
	Permissions    []string `json:"permissions"`
	Conditions     []string `json:"conditions"`
	Limitations    []string `json:"limitations"`
	Body           string   `json:"body"`
	Featured       bool     `json:"featured"`
}

type TrimmedLicense struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

var cachedLicenses []TrimmedLicense
var url = "https://api.github.com/licenses"

func init() {
	resp, httpErr := http.Get(url)
	helpers.HandlePanic(httpErr)

	defer resp.Body.Close()
	body, ioErr := io.ReadAll(resp.Body)
	helpers.HandlePanic(ioErr)

	var licenses []License
	unmarshalErr := json.Unmarshal(body, &licenses)
	helpers.HandlePanic(unmarshalErr)

	for _, value := range licenses {
		cachedLicenses = append(cachedLicenses, TrimmedLicense{
			Key:  value.Key,
			Name: value.Name,
		})
	}
}

func FetchFullLicense(key string) FullLicense {
	resp, httpErr := http.Get(url + "/" + key)
	helpers.HandlePanic(httpErr)

	defer resp.Body.Close()
	body, ioErr := io.ReadAll(resp.Body)
	helpers.HandlePanic(ioErr)

	var license FullLicense
	unmarshalErr := json.Unmarshal(body, &license)
	helpers.HandlePanic(unmarshalErr)

	return license
}

func GetLicenseList() []TrimmedLicense {
	return cachedLicenses
}
