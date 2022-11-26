package main

import (
	"fmt"

	"github.com/amirfakhrullah/license-gen/pkg/licenses"
)

func main()  {
	licenseNames := licenses.GetAllLicenseNames()
	fmt.Println(licenseNames)
}