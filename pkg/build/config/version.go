package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Metadata struct {
	GrafanaVersion string `json:"version,omitempty"`
	ReleaseMode    string `json:"releaseMode,omitempty"`
	GrabplVersion  string `json:"grabplVersion,omitempty"`
}

type PluginSignature struct {
	Sign      bool `json:"sign,omitempty"`
	AdminSign bool `json:"adminSign,omitempty"`
}

type Docker struct {
	ShouldSave    bool           `json:"shouldSave,omitempty"`
	Architectures []Architecture `json:"archs,omitempty"`
}

// Version represents the "version.json" that defines all of the different variables used to build Grafana
type Version struct {
	Variants                  []Variant       `json:"variants,omitempty"`
	PluginSignature           PluginSignature `json:"pluginSignature,omitempty"`
	Docker                    Docker          `json:"docker,omitempty"`
	PackagesBucket            string          `json:"packagesBucket,omitempty"`
	PackagesBucketEnterprise2 string          `json:"packagesBucketEnterprise2,omitempty"`
	CDNAssetsBucket           string          `json:"CDNAssetsBucket,omitempty"`
	CDNAssetsDir              string          `json:"CDNAssetsDir,omitempty"`
	StorybookBucket           string          `json:"storybookBucket,omitempty"`
	StorybookSrcDir           string          `json:"storybookSrcDir,omitempty"`
}

// Versions is a map of versions. Each key of the Versions map is an event that uses the the config as the value for that key.
// For example, the 'pull_request' key will have data in it that might cause Grafana to be built differently in a pull request,
// than the way it will be built in 'main'
type VersionMap map[string]Version

// GetMetadata attempts to read the JSON file located at 'path' and decode it as a Metadata{} type.
func GetMetadata(path string) (*Metadata, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return DecodeMetadata(file)
}

// DecodeMetadata decodes the data in the io.Reader 'r' as Metadata.
func DecodeMetadata(r io.Reader) (*Metadata, error) {
	m := &Metadata{}
	if err := json.NewDecoder(r).Decode(m); err != nil {
		return nil, err
	}

	return m, nil
}

// GetVersions reads the embedded config.json and decodes it.
func GetVersion(mode string) (*Version, error) {
	if v, ok := Versions[mode]; ok {
		return &v, nil
	}

	return nil, fmt.Errorf("mode not found in version list")
}
