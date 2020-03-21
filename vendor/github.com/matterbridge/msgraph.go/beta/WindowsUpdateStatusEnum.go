// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsUpdateStatus undocumented
type WindowsUpdateStatus int

const (
	// WindowsUpdateStatusVUpToDate undocumented
	WindowsUpdateStatusVUpToDate WindowsUpdateStatus = 0
	// WindowsUpdateStatusVPendingInstallation undocumented
	WindowsUpdateStatusVPendingInstallation WindowsUpdateStatus = 1
	// WindowsUpdateStatusVPendingReboot undocumented
	WindowsUpdateStatusVPendingReboot WindowsUpdateStatus = 2
	// WindowsUpdateStatusVFailed undocumented
	WindowsUpdateStatusVFailed WindowsUpdateStatus = 3
)

// WindowsUpdateStatusPUpToDate returns a pointer to WindowsUpdateStatusVUpToDate
func WindowsUpdateStatusPUpToDate() *WindowsUpdateStatus {
	v := WindowsUpdateStatusVUpToDate
	return &v
}

// WindowsUpdateStatusPPendingInstallation returns a pointer to WindowsUpdateStatusVPendingInstallation
func WindowsUpdateStatusPPendingInstallation() *WindowsUpdateStatus {
	v := WindowsUpdateStatusVPendingInstallation
	return &v
}

// WindowsUpdateStatusPPendingReboot returns a pointer to WindowsUpdateStatusVPendingReboot
func WindowsUpdateStatusPPendingReboot() *WindowsUpdateStatus {
	v := WindowsUpdateStatusVPendingReboot
	return &v
}

// WindowsUpdateStatusPFailed returns a pointer to WindowsUpdateStatusVFailed
func WindowsUpdateStatusPFailed() *WindowsUpdateStatus {
	v := WindowsUpdateStatusVFailed
	return &v
}