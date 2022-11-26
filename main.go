package main

import (
	"fmt"
	"os"

	"github.com/amirfakhrullah/license-gen/pkg/cli"
	"github.com/amirfakhrullah/license-gen/pkg/helpers"
	"github.com/amirfakhrullah/license-gen/pkg/licenses"
)

func main() {
	lic := licenses.GetLicenseList()

	i := cli.Select(lic)
	name := cli.GetName()

	licenses.FetchFullLicense(lic[i].Key)
	licContent := licenses.Fill_License(name, helpers.GetYear())

	f, osErr := os.Create("LICENSE")
	helpers.HandlePanic(osErr)

	_, writeErr := f.WriteString(licContent)
	helpers.HandlePanic(writeErr)

	fmt.Printf("✅ Successfully added %v\n", lic[i].Name)
}
