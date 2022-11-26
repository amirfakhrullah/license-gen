package cli

import (
	"errors"
	"time"

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
	validate := func(input string) error {
		if len(input) == 0 {
			return errors.New("value is required")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Your name",
		Validate: validate,
	}

	name, err := prompt.Run()
	helpers.HandlePanic(err)

	return name
}

func GetYear() string {
	validate := func(input string) error {
		if len(input) == 0 {
			return nil
		}
		_, err := time.Parse("2006", input)
		if err != nil {
			return err
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Year for the license (If skip, will default to this year)",
		Validate: validate,
	}

	year, err := prompt.Run()
	helpers.HandlePanic(err)

	return year
}

func ConfirmProceed() bool {
	prompt := promptui.Select{
		Label: "LICENSE file found in your directory. If you proceed, this will replace the current license",
		Items: []string{"Proceed", "Cancel"},
	}

	i, _, promptErr := prompt.Run()
	helpers.HandlePanic(promptErr)

	return i == 0
}
