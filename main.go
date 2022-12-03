package main

import (
	"fmt"
	"os"

	"github.com/amirfakhrullah/license-gen/pkg/cli"
	"github.com/amirfakhrullah/license-gen/pkg/helpers"
	"github.com/amirfakhrullah/license-gen/pkg/licenses"
)

var fileName = "LICENSE"

func main() {
	existedLicenseList, fileErr := helpers.IsLicenseExist()
	helpers.HandlePanic(&fileErr)

	// get confirmation to proceed if there's existing LICENSE
	if len(existedLicenseList) > 0 {
		toProceed := cli.ConfirmProceed()
		if !toProceed {
			return
		}
	}

	lic := *licenses.GetLicenseList()
	defaultYear := helpers.GetYear()

	i := cli.Select(&lic)
	name := cli.GetName()
	year := cli.GetYear(&defaultYear)

	// if no year input is passed, use defaultYear as value
	if len(year) == 0 {
		year = defaultYear
	}

	licenses.FetchFullLicense(lic[i].Key)
	licContent := licenses.Fill_License(&name, &year)

	// execute file deletion process for existedLicenseList with extensions (.txt, .md, ...)
	for _, existedLic := range existedLicenseList {
		if existedLic == fileName {
			continue
		}
		osErr := os.Remove(existedLic)
		helpers.HandlePanic(&osErr)
	}

	f, osErr := os.Create(fileName)
	helpers.HandlePanic(&osErr)

	_, writeErr := f.WriteString(*licContent)
	helpers.HandlePanic(&writeErr)

	fmt.Printf("✅ Successfully added %v\n", lic[i].Name)
}
