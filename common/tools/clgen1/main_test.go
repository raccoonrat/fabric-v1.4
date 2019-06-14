/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package main

import (
	"reflect"
	"testing"

	"github.com/hyperledger/fabric/common/tools/clgen/ca"
	"github.com/hyperledger/fabric/common/tools/clgen/kgc"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_getConfig(t *testing.T) {
	tests := []struct {
		name    string
		want    *Config
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("getConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extend(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			extend()
		})
	}
}

func Test_extendPeerOrg(t *testing.T) {
	type args struct {
		orgSpec OrgSpec
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			extendPeerOrg(tt.args.orgSpec)
		})
	}
}

func Test_extendOrdererOrg(t *testing.T) {
	type args struct {
		orgSpec OrgSpec
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			extendOrdererOrg(tt.args.orgSpec)
		})
	}
}

func Test_generate(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generate()
		})
	}
}

func Test_parseTemplate(t *testing.T) {
	type args struct {
		input string
		data  interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseTemplate(tt.args.input, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseTemplateWithDefault(t *testing.T) {
	type args struct {
		input        string
		defaultInput string
		data         interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseTemplateWithDefault(tt.args.input, tt.args.defaultInput, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseTemplateWithDefault() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseTemplateWithDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_renderNodeSpec(t *testing.T) {
	type args struct {
		domain string
		spec   *NodeSpec
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := renderNodeSpec(tt.args.domain, tt.args.spec); (err != nil) != tt.wantErr {
				t.Errorf("renderNodeSpec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_renderOrgSpec(t *testing.T) {
	type args struct {
		orgSpec *OrgSpec
		prefix  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := renderOrgSpec(tt.args.orgSpec, tt.args.prefix); (err != nil) != tt.wantErr {
				t.Errorf("renderOrgSpec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_generatePeerOrg(t *testing.T) {
	type args struct {
		baseDir string
		orgSpec OrgSpec
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generatePeerOrg(tt.args.baseDir, tt.args.orgSpec)
		})
	}
}

func Test_copyAdminCert(t *testing.T) {
	type args struct {
		usersDir      string
		adminCertsDir string
		adminUserName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := copyAdminCert(tt.args.usersDir, tt.args.adminCertsDir, tt.args.adminUserName); (err != nil) != tt.wantErr {
				t.Errorf("copyAdminCert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_generateNodes(t *testing.T) {
	type args struct {
		baseDir  string
		nodes    []NodeSpec
		signKGC  *kgc.KGC
		tlsCA    *ca.CA
		nodeType int
		nodeOUs  bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generateNodes(tt.args.baseDir, tt.args.nodes, tt.args.signKGC, tt.args.tlsCA, tt.args.nodeType, tt.args.nodeOUs)
		})
	}
}

func Test_generateOrdererOrg(t *testing.T) {
	type args struct {
		baseDir string
		orgSpec OrgSpec
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generateOrdererOrg(tt.args.baseDir, tt.args.orgSpec)
		})
	}
}

func Test_copyFile(t *testing.T) {
	type args struct {
		src string
		dst string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := copyFile(tt.args.src, tt.args.dst); (err != nil) != tt.wantErr {
				t.Errorf("copyFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_printVersion(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printVersion()
		})
	}
}

func Test_getCA(t *testing.T) {
	type args struct {
		caDir string
		spec  OrgSpec
		name  string
	}
	tests := []struct {
		name string
		args args
		want *ca.CA
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCA(tt.args.caDir, tt.args.spec, tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getKGC(t *testing.T) {
	type args struct {
		kgcDir string
		spec   OrgSpec
		name   string
	}
	tests := []struct {
		name string
		args args
		want *kgc.KGC
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getKGC(tt.args.kgcDir, tt.args.spec, tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getKGC() = %v, want %v", got, tt.want)
			}
		})
	}
}
