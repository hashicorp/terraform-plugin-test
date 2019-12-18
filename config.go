package tftest

import (
	"fmt"
	"os"
)

// Config is used to configure the test helper. In most normal test programs
// the configuration is discovered automatically by an Init* function using
// DiscoverConfig, but this is exposed so that more complex scenarios can be
// implemented by direct configuration.
type Config struct {
	PluginName         string
	SourceDir          string
	TerraformExec      string
	CurrentPluginExec  string
	PreviousPluginExec string
}

// DiscoverConfig uses environment variables and other means to automatically
// discover a reasonable test helper configuration.
func DiscoverConfig(pluginName string, sourceDir string) (*Config, error) {
	var tfExec string
	var err error
	tfVersion := os.Getenv("TF_ACC_TERRAFORM_VERSION")
	if tfVersion == "" {
		tfExec = FindTerraform()
		if tfExec == "" {
			return nil, fmt.Errorf("unable to find 'terraform' executable for testing; either place it in PATH or set TFTEST_TERRAFORM explicitly to a direct executable path")
		}
	} else {
		tfExec, err = InstallTerraform(tfVersion)
		if err != nil {
			return nil, fmt.Errorf("could not install Terraform version %s: %s", tfVersion, err)
		}
	}

	prevExec := os.Getenv("TFTEST_PREVIOUS_EXEC")
	if prevExec != "" {
		if info, err := os.Stat(prevExec); err != nil {
			return nil, fmt.Errorf("TFTEST_PREVIOUS_EXEC of %s cannot be used: %s", prevExec, err)
		} else if info.IsDir() {
			return nil, fmt.Errorf("TFTEST_PREVIOUS_EXEC of %s is directory, not file", prevExec)
		}
	}

	return &Config{
		PluginName:         pluginName,
		SourceDir:          sourceDir,
		TerraformExec:      tfExec,
		CurrentPluginExec:  os.Args[0],
		PreviousPluginExec: os.Getenv("TFTEST_PREVIOUS_EXEC"),
	}, nil
}
