// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Tetragon

package kernels

import (
	"fmt"

	"golang.org/x/sys/windows"
)

func GetKernelVersion(kernelVersion, procfs string) (int, string, error) {
	var version int
	var verStr string

	if kernelVersion != "" {
		version = int(KernelStringToNumeric(kernelVersion))
		verStr = kernelVersion
	} else {
		osVersionInfo := windows.RtlGetVersion()

		verStr = fmt.Sprintf("%d.%d.%d",
			osVersionInfo.dwMajorVersion,
			osVersionInfo.dwMinorVersion,
			osVersionInfo.dwBuildNumber)

		version = int(osVersionInfo.dwMajorVersion<<16) + int(osVersionInfo.dwMinorVersion<<8) + int(osVersionInfo.dwBuildNumber)
	}
	return version, verStr, nil
}

func GenericKprobeObjs() (string, string) {
	return "", ""
}

func MinKernelVersion(kernel string) bool {

	runningVersion, _, _ := GetKernelVersion("", "")

	minVersion := int(KernelStringToNumeric(kernel))

	return minVersion <= runningVersion
}
