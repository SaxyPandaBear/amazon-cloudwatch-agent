// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

//go:build windows
// +build windows

package paths

import (
	"os"
	"path/filepath"
)

const (
	AgentDir             = "\\Amazon\\AmazonCloudWatchAgent\\"
	JsonDir              = "\\Configs"
	BinaryDir            = "bin"
	TranslatorBinaryName = "config-translator.exe"
	AgentBinaryName      = "amazon-cloudwatch-agent.exe"
	WizardBinaryName     = "amazon-cloudwatch-agent-config-wizard.exe"
	AgentStartName       = "amazon-cloudwatch-agent-ctl.ps1"
)

func init() {
	programFiles := os.Getenv("ProgramFiles")
	var programData string
	if _, ok := os.LookupEnv("ProgramData"); ok {
		programData = os.Getenv("ProgramData")
	} else {
		// Windows 2003
		programData = filepath.Join(os.Getenv("ALLUSERSPROFILE"), "Application Data")
	}

	AgentRootDir := filepath.Join(programFiles, AgentDir)
	AgentConfigDir := filepath.Join(programData, AgentDir)
	JsonConfigPath = filepath.Join(AgentConfigDir, JSON)
	JsonDirPath = filepath.Join(AgentConfigDir, JsonDir)
	EnvConfigPath = filepath.Join(AgentConfigDir, ENV)
	TomlConfigPath = filepath.Join(AgentConfigDir, TOML)
	YamlConfigPath = filepath.Join(AgentConfigDir, YAML)
	CommonConfigPath = filepath.Join(AgentConfigDir, COMMON_CONFIG)
	AgentLogFilePath = filepath.Join(AgentConfigDir, AGENT_LOG_FILE)
	TranslatorBinaryPath = filepath.Join(AgentRootDir, TranslatorBinaryName)
	AgentBinaryPath = filepath.Join(AgentRootDir, AgentBinaryName)
}
