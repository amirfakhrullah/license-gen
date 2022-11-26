package cli

import (
	"github.com/amirfakhrullah/license-gen/pkg/helpers"
	"github.com/amirfakhrullah/license-gen/pkg/licenses"
	"github.com/manifoldco/promptui"
)

func Select(licenses []licenses.TrimmedLicense) int {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "📌 {{ .Name | cyan }}",
		Inactive: "  {{ .Name }}",
		Selected: "📌 {{ .Name | white }}",
	}

	prompt := promptui.Select{
		Label:     "Which license do you want",
		Items:     licenses,
		Templates: templates,
		Size:      13,
	}

	i, _, promptErr := prompt.Run()
	helpers.HandlePanic(promptErr)

	return i
}

func GetName() string {
	prompt := promptui.Prompt{
		Label: "Your name",
	}

	name, err := prompt.Run()
	helpers.HandlePanic(err)

	return name
}
