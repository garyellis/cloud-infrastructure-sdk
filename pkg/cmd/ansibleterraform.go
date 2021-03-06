package cmd

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/config"
	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold"
	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/ansibleterraform"
	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/scripts"
	"github.com/garyellis/cloud-infrastructure-sdk/pkg/util/projectutil"
)

// InitAnsibleTerraformScaffold creates or updates the ansible/terraform project
func InitAnsibleTerraformScaffold(configFilePath, terragruntVarsFilePath, cliName, cliVersion, projectName, appName, infraProvider, dcName string, envNames []string, vaultAddr, vaultSSHCa, vaultSSHRole, sshUser, awsRegion, s3BucketName, s3BucketRegion, vsphereServer string) error {
	userCfg := config.NewConfig()
	userCfg.ReadConfigFile(configFilePath)

	if s3BucketName != "" && s3BucketRegion != "" {
		userCfg.AnsibleTerraform.S3BucketNamePrefix = s3BucketName
		userCfg.AnsibleTerraform.S3BucketRegion = s3BucketRegion
	}

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
		&scripts.TerraformHelpersSh{
			TerraformVersion:      userCfg.TerraformHelpers.TerraformVersion,
			TerraformReleasesURL:  userCfg.TerraformHelpers.TerraformReleasesURL,
			TerragruntVersion:     userCfg.TerraformHelpers.TerragruntVersion,
			TerragruntDownloadURL: userCfg.TerraformHelpers.TerragruntDownloadURL,
			TfenvRepoURL:          userCfg.TerraformHelpers.TfenvRepoURL,
		},
		&scripts.DockerHelpersSh{},
		&scripts.PythonHelpersSh{},
		&scripts.VaultHelpersSh{},
		&scripts.TerraformPluginsTxt{PluginUrls: userCfg.TerraformHelpers.TerraformPluginURLs},
		&ansibleterraform.ReadmeMd{},
		&ansibleterraform.AWSMainTf{
			AppName:                           appName,
			TFModuleCloudInitSource:           userCfg.TerraformModuleSources.CloudInit,
			TFModuleVaultApproleSource:        userCfg.TerraformModuleSources.VaultApprole,
			TFModuleSecurityGroupSource:       userCfg.TerraformModuleSources.SecurityGroup,
			TFModuleEc2InstanceSource:         userCfg.TerraformModuleSources.Ec2Instance,
			TFModuleNetworkLoadBalancerSource: userCfg.TerraformModuleSources.NetworkLoadBalancer,
			TFModuleRoute53ZoneSource:         userCfg.TerraformModuleSources.Route53Zone,
		},
		&ansibleterraform.AWSVariablesTf{AppName: appName},
		&ansibleterraform.AWSLocalsTf{AppName: appName},
		&ansibleterraform.AWSOutputsTf{AppName: appName},
		&ansibleterraform.AWSUserdataSh{AppName: appName},
		&ansibleterraform.AwsAnsileInventory{AppName: appName},
		&ansibleterraform.VMwareMainTf{
			AppName:                             appName,
			TFModuleCloudInitSource:             userCfg.TerraformModuleSources.CloudInit,
			TFModuleVaultApproleSource:          userCfg.TerraformModuleSources.VaultApprole,
			TFModuleVsphereVirtualMachineSource: userCfg.TerraformModuleSources.VsphereVirtualMachine,
		},
		&ansibleterraform.VMwareVariablesTf{AppName: appName},
		&ansibleterraform.VMwareOutputsTf{AppName: appName},
		&ansibleterraform.VMwareUserdataSh{AppName: appName},
		&ansibleterraform.VMwareAnsileInventory{AppName: appName},
		&ansibleterraform.RequirementsTxt{},
		&ansibleterraform.AnsibleCfg{},
		&ansibleterraform.RequirementsYml{
			OSRoles:  userCfg.AnsibleTerraform.AnsibleRoleSources.OSRoleSources,
			AppRoles: userCfg.AnsibleTerraform.AnsibleRoleSources.AppRoleSources,
		},
		&ansibleterraform.SiteYml{},
		&ansibleterraform.OSYml{
			AnsibleRoles: userCfg.AnsibleTerraform.AnsibleRoleSources.OSRoleSources,
		},
		&ansibleterraform.ApplicationYml{
			AppName:      appName,
			AnsibleRoles: userCfg.AnsibleTerraform.AnsibleRoleSources.AppRoleSources,
		},
		&ansibleterraform.VarsOverrideYml{},
	)

	// run through instances of templates
	for _, envName := range envNames {
		err = s.Execute(cfg,
			&ansibleterraform.TerragruntBaseHcl{
				EnvName:        envName,
				DCName:         dcName,
				S3BucketName:   userCfg.AnsibleTerraform.S3BucketNamePrefix,
				S3BucketRegion: userCfg.AnsibleTerraform.S3BucketRegion,
				S3KeyPrefix:    "$S3_KEY_PREFIX",
			},

			&ansibleterraform.AnsibleInventoryGroupVarsYml{EnvName: envName, AppName: appName, DCName: dcName},
		)
		if err != nil {
			return err
		}

		// render infrastructure provider specific templates
		if infraProvider == "aws" {
			terragruntVarsCfg := config.NewTerragruntVarsConfig()
			terragruntVarsCfg.ReadConfigFile(terragruntVarsFilePath)

			err = s.Execute(cfg,
				&ansibleterraform.EnvAwsSh{
					EnvName:      envName,
					AppName:      appName,
					DCName:       dcName,
					AWSRegion:    awsRegion,
					VaultAddr:    vaultAddr,
					VaultSSHCa:   vaultSSHCa,
					VaultSSHRole: vaultSSHRole,
					SSHUser:      sshUser,
				},
				&ansibleterraform.TerragruntAwsHcl{EnvName: envName, AppName: appName, DCName: dcName},
				&ansibleterraform.TerragruntAwsVars{
					EnvName:        envName,
					AppName:        appName,
					DCName:         dcName,
					AWSRegion:      awsRegion,
					VaultAddr:      vaultAddr,
					VaultSshCAPath: vaultSSHCa,
					AwsAccountID:   terragruntVarsCfg.AwsAccountID,
					AmiID:          terragruntVarsCfg.AmiID,
					KeyName:        terragruntVarsCfg.KeyName,
					VpcID:          terragruntVarsCfg.VpcID,
					SGAttachments:  terragruntVarsCfg.SGAttachments,
					LBSubnetIDs:    terragruntVarsCfg.LBSubnetIDs,
					NodesSubnetIDs: terragruntVarsCfg.NodesSubnetIDs,
					DNSDomain:      terragruntVarsCfg.DNSDomain,
					DNSZoneID:      terragruntVarsCfg.DNSZoneID,
				},
			)
		} else if infraProvider == "vmware" {
			terragruntVarsCfg := config.NewTerragruntVsphereVarsConfig()
			terragruntVarsCfg.ReadConfigFile(terragruntVarsFilePath)
			err = s.Execute(cfg,
				&ansibleterraform.EnvVmwareSh{
					EnvName:       envName,
					AppName:       appName,
					DCName:        dcName,
					VsphereServer: vsphereServer,
					VaultAddr:     vaultAddr,
					VaultSSHCa:    vaultSSHCa,
					VaultSSHRole:  vaultSSHRole,
				},
				&ansibleterraform.TerragruntVMwareHcl{EnvName: envName, AppName: appName, DCName: dcName},
				&ansibleterraform.TerragruntVMwareVars{
					EnvName:              envName,
					AppName:              appName,
					DCName:               dcName,
					VaultAddr:            vaultAddr,
					VaultSshCAPath:       vaultSSHCa,
					VsphereDCName:        terragruntVarsCfg.VsphereDCName,
					VsphereClusterName:   terragruntVarsCfg.VsphereClusterName,
					VsphereFolder:        terragruntVarsCfg.VsphereFolder,
					VsphereIpv4Gateway:   terragruntVarsCfg.VsphereIpv4Gateway,
					VsphereNetworkNames:  terragruntVarsCfg.VsphereNetworkNames,
					VsphereDataStoreName: terragruntVarsCfg.VsphereDataStoreName,
					VsphereNodes:         terragruntVarsCfg.VsphereNodes,
				},
			)
		}

	}
	if err != nil {
		return err
	}
	return nil
}
