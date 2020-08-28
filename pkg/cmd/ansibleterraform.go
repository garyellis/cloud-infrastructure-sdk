package cmd

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold"
	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/ansibleterraform"
	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/scripts"
	"github.com/garyellis/cloud-infrastructure-sdk/pkg/util/projectutil"
)

// DoAnsibleTerraformScaffold creates or updates the ansible/terraform project
func DoAnsibleTerraformScaffold(cliName, cliVersion, projectName, appName, infraProvider, dcName string, envNames []string) error {
	cfg := &input.Config{
		AbsProjectPath: filepath.Join(projectutil.MustGetwd(), projectName),
		ProjectName:    projectName,
		CliName:        cliName,
		CliVersion:     cliVersion,
	}

	s := &scaffold.Scaffold{}
	err := s.Execute(cfg,
		&ansibleterraform.Makefile{AppName: appName},
		&scaffold.VersionSh{},
		&scripts.HelpersSh{},
		&scripts.AwsHelpersSh{},
		&scripts.TerraformHelpersSh{},
		&scripts.DockerHelpersSh{},
		&scripts.PythonHelpersSh{},
		&ansibleterraform.ReadmeMd{},
		&ansibleterraform.AWSMainTf{AppName: appName},
		&ansibleterraform.AWSVariablesTf{AppName: appName},
		&ansibleterraform.AWSLocalsTf{AppName: appName},
		&ansibleterraform.AWSOutputsTf{AppName: appName},
		&ansibleterraform.AWSUserdataSh{AppName: appName},
		&ansibleterraform.AwsAnsileInventory{AppName: appName},
		&ansibleterraform.VMwareMainTf{AppName: appName},
		&ansibleterraform.VMwareVariablesTf{AppName: appName},
		&ansibleterraform.VMwareOutputsTf{AppName: appName},
		&ansibleterraform.VMwareUserdataSh{AppName: appName},
		&ansibleterraform.VMwareAnsileInventory{AppName: appName},
		&ansibleterraform.AnsibleCfg{},
		&ansibleterraform.RequirementsYml{},
		&ansibleterraform.SiteYml{},
		&ansibleterraform.OSYml{},
		&ansibleterraform.MiddlewareYml{},
	)

	// run through instances of templates
	for _, envName := range envNames {
		err = s.Execute(cfg,
			&ansibleterraform.TerragruntBaseHcl{
				EnvName:        envName,
				DCName:         dcName,
				S3BucketName:   "$S3_BUCKET_NAME",
				S3BucketRegion: "us-west-2",
				S3KeyPrefix:    "$S3_KEY_PREFIX",
			},
			&ansibleterraform.EnvSh{EnvName: envName, AppName: appName, DCName: dcName},
			&ansibleterraform.AnsibleInventoryGroupVarsYml{EnvName: envName, AppName: appName, DCName: dcName},
		)
		if err != nil {
			return err
		}

		// render infrastructure provider specific templates
		if infraProvider == "aws" {
			err = s.Execute(cfg,
				&ansibleterraform.TerragruntAwsHcl{EnvName: envName, AppName: appName, DCName: dcName},
				&ansibleterraform.TerragruntAwsVars{EnvName: envName, AppName: appName, DCName: dcName},
			)
		} else if infraProvider == "vmware" {
			err = s.Execute(cfg,
				&ansibleterraform.TerragruntVMwareHcl{EnvName: envName, AppName: appName, DCName: dcName},
				&ansibleterraform.TerragruntVMwareVars{EnvName: envName, AppName: appName, DCName: dcName},
			)
		}

	}
	if err != nil {
		return err
	}
	return nil
}
