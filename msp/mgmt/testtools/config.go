/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package msptesttools

import (
	"os"

	"github.com/hyperledger/fabric/common/util"
	"github.com/hyperledger/fabric/core/config"
	"github.com/hyperledger/fabric/core/config/configtest"
	"github.com/hyperledger/fabric/msp"
	"github.com/hyperledger/fabric/msp/mgmt"
	"github.com/spf13/viper"
)

// LoadTestMSPSetup sets up the local MSP
// and a chain MSP for the default chain
func LoadMSPSetupForTesting() error {
	defer viper.Reset()
	dir, err := configtest.GetDevMspDir()
	if err != nil {
		return err
	}

	if os.Getenv("FABRIC_CFG_PATH") == "" {
		cfgDir, _ := configtest.GetDevConfigDir()
		config.AddConfigPath(nil, cfgDir)
	}
	err = config.InitViper(nil, "core")
	if err != nil {
		return err
	}
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		return nil
	}
	mspType := viper.GetString("peer.localMspType")
	if mspType == "" {
		mspType = msp.ProviderTypeToString(msp.FABRIC)
	}
	conf, err := msp.GetLocalMspConfigWithType(dir, nil, "SampleOrg", mspType)
	if err != nil {
		return err
	}

	err = mgmt.GetLocalMSP().Setup(conf)
	if err != nil {
		return err
	}

	err = mgmt.GetManagerForChain(util.GetTestChainID()).Setup([]msp.MSP{mgmt.GetLocalMSP()})
	if err != nil {
		return err
	}

	return nil
}

// Loads the development local MSP for use in testing.  Not valid for production/runtime context
func LoadDevMsp() error {
	mspDir, err := configtest.GetDevMspDir()
	if err != nil {
		return err
	}
	if os.Getenv("FABRIC_CFG_PATH") == "" {
		cfgDir, _ := configtest.GetDevConfigDir()
		config.AddConfigPath(nil, cfgDir)
	}
	return mgmt.LoadLocalMsp(mspDir, nil, "SampleOrg")
}
