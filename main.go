package main

import (
	"fmt"
	"os"

	"github.com/amirfakhrullah/license-gen/pkg/cli"
	"github.com/amirfakhrullah/license-gen/pkg/helpers"
	"github.com/amirfakhrullah/license-gen/pkg/licenses"
)

func main() {
	isExist, fileErr := helpers.IsLicenseExist()
	helpers.HandlePanic(&fileErr)

	if isExist {
		toProceed := cli.ConfirmProceed()
		if !toProceed {
			return
		}
	}

	lic := *licenses.GetLicenseList()

	i := cli.Select(&lic)
	name := cli.GetName()
	year := cli.GetYear()

	if len(year) == 0 {
		year = helpers.GetYear()
	}

	licenses.FetchFullLicense(lic[i].Key)
	licContent := licenses.Fill_License(name, year)

	f, osErr := os.Create("LICENSE")
	helpers.HandlePanic(&osErr)

	_, writeErr := f.WriteString(*licContent)
	helpers.HandlePanic(&writeErr)

	fmt.Printf("✅ Successfully added %v\n", lic[i].Name)
}
