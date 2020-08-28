package cmd

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold"
	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/scripts"
	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/tflive"
	"github.com/garyellis/cloud-infrastructure-sdk/pkg/util/projectutil"
)

func DoTfLiveScaffold(projectName string, envNames []string, appName string) error {
	cfg := &input.Config{
		AbsProjectPath: filepath.Join(projectutil.MustGetwd(), projectName),
		ProjectName:    projectName,
	}

	s := &scaffold.Scaffold{}
	err := s.Execute(cfg,
		&scaffold.VersionSh{},
		&scripts.HelpersSh{},
		&tflive.ReadmeMd{},
		&tflive.MetadataYamlSh{},
		&tflive.TfModuleAwsMainTf{AppName: appName},
		&tflive.TfModuleAwsVariablesTf{AppName: appName},
		&tflive.TfModuleAwsLocalsTf{AppName: appName},
		&tflive.TfModuleAwsOutputsTf{AppName: appName},
	)

	if err != nil {
		return err
	}

	// run instances of templates
	for _, envName := range envNames {
		err = s.Execute(cfg,
			&tflive.VarsYaml{EnvName: envName, AppName: appName},
			&tflive.RootTerragruntHcl{EnvName: envName, AppName: appName},
			&tflive.TerragruntHcl{EnvName: envName, AppName: appName},
		)
		if err != nil {
			return err
		}

		/*
			err = s.Execute(cfg,
				&ansiblelive.GroupVarsAllYml{EnvName: envName, AppName: appName},
			)
			if err != nil {
				return err
			}
		*/
	}
	return nil
}
