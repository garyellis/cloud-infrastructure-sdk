// Modified from github.com/operator-framework/operator-sdk/blob/master/internal/util/projutil/project_util.go

package projectutil

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
)

type ProjectType = string

const (
	ProjectTypeTerraformLive   ProjectType = "terraform-live"
	ProjectTypeTerraformModule ProjectType = "terraform-module"
	ProjectTypeAnsibleTeraform ProjectType = "ansible/terraform"
	SdkMetadataFile                        = ".cloud-infrastructure-sdk"
)

func MustInProjectRoot() {
	if err := CheckProjectRoot(); err != nil {
		log.Fatal(err)
	}
}

func CheckProjectRoot() error {
	if _, err := os.Stat(SdkMetadataFile); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("must run command in project root dir: project structure requires %s",
				SdkMetadataFile)
		}
		return fmt.Errorf("error while checking if current directory is the project root: %v", err)
	}
	return nil
}

func MustGetwd() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: (%v)", err)
	}
	return wd
}

func getHomeDir() (string, error) {
	hd, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return homedir.Expand(hd)
}
