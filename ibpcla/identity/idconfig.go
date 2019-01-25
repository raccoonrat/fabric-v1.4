//SPDX-License-Identifier: Apache-2.0
package ibpcla

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
)

//IdConfig contains the crypto material to set up an ibpcla identity

type IdConfig struct {
	// Pk is the oartial public key from server to client
	Pk []byte `protobuf:"bytes,1,opt,name=Pk,proto3" json:"Pk,omitempty"`
	// Sk is the final secret key of client
	Sk []byte `protobuf:"bytes,2,opt,name=sk,proto3" json:"sk,omitempty"`
	// Serial is the serial number to identify the PA
	Serial string `protobuf:"bytes,3,opt,name=serial,proto3" json:"serial,omitempty"`
	// OrganizationalUnitIdentifier defines the organizational unit the default signer is in
	OrganizationalUnitIdentifier string `protobuf:"bytes,4,opt,name=organizational_unit_identifier,json=organizationalUnitIdentifier" json:"organizational_unit_identifier,omitempty"`
	// Role defines whether the default signer is admin, member, peer, or client
	Role string `protobuf:"bytes,5,opt,name=role,json=role" json:"role,omitempty"`
	//EnrollmentID contains the enrollment id of the client
	EnrollmentID string `protobuf:"bytes,6,opt,name=enrollment_id,json=enrollmentId" json:"enrollment_id,omitempty"`
	//KRI contains a serilized Key Revocation Information
	KeyRevocationInformation string `protobuf:"bytes,7,opt,name=Key_Revocation_Information,proto3" json:"Key_Revocation_Information,omitempty"`
}

// GetPk returns partial publickey associated with this id config
func (i *IdConfig) GetPk() []byte {
	return i.Pk
}

// GetSk returns final secrect key associated with this id config
func (i *IdConfig) GetSk() []byte {
	return i.Sk
}

// GetSk returns final secrect key associated with this id config
func (i *IdConfig) GetSerial() string {
	return i.Serial
}

// GetOrganizationalUnitIdentifier returns OU of the user associated with this signer config
func (i *IdConfig) GetOrganizationalUnitIdentifier() string {
	return i.OrganizationalUnitIdentifier
}

// GetRole returns true if the user associated with this signer config is an admin, else
// returns role
func (i *IdConfig) GetRole() string {
	return i.Role
}

// GetKeyRevocationInformation returns PRL
func (i *IdConfig) GetKeyRevocationInformation() string {
	return i.KeyRevocationInformation
}

// GetEnrollmentID returns enrollment ID of the user associated with this id config
func (i *IdConfig) GetEnrollmentID() string {
	return i.EnrollmentID
}

// Store stores this ibpcla credential to the given location
func (i *IdConfig) Store(path string) error {
	idConfigBytes, err := json.Marshal(i)
	if err != nil {
		return errors.Wrapf(err, "Failed to marshal IDConfig")
	}
	err = ioutil.WriteFile(path, idConfigBytes, 0644)
	if err != nil {
		return errors.WithMessage(err, "Failed to store the IDconfig")
	}
	return nil
}

// Load loads the ibpcla config from the location specified by the
// IDConfigFile attribute
func (i *IdConfig) Load(path string) error {
	idConfigBytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("No ibpcla config found at %s: %s", path, err.Error())
	}
	err = json.Unmarshal(idConfigBytes, i)
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("Failed to ynmarshal IDConfig bytes from %s", path))
	}
	return nil
}
