package _config

import (
	"encoding/json"
	"os"

	_struct "github.com/crlspe/frame-go/util/struct"
)

type configManager struct {
	_useTag     string
	_hideKeyTag string
	_settings   any
}

func NewConfigManager(settings any, useTag string, hideKeyTag string) configManager {
	return configManager{
		_useTag:     useTag,
		_hideKeyTag: hideKeyTag,
		_settings:   settings,
	}
}

func (c *configManager) LoadFromJson(jsonFilePath string) error {
	var jsonFile, err = os.Open(jsonFilePath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	var decoder = json.NewDecoder(jsonFile)
	err = decoder.Decode(c._settings)
	if err != nil {
		return err
	}
	return nil
}

func (c *configManager) SaveToJson(jsonFilePath string, includeHiddenKeyTags ...string) error {
	var jsonFile, err = os.Create(jsonFilePath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	var configData = _struct.ToMap(c._settings, c._useTag, c._hideKeyTag, includeHiddenKeyTags...)

	var encoder = json.NewEncoder(jsonFile)
	encoder.SetIndent("", "    ")
	if err = encoder.Encode(configData); err != nil {
		return err
	}

	return nil
}
