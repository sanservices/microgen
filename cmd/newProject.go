package cmd

import (
	"fmt"
	"go/build"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// createProjectCmd is the command to create
// a new micro service using --name and --module as flags.
var createProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "Creates a new project in current directory.",
	RunE:  createProject,
}

func init() {
	createProjectCmd.Flags().StringP("name", "n", "", "New service's name")
	createProjectCmd.Flags().StringP("module", "m", "", "Service's module (example: github.com/foo/bar)")

	createProjectCmd.MarkFlagRequired("name")
	createProjectCmd.MarkFlagRequired("module")

	// Adding createProjectCmd as a top level command.
	newCmd.AddCommand(createProjectCmd)
}

// createProject creates a new micro service project
// in the current directory.
func createProject(cmd *cobra.Command, args []string) error {

	now := time.Now()

	project, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}

	module, err := cmd.Flags().GetString("module")
	if err != nil {
		return err
	}

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}

	// Create the root folder with project's name
	rootFolder := strings.ToLower(strings.Replace(project, " ", "_", -1))
	templatePath := gopath + "/src/github.com/san-services/microgen/templates/template_v1"

	// Loop through the template-project folders and files
	// to generate a new project.
	err = filepath.Walk(templatePath, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		log.Println(info.Name())

		// Remove generator path from info
		// to be able to add it in the current directory.
		trimedPath := strings.Replace(path, templatePath, "", -1)
		pathToSave := rootFolder + trimedPath

		if info.IsDir() {
			_ = os.MkdirAll(pathToSave, 0775)
			return nil
		}

		content, err := ioutil.ReadFile(path)
		if err != nil {
			log.Println("Could not read file")
			return err
		}

		// Replace goproposal module
		// with the one specified by the user.
		contentStr := string(content)
		file := strings.Replace(contentStr, "goproposal", module, -1)
		if info.Name() == "main.go" {
			// Change swagger description name.
			file = strings.Replace(file, "Goproposal", project, -1)
		}

		return ioutil.WriteFile(pathToSave, []byte(file), 0644)
	})

	// Show info about execution time
	elapsed := time.Since(now)
	msg := fmt.Sprintf("The project was created in: %f ms ", float64(elapsed.Nanoseconds())/float64(time.Millisecond))
	fmt.Println(msg)

	return err
}
