// Copyright 2020 NetApp, Inc. All Rights Reserved.

//go:build linux

package utils

import (
	"context"
	"net"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

// TestDetermineNFSPackages tests the determineNFSPackages function
// It checks the packages returned for a variety of distros
//
// -- Doc autogenerated on 2022-05-12 15:01:28.925004 --
func TestDetermineNFSPackages(t *testing.T) {
	log.Debug("Running TestDetermineNFSPackages...")

	tests := map[string]struct {
		host             HostSystem
		expectedPackages []string
		errorExpected    bool
	}{
		"Ubuntu": {
			host: HostSystem{OS: SystemOS{
				Distro:  Ubuntu,
				Version: "18.04",
				Release: "18.04.1",
			}},
			expectedPackages: []string{"nfs-common"},
			errorExpected:    false,
		},
		"Centos": {
			host: HostSystem{OS: SystemOS{
				Distro:  Centos,
				Version: "8",
				Release: "8.1",
			}},
			expectedPackages: []string{"nfs-utils"},
			errorExpected:    false,
		},
		"RHEL": {
			host: HostSystem{OS: SystemOS{
				Distro:  RHEL,
				Version: "8",
				Release: "8.1",
			}},
			expectedPackages: []string{"nfs-utils"},
			errorExpected:    false,
		},
		"Foobar": {
			host: HostSystem{OS: SystemOS{
				Distro:  "Foobar",
				Version: "",
				Release: "",
			}},
			expectedPackages: nil,
			errorExpected:    true,
		},
	}
	for testName, test := range tests {
		t.Logf("Running test case '%s'", testName)

		packages, err := determineNFSPackages(context.Background(), test.host)
		assert.Equal(t, test.errorExpected, err != nil, "Unexpected error value")
		assert.ElementsMatch(t, test.expectedPackages, packages, "Incorrect packages returned")
	}
}

// TestGetPackageManagerForHost tests the getPackageManagerForHost function
// It checks that the correct
//
// -- Doc autogenerated on 2022-05-12 15:01:28.925004 --
func TestGetPackageManagerForHost(t *testing.T) {
	log.Debug("Running TestGetPackageManagerForHost...")

	tests := map[string]struct {
		host          HostSystem
		expectedPM    string
		errorExpected bool
	}{
		"Ubuntu": {
			host: HostSystem{OS: SystemOS{
				Distro:  Ubuntu,
				Version: "18.04",
				Release: "18.04.1",
			}},
			expectedPM:    "apt",
			errorExpected: false,
		},
		"Centos7": {
			host: HostSystem{OS: SystemOS{
				Distro:  Centos,
				Version: "7",
				Release: "7.1",
			}},
			expectedPM:    "yum",
			errorExpected: false,
		},
		"RHEL7": {
			host: HostSystem{OS: SystemOS{
				Distro:  RHEL,
				Version: "7",
				Release: "7.1",
			}},
			expectedPM:    "yum",
			errorExpected: false,
		},
		"Centos8": {
			host: HostSystem{OS: SystemOS{
				Distro:  Centos,
				Version: "8",
				Release: "8.1",
			}},
			expectedPM:    "dnf",
			errorExpected: false,
		},
		"RHEL8": {
			host: HostSystem{OS: SystemOS{
				Distro:  RHEL,
				Version: "8",
				Release: "8.1",
			}},
			expectedPM:    "dnf",
			errorExpected: false,
		},
		"Foobar": {
			host: HostSystem{OS: SystemOS{
				Distro:  "Foobar",
				Version: "",
				Release: "",
			}},
			expectedPM:    "",
			errorExpected: true,
		},
	}
	for testName, test := range tests {
		t.Logf("Running test case '%s'", testName)

		pm, err := getPackageManagerForHost(context.Background(), test.host)
		assert.Equal(t, test.errorExpected, err != nil, "Unexpected error value")
		assert.Equal(t, test.expectedPM, pm, "Incorrect package manager returned")
	}
}
// TestGetIPAddresses tests the getIPAddresses function
// It checks that the returned addresses are not loopback and are global unicast
//
// -- Doc autogenerated on 2022-05-12 15:01:28.925004 --
func TestGetIPAddresses(t *testing.T) {

	addrs, err := getIPAddresses(context.TODO())
	if err != nil {
		t.Error(err)
	}

	assert.Greater(t, len(addrs), 0, "No IP addresses found")

	for _, addr := range addrs {

		parsedAddr := net.ParseIP(strings.Split(addr.String(), "/")[0])
		assert.False(t, parsedAddr.IsLoopback(), "Address is loopback")
		assert.True(t, parsedAddr.IsGlobalUnicast(), "Address is not global unicast")
	}
}

// TestGetIPAddressesExceptingDummyInterfaces tests the getIPAddressesExceptingDummyInterfaces function
// It checks that the function returns at least one address, and that all addresses are global unicast
// and not loopback
//
// -- Doc autogenerated on 2022-05-12 15:01:28.925004 --
func TestGetIPAddressesExceptingDummyInterfaces(t *testing.T) {

	addrs, err := getIPAddressesExceptingDummyInterfaces(context.TODO())
	if err != nil {
		t.Error(err)
	}

	assert.Greater(t, len(addrs), 0, "No IP addresses found")

	for _, addr := range addrs {

		parsedAddr := net.ParseIP(strings.Split(addr.String(), "/")[0])
		assert.False(t, parsedAddr.IsLoopback(), "Address is loopback")
		assert.True(t, parsedAddr.IsGlobalUnicast(), "Address is not global unicast")
	}
}

// TestGetIPAddressesExceptingNondefaultRoutes tests the getIPAddressesExceptingNondefaultRoutes function
// It checks that the function returns at least one IP address and that the IP addresses returned are valid
//
// -- Doc autogenerated on 2022-05-12 15:01:28.925004 --
func TestGetIPAddressesExceptingNondefaultRoutes(t *testing.T) {

	addrs, err := getIPAddressesExceptingNondefaultRoutes(context.TODO())
	if err != nil {
		t.Error(err)
	}

	assert.Greater(t, len(addrs), 0, "No IP addresses found")

	for _, addr := range addrs {

		parsedAddr := net.ParseIP(strings.Split(addr.String(), "/")[0])
		assert.False(t, parsedAddr.IsLoopback(), "Address is loopback")
		assert.True(t, parsedAddr.IsGlobalUnicast(), "Address is not global unicast")
	}
}
