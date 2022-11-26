package main

import (
	"fmt"
	"github.com/amirfakhrullah/license-gen/pkg/cli"
	"github.com/amirfakhrullah/license-gen/pkg/licenses"
)

func main() {
	licences := licenses.GetLicenseList()

	i := cli.Select(licences)
	name := cli.GetName()
	fmt.Println(i, name)
}
