// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ConfigurationManagerClientState undocumented
type ConfigurationManagerClientState int

const (
	// ConfigurationManagerClientStateVUnknown undocumented
	ConfigurationManagerClientStateVUnknown ConfigurationManagerClientState = 0
	// ConfigurationManagerClientStateVInstalled undocumented
	ConfigurationManagerClientStateVInstalled ConfigurationManagerClientState = 1
	// ConfigurationManagerClientStateVHealthy undocumented
	ConfigurationManagerClientStateVHealthy ConfigurationManagerClientState = 7
	// ConfigurationManagerClientStateVInstallFailed undocumented
	ConfigurationManagerClientStateVInstallFailed ConfigurationManagerClientState = 8
	// ConfigurationManagerClientStateVUpdateFailed undocumented
	ConfigurationManagerClientStateVUpdateFailed ConfigurationManagerClientState = 11
	// ConfigurationManagerClientStateVCommunicationError undocumented
	ConfigurationManagerClientStateVCommunicationError ConfigurationManagerClientState = 19
)

// ConfigurationManagerClientStatePUnknown returns a pointer to ConfigurationManagerClientStateVUnknown
func ConfigurationManagerClientStatePUnknown() *ConfigurationManagerClientState {
	v := ConfigurationManagerClientStateVUnknown
	return &v
}

// ConfigurationManagerClientStatePInstalled returns a pointer to ConfigurationManagerClientStateVInstalled
func ConfigurationManagerClientStatePInstalled() *ConfigurationManagerClientState {
	v := ConfigurationManagerClientStateVInstalled
	return &v
}

// ConfigurationManagerClientStatePHealthy returns a pointer to ConfigurationManagerClientStateVHealthy
func ConfigurationManagerClientStatePHealthy() *ConfigurationManagerClientState {
	v := ConfigurationManagerClientStateVHealthy
	return &v
}

// ConfigurationManagerClientStatePInstallFailed returns a pointer to ConfigurationManagerClientStateVInstallFailed
func ConfigurationManagerClientStatePInstallFailed() *ConfigurationManagerClientState {
	v := ConfigurationManagerClientStateVInstallFailed
	return &v
}

// ConfigurationManagerClientStatePUpdateFailed returns a pointer to ConfigurationManagerClientStateVUpdateFailed
func ConfigurationManagerClientStatePUpdateFailed() *ConfigurationManagerClientState {
	v := ConfigurationManagerClientStateVUpdateFailed
	return &v
}

// ConfigurationManagerClientStatePCommunicationError returns a pointer to ConfigurationManagerClientStateVCommunicationError
func ConfigurationManagerClientStatePCommunicationError() *ConfigurationManagerClientState {
	v := ConfigurationManagerClientStateVCommunicationError
	return &v
}