package licenses

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/amirfakhrullah/license-gen/pkg/helpers"
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
var selectedLicense FullLicense
var url = "https://api.github.com/licenses"

func init() {
	resp, httpErr := http.Get(url)
	helpers.HandlePanic(&httpErr)

	defer resp.Body.Close()
	body, ioErr := io.ReadAll(resp.Body)
	helpers.HandlePanic(&ioErr)

	var licenses []License
	unmarshalErr := json.Unmarshal(body, &licenses)
	helpers.HandlePanic(&unmarshalErr)

	for _, value := range licenses {
		cachedLicenses = append(cachedLicenses, TrimmedLicense{
			Key:  value.Key,
			Name: value.Name,
		})
	}
}

func FetchFullLicense(key string) {
	resp, httpErr := http.Get(url + "/" + key)
	helpers.HandlePanic(&httpErr)

	defer resp.Body.Close()
	body, ioErr := io.ReadAll(resp.Body)
	helpers.HandlePanic(&ioErr)

	unmarshalErr := json.Unmarshal(body, &selectedLicense)
	helpers.HandlePanic(&unmarshalErr)
}

func GetLicenseList() *[]TrimmedLicense {
	return &cachedLicenses
}

func Fill_License(name *string, year *string) *string {
	nameTagList := []string{"[fullname]", "[name of copyright owner]", "<name of author>"}
	yearTagList := []string{"[year]", "[yyyy]", "<year>"}

	for _, nameTag := range nameTagList {
		selectedLicense.Body = strings.ReplaceAll(selectedLicense.Body, nameTag, *name)
	}
	for _, yearTag := range yearTagList {
		selectedLicense.Body = strings.ReplaceAll(selectedLicense.Body, yearTag, *year)
	}

	return &selectedLicense.Body
}
