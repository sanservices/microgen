package cmd

import (
	"fmt"
	"go/build"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
)

// createAPICmd creates a new api layer
var createAPICmd = &cobra.Command{
	Use:   "api",
	Short: "Creates a new project in current directory.",
	RunE:  createApi,
}

func init() {
	createAPICmd.Flags().IntP("version", "v", 0, "Version number")
	createAPICmd.MarkFlagRequired("version")

	newCmd.AddCommand(createAPICmd)
}

const afterProcessHelp = `
New api was genereated. Please follow these steps:

Add handler to main.go:
- Create a new handler with the api version ex.: "handler := v2.NewHandler(config)"
- Add the handler to the apicore handlers list

Modify the modd.conf file:
- To generate the swagger docs for the new api version you can copy the *prep* steps corresponding to v1.`

func createApi(cmd *cobra.Command, args []string) error {

	inVersion, err := cmd.Flags().GetInt("version")
	if err != nil {
		return err
	}

	// Getting current project's module to use it in generated code.
	modContent, err := ioutil.ReadFile("go.mod")
	if err != nil {
		log.Println("Could not read file")
		return err
	}

	modStr := string(modContent)
	spacesExpected := 2
	module := ""
	for _, c := range modStr {
		if unicode.IsSpace(c) && spacesExpected > 0 {
			spacesExpected--
			continue
		} else if spacesExpected == 2 {
			continue
		} else if spacesExpected == 0 {
			break
		}

		module = fmt.Sprintf("%s%s", module, string(c))
	}

	// Get gopath to look for template folder.
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}

	orgPath := "/src/github.com/san-services/microgen"

	templatePath := gopath + orgPath + "/templates/template_api"

	version := fmt.Sprintf("v%d", inVersion)

	// Generate api package.
	err = filepath.Walk(templatePath, func(path string, info os.FileInfo, err error) error {
		trimedPath := strings.Replace(path, templatePath, "", -1)
		pathToSave := "internal/api/" + version + trimedPath

		if info.IsDir() {
			_ = os.MkdirAll(pathToSave, 0775)
			return nil
		}

		content, err := ioutil.ReadFile(path)
		if err != nil {
			log.Println("Could not read file")
			return err
		}

		contentStr := string(content)

		// Replace goproposal with the module specified by the user.
		file := strings.Replace(contentStr, "goproposal", module, -1)

		// Replace api_version with the correct version number.
		file = strings.Replace(file, "api_version", version, -1)

		return ioutil.WriteFile(pathToSave, []byte(file), 0644)
	})

	swaggerTemplate := gopath + orgPath + "/templates/template_swagger"

	// Generate swagger bundle.
	err = filepath.Walk(swaggerTemplate, func(path string, info os.FileInfo, err error) error {
		trimedPath := strings.Replace(path, swaggerTemplate, "", -1)
		pathToSave := "files/swaggerui/" + version + trimedPath

		if info.IsDir() {
			_ = os.MkdirAll(pathToSave, 0775)
			return nil
		}

		content, err := ioutil.ReadFile(path)
		if err != nil {
			log.Println("Could not read file")
			return err
		}

		return ioutil.WriteFile(pathToSave, content, 0644)
	})

	fmt.Println(afterProcessHelp)
	return err
}
