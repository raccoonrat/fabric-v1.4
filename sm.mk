# Copyright IBM Corp All Rights Reserved.
# Copyright London Stock Exchange Group All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0

BUILD_DIR ?= .build

.PHONY: smPlugin
smPlugin:
	go build -o $(BUILD_DIR)/smPlugin.so --buildmode=plugin ./examples/plugins/smPlugin

$(BUILD_DIR)/smPlugin.so: $(BUILD_DIR)/image/buildenv/$(DUMMY)
	@echo "Building smPlugin.so"
	@$(DRUN) \
        -v  $(abspath ./):/opt/gopath/src/github.com/hyperledger/fabric \
	    -v  $(abspath ./.build):/out \
	    -w /opt/gopath/src/github.com/hyperledger/fabric/examples/plugins/smPlugin \
	    hyperledger/fabric-buildenv:latest \
	    go build --buildmode=plugin -o /out/smPlugin.so
