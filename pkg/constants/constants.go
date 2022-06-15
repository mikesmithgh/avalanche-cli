// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
package constants

import "time"

const (
	DefaultPerms755 = 0o755

	SubnetEVMReleaseVersion   = "v0.2.3"
	AvalancheGoReleaseVersion = "v1.7.13"

	LatestAvagoReleaseURL = "https://api.github.com/repos/ava-labs/avalanchego/releases/latest"
	SubnetEVMReleaseURL   = "https://api.github.com/repos/ava-labs/subnet-evm/releases/latest"

	ServerRunFile      = "/tmp/gRPCserver.run"
	AvalancheCliBinDir = "bin"
	SidecarSuffix      = "_sidecar.json"
	GenesisSuffix      = "_genesis.json"

	SidecarVersion = "1.0.0"

	RequestTimeout = 3 * time.Minute

	MaxCopy = 2147483648 // 2 GB
)
