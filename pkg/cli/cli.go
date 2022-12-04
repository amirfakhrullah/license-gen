package cli

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/amirfakhrullah/license-gen/pkg/helpers"
	"github.com/amirfakhrullah/license-gen/pkg/licenses"
	"github.com/manifoldco/promptui"
)

func Select(licenses *[]licenses.TrimmedLicense) int {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "📌 {{ .Name | cyan }}",
		Inactive: "  {{ .Name }}",
		Selected: "📌 {{ .Name | white }}",
	}

	prompt := promptui.Select{
		Label:     "Which license do you want",
		Items:     *licenses,
		Templates: templates,
		Size:      len(*licenses),
	}

	i, _, promptErr := prompt.Run()
	helpers.HandlePanic(&promptErr)

	return i
}

func GetName() string {
	validate := func(input string) error {
		if len(input) == 0 {
			return errors.New("name is required")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Your name",
		Validate: validate,
	}

	name, err := prompt.Run()
	helpers.HandlePanic(&err)

	return name
}

func GetYear() string {
	validate := func(input string) error {
		if len(input) == 0 {
			return nil
		}
		_, err := time.Parse("2006", input)
		if err != nil {
			return errors.New("please use this format: yyyy")
		}
		return nil
	}

	defaultYear := helpers.GetDefaultYear()

	prompt := promptui.Prompt{
		Label:    fmt.Sprintf("Year (default to %s)", defaultYear),
		Validate: validate,
	}

	year, err := prompt.Run()
	helpers.HandlePanic(&err)

	// if no year input is passed, use defaultYear as value
	if len(year) == 0 {
		return defaultYear
	}
	return year
}

func ConfirmProceed(files *[]string) bool {
	fmt.Printf("⚠️  LICENSE file(s) found in your directory: %v\n", strings.Join(*files, ", "))
	prompt := promptui.Select{
		Label: "Continuing this will erase/replace the file(s) above. Do you wish to proceed?",
		Items: []string{"Proceed", "Cancel"},
	}

	i, _, promptErr := prompt.Run()
	helpers.HandlePanic(&promptErr)

	return i == 0
}

func ToMentionInReadMe() bool  {
	prompt := promptui.Select{
		Label: "Do you want the license to be mentioned at the bottom of your README?",
		Items: []string{"Yes", "No"},
	}

	i, _, promptErr := prompt.Run()
	helpers.HandlePanic(&promptErr)

	return i == 0
}
