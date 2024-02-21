// certificateManager
// Écrit par J.F.Gratton (jean-francois@famillegratton.net)
// addRemoveEnvFile.go, jfgratton : 2024-02-20

package env

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func RemoveEnvFile(envfiles []string) error {
	for _, envfile := range envfiles {
		if !strings.HasSuffix(envfile, ".json") {
			envfile += ".json"
		}
		if err := os.Remove(filepath.Join(os.Getenv("HOME"), ".config", "JFG", "nxrmutils", envfile)); err != nil {
			return err
		}
		if err := os.Remove(filepath.Join(os.Getenv("HOME"), ".config", "JFG", "nxrmutils", envfile)); err != nil {
			return err
		}
		fmt.Printf("%s removed succesfully\n", envfile)
	}
	return nil
}

func AddEnvFile(envfile string) error {
	var env NXRMinfo
	var err error

	if envfile == "" {
		envfile = EnvConfigFile
	}
	if !strings.HasSuffix(envfile, ".json") {
		envfile += ".json"
	}

	env = prompt4EnvironmentValues()

	if err = env.SaveEnvironmentFile(envfile); err != nil {
		return err
	}
	return nil
}

func prompt4EnvironmentValues() NXRMinfo {
	var env NXRMinfo

	return env
}

func LoadEnvironmentFile() (NXRMinfo, error) {
	var payload NXRMinfo
	var err error

	if !strings.HasSuffix(EnvConfigFile, ".json") {
		EnvConfigFile += ".json"
	}
	rcFile := filepath.Join(os.Getenv("HOME"), ".config", "JFG", "nxrmuploader", EnvConfigFile)
	jFile, err := os.ReadFile(rcFile)
	if err != nil {
		return NXRMinfo{}, err
	}
	err = json.Unmarshal(jFile, &payload)
	if err != nil {
		return NXRMinfo{}, err
	} else {
		return payload, nil
	}
}

// Save the above structure into a JSON file in the user's .config/certificatemanager directory
func (e NXRMinfo) SaveEnvironmentFile(outputfile string) error {
	if outputfile == "" {
		outputfile = EnvConfigFile
	}
	jStream, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return err
	}
	rcFile := filepath.Join(os.Getenv("HOME"), ".config", "JFG", "nxrmutils", outputfile)
	err = os.WriteFile(rcFile, jStream, 0600)

	return err
}
