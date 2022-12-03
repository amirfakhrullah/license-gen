package main

import (
	"fmt"

	"github.com/amirfakhrullah/license-gen/pkg/cli"
	"github.com/amirfakhrullah/license-gen/pkg/helpers"
	"github.com/amirfakhrullah/license-gen/pkg/licenses"
)

func main() {
	existedLicenseList, fileErr := helpers.IsLicenseExist()
	helpers.HandlePanic(&fileErr)

	// get confirmation to proceed if there's existing LICENSE
	if len(existedLicenseList) > 0 {
		toProceed := cli.ConfirmProceed(&existedLicenseList)
		if !toProceed {
			return
		}
	}

	lic := *licenses.GetLicenseList()

	i := cli.Select(&lic)
	name := cli.GetName()
	year := cli.GetYear()

	licenses.FetchFullLicense(lic[i].Key)
	licContent := licenses.Fill_License(&name, &year)

	// execute file deletion process for files in existedLicenseList with extensions (.txt, .md, ...)
	delErr := helpers.DeleteExistingLicenseFiles(&existedLicenseList)
	helpers.HandlePanic(&delErr)

	// file writes
	writeErr := helpers.CreateAndWriteLicense(licContent)
	helpers.HandlePanic(&writeErr)

	fmt.Printf("✅ Successfully added %v\n", lic[i].Name)
}
