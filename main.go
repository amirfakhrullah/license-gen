package main

import (
	"fmt"
	"github.com/amirfakhrullah/license-gen/pkg/cli"
	"github.com/amirfakhrullah/license-gen/pkg/helpers"
	"github.com/amirfakhrullah/license-gen/pkg/licenses"
)

func main() {
	lic := licenses.GetLicenseList()

	i := cli.Select(lic)
	name := cli.GetName()
	fmt.Println(i, name)

	licenses.FetchFullLicense(lic[i].Key)
	licContent := licenses.Fill_License(name, helpers.GetYear())
	fmt.Println(licContent)
}
