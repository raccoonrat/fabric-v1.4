// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msp_config.proto

package msp

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// MSPConfig collects all the configuration information for
// an MSP. The Config field should be unmarshalled in a way
// that depends on the Type
type MSPConfig struct {
	// Type holds the type of the MSP; the default one would
	// be of type FABRIC implementing an X.509 based provider
	Type int32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	// Config is MSP dependent configuration info
	Config               []byte   `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MSPConfig) Reset()         { *m = MSPConfig{} }
func (m *MSPConfig) String() string { return proto.CompactTextString(m) }
func (*MSPConfig) ProtoMessage()    {}
func (*MSPConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1cbdb818a978d24, []int{0}
}

func (m *MSPConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MSPConfig.Unmarshal(m, b)
}
func (m *MSPConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MSPConfig.Marshal(b, m, deterministic)
}
func (m *MSPConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MSPConfig.Merge(m, src)
}
func (m *MSPConfig) XXX_Size() int {
	return xxx_messageInfo_MSPConfig.Size(m)
}
func (m *MSPConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MSPConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MSPConfig proto.InternalMessageInfo

func (m *MSPConfig) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *MSPConfig) GetConfig() []byte {
	if m != nil {
		return m.Config
	}
	return nil
}

// FabricMSPConfig collects all the configuration information for
// a Fabric MSP.
// Here we assume a default certificate validation policy, where
// any certificate signed by any of the listed rootCA certs would
// be considered as valid under this MSP.
// This MSP may or may not come with a signing identity. If it does,
// it can also issue signing identities. If it does not, it can only
// be used to validate and verify certificates.
type FabricMSPConfig struct {
	// Name holds the identifier of the MSP; MSP identifier
	// is chosen by the application that governs this MSP.
	// For example, and assuming the default implementation of MSP,
	// that is X.509-based and considers a single Issuer,
	// this can refer to the Subject OU field or the Issuer OU field.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// List of root certificates trusted by this MSP
	// they are used upon certificate validation (see
	// comment for IntermediateCerts below)
	RootCerts [][]byte `protobuf:"bytes,2,rep,name=root_certs,json=rootCerts,proto3" json:"root_certs,omitempty"`
	// List of intermediate certificates trusted by this MSP;
	// they are used upon certificate validation as follows:
	// validation attempts to build a path from the certificate
	// to be validated (which is at one end of the path) and
	// one of the certs in the RootCerts field (which is at
	// the other end of the path). If the path is longer than
	// 2, certificates in the middle are searched within the
	// IntermediateCerts pool
	IntermediateCerts [][]byte `protobuf:"bytes,3,rep,name=intermediate_certs,json=intermediateCerts,proto3" json:"intermediate_certs,omitempty"`
	// Identity denoting the administrator of this MSP
	Admins [][]byte `protobuf:"bytes,4,rep,name=admins,proto3" json:"admins,omitempty"`
	// Identity revocation list
	RevocationList [][]byte `protobuf:"bytes,5,rep,name=revocation_list,json=revocationList,proto3" json:"revocation_list,omitempty"`
	// SigningIdentity holds information on the signing identity
	// this peer is to use, and which is to be imported by the
	// MSP defined before
	SigningIdentity *SigningIdentityInfo `protobuf:"bytes,6,opt,name=signing_identity,json=signingIdentity,proto3" json:"signing_identity,omitempty"`
	// OrganizationalUnitIdentifiers holds one or more
	// fabric organizational unit identifiers that belong to
	// this MSP configuration
	OrganizationalUnitIdentifiers []*FabricOUIdentifier `protobuf:"bytes,7,rep,name=organizational_unit_identifiers,json=organizationalUnitIdentifiers,proto3" json:"organizational_unit_identifiers,omitempty"`
	// FabricCryptoConfig contains the configuration parameters
	// for the cryptographic algorithms used by this MSP
	CryptoConfig *FabricCryptoConfig `protobuf:"bytes,8,opt,name=crypto_config,json=cryptoConfig,proto3" json:"crypto_config,omitempty"`
	// List of TLS root certificates trusted by this MSP.
	// They are returned by GetTLSRootCerts.
	TlsRootCerts [][]byte `protobuf:"bytes,9,rep,name=tls_root_certs,json=tlsRootCerts,proto3" json:"tls_root_certs,omitempty"`
	// List of TLS intermediate certificates trusted by this MSP;
	// They are returned by GetTLSIntermediateCerts.
	TlsIntermediateCerts [][]byte `protobuf:"bytes,10,rep,name=tls_intermediate_certs,json=tlsIntermediateCerts,proto3" json:"tls_intermediate_certs,omitempty"`
	// fabric_node_ous contains the configuration to distinguish clients from peers from orderers
	// based on the OUs.
	FabricNodeOus        *FabricNodeOUs `protobuf:"bytes,11,opt,name=fabric_node_ous,json=fabricNodeOus,proto3" json:"fabric_node_ous,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FabricMSPConfig) Reset()         { *m = FabricMSPConfig{} }
func (m *FabricMSPConfig) String() string { return proto.CompactTextString(m) }
func (*FabricMSPConfig) ProtoMessage()    {}
func (*FabricMSPConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1cbdb818a978d24, []int{1}
}

func (m *FabricMSPConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FabricMSPConfig.Unmarshal(m, b)
}
func (m *FabricMSPConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FabricMSPConfig.Marshal(b, m, deterministic)
}
func (m *FabricMSPConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FabricMSPConfig.Merge(m, src)
}
func (m *FabricMSPConfig) XXX_Size() int {
	return xxx_messageInfo_FabricMSPConfig.Size(m)
}
func (m *FabricMSPConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FabricMSPConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FabricMSPConfig proto.InternalMessageInfo

func (m *FabricMSPConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FabricMSPConfig) GetRootCerts() [][]byte {
	if m != nil {
		return m.RootCerts
	}
	return nil
}

func (m *FabricMSPConfig) GetIntermediateCerts() [][]byte {
	if m != nil {
		return m.IntermediateCerts
	}
	return nil
}

func (m *FabricMSPConfig) GetAdmins() [][]byte {
	if m != nil {
		return m.Admins
	}
	return nil
}

func (m *FabricMSPConfig) GetRevocationList() [][]byte {
	if m != nil {
		return m.RevocationList
	}
	return nil
}

func (m *FabricMSPConfig) GetSigningIdentity() *SigningIdentityInfo {
	if m != nil {
		return m.SigningIdentity
	}
	return nil
}

func (m *FabricMSPConfig) GetOrganizationalUnitIdentifiers() []*FabricOUIdentifier {
	if m != nil {
		return m.OrganizationalUnitIdentifiers
	}
	return nil
}

func (m *FabricMSPConfig) GetCryptoConfig() *FabricCryptoConfig {
	if m != nil {
		return m.CryptoConfig
	}
	return nil
}

func (m *FabricMSPConfig) GetTlsRootCerts() [][]byte {
	if m != nil {
		return m.TlsRootCerts
	}
	return nil
}

func (m *FabricMSPConfig) GetTlsIntermediateCerts() [][]byte {
	if m != nil {
		return m.TlsIntermediateCerts
	}
	return nil
}

func (m *FabricMSPConfig) GetFabricNodeOus() *FabricNodeOUs {
	if m != nil {
		return m.FabricNodeOus
	}
	return nil
}

// FabricCryptoConfig contains configuration parameters
// for the cryptographic algorithms used by the MSP
// this configuration refers to
type FabricCryptoConfig struct {
	// SignatureHashFamily is a string representing the hash family to be used
	// during sign and verify operations.
	// Allowed values are "SHA2" and "SHA3".
	SignatureHashFamily string `protobuf:"bytes,1,opt,name=signature_hash_family,json=signatureHashFamily,proto3" json:"signature_hash_family,omitempty"`
	// IdentityIdentifierHashFunction is a string representing the hash function
	// to be used during the computation of the identity identifier of an MSP identity.
	// Allowed values are "SHA256", "SHA384" and "SHA3_256", "SHA3_384".
	IdentityIdentifierHashFunction string   `protobuf:"bytes,2,opt,name=identity_identifier_hash_function,json=identityIdentifierHashFunction,proto3" json:"identity_identifier_hash_function,omitempty"`
	XXX_NoUnkeyedLiteral           struct{} `json:"-"`
	XXX_unrecognized               []byte   `json:"-"`
	XXX_sizecache                  int32    `json:"-"`
}

func (m *FabricCryptoConfig) Reset()         { *m = FabricCryptoConfig{} }
func (m *FabricCryptoConfig) String() string { return proto.CompactTextString(m) }
func (*FabricCryptoConfig) ProtoMessage()    {}
func (*FabricCryptoConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1cbdb818a978d24, []int{2}
}

func (m *FabricCryptoConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FabricCryptoConfig.Unmarshal(m, b)
}
func (m *FabricCryptoConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FabricCryptoConfig.Marshal(b, m, deterministic)
}
func (m *FabricCryptoConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FabricCryptoConfig.Merge(m, src)
}
func (m *FabricCryptoConfig) XXX_Size() int {
	return xxx_messageInfo_FabricCryptoConfig.Size(m)
}
func (m *FabricCryptoConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FabricCryptoConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FabricCryptoConfig proto.InternalMessageInfo

func (m *FabricCryptoConfig) GetSignatureHashFamily() string {
	if m != nil {
		return m.SignatureHashFamily
	}
	return ""
}

func (m *FabricCryptoConfig) GetIdentityIdentifierHashFunction() string {
	if m != nil {
		return m.IdentityIdentifierHashFunction
	}
	return ""
}

// IdemixMSPConfig collects all the configuration information for
// an Idemix MSP.
type IdemixMSPConfig struct {
	// Name holds the identifier of the MSP
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// ipk represents the (serialized) issuer public key
	Ipk []byte `protobuf:"bytes,2,opt,name=ipk,proto3" json:"ipk,omitempty"`
	// signer may contain crypto material to configure a default signer
	Signer *IdemixMSPSignerConfig `protobuf:"bytes,3,opt,name=signer,proto3" json:"signer,omitempty"`
	// revocation_pk is the public key used for revocation of credentials
	RevocationPk []byte `protobuf:"bytes,4,opt,name=revocation_pk,json=revocationPk,proto3" json:"revocation_pk,omitempty"`
	// epoch represents the current epoch (time interval) used for revocation
	Epoch                int64    `protobuf:"varint,5,opt,name=epoch,proto3" json:"epoch,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdemixMSPConfig) Reset()         { *m = IdemixMSPConfig{} }
func (m *IdemixMSPConfig) String() string { return proto.CompactTextString(m) }
func (*IdemixMSPConfig) ProtoMessage()    {}
func (*IdemixMSPConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1cbdb818a978d24, []int{3}
}

func (m *IdemixMSPConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdemixMSPConfig.Unmarshal(m, b)
}
func (m *IdemixMSPConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdemixMSPConfig.Marshal(b, m, deterministic)
}
func (m *IdemixMSPConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdemixMSPConfig.Merge(m, src)
}
func (m *IdemixMSPConfig) XXX_Size() int {
	return xxx_messageInfo_IdemixMSPConfig.Size(m)
}
func (m *IdemixMSPConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_IdemixMSPConfig.DiscardUnknown(m)
}

var xxx_messageInfo_IdemixMSPConfig proto.InternalMessageInfo

func (m *IdemixMSPConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *IdemixMSPConfig) GetIpk() []byte {
	if m != nil {
		return m.Ipk
	}
	return nil
}

func (m *IdemixMSPConfig) GetSigner() *IdemixMSPSignerConfig {
	if m != nil {
		return m.Signer
	}
	return nil
}

func (m *IdemixMSPConfig) GetRevocationPk() []byte {
	if m != nil {
		return m.RevocationPk
	}
	return nil
}

func (m *IdemixMSPConfig) GetEpoch() int64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

// IdemixMSPSIgnerConfig contains the crypto material to set up an idemix signing identity
type IdemixMSPSignerConfig struct {
	// cred represents the serialized idemix credential of the default signer
	Cred []byte `protobuf:"bytes,1,opt,name=cred,proto3" json:"cred,omitempty"`
	// sk is the secret key of the default signer, corresponding to credential Cred
	Sk []byte `protobuf:"bytes,2,opt,name=sk,proto3" json:"sk,omitempty"`
	// organizational_unit_identifier defines the organizational unit the default signer is in
	OrganizationalUnitIdentifier string `protobuf:"bytes,3,opt,name=organizational_unit_identifier,json=organizationalUnitIdentifier,proto3" json:"organizational_unit_identifier,omitempty"`
	// role defines whether the default signer is admin, peer, member or client
	Role int32 `protobuf:"varint,4,opt,name=role,proto3" json:"role,omitempty"`
	// enrollment_id contains the enrollment id of this signer
	EnrollmentId string `protobuf:"bytes,5,opt,name=enrollment_id,json=enrollmentId,proto3" json:"enrollment_id,omitempty"`
	// credential_revocation_information contains a serialized CredentialRevocationInformation
	CredentialRevocationInformation []byte   `protobuf:"bytes,6,opt,name=credential_revocation_information,json=credentialRevocationInformation,proto3" json:"credential_revocation_information,omitempty"`
	XXX_NoUnkeyedLiteral            struct{} `json:"-"`
	XXX_unrecognized                []byte   `json:"-"`
	XXX_sizecache                   int32    `json:"-"`
}

func (m *IdemixMSPSignerConfig) Reset()         { *m = IdemixMSPSignerConfig{} }
func (m *IdemixMSPSignerConfig) String() string { return proto.CompactTextString(m) }
func (*IdemixMSPSignerConfig) ProtoMessage()    {}
func (*IdemixMSPSignerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1cbdb818a978d24, []int{4}
}

func (m *IdemixMSPSignerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdemixMSPSignerConfig.Unmarshal(m, b)
}
func (m *IdemixMSPSignerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdemixMSPSignerConfig.Marshal(b, m, deterministic)
}
func (m *IdemixMSPSignerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdemixMSPSignerConfig.Merge(m, src)
}
func (m *IdemixMSPSignerConfig) XXX_Size() int {
	return xxx_messageInfo_IdemixMSPSignerConfig.Size(m)
}
func (m *IdemixMSPSignerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_IdemixMSPSignerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_IdemixMSPSignerConfig proto.InternalMessageInfo

func (m *IdemixMSPSignerConfig) GetCred() []byte {
	if m != nil {
		return m.Cred
	}
	return nil
}

func (m *IdemixMSPSignerConfig) GetSk() []byte {
	if m != nil {
		return m.Sk
	}
	return nil
}

func (m *IdemixMSPSignerConfig) GetOrganizationalUnitIdentifier() string {
	if m != nil {
		return m.OrganizationalUnitIdentifier
	}
	return ""
}

func (m *IdemixMSPSignerConfig) GetRole() int32 {
	if m != nil {
		return m.Role
	}
	return 0
}

func (m *IdemixMSPSignerConfig) GetEnrollmentId() string {
	if m != nil {
		return m.EnrollmentId
	}
	return ""
}

func (m *IdemixMSPSignerConfig) GetCredentialRevocationInformation() []byte {
	if m != nil {
		return m.CredentialRevocationInformation
	}
	return nil
}

// SigningIdentityInfo represents the configuration information
// related to the signing identity the peer is to use for generating
// endorsements
type SigningIdentityInfo struct {
	// PublicSigner carries the public information of the signing
	// identity. For an X.509 provider this would be represented by
	// an X.509 certificate
	PublicSigner []byte `protobuf:"bytes,1,opt,name=public_signer,json=publicSigner,proto3" json:"public_signer,omitempty"`
	// PrivateSigner denotes a reference to the private key of the
	// peer's signing identity
	PrivateSigner        *KeyInfo `protobuf:"bytes,2,opt,name=private_signer,json=privateSigner,proto3" json:"private_signer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SigningIdentityInfo) Reset()         { *m = SigningIdentityInfo{} }
func (m *SigningIdentityInfo) String() string { return proto.CompactTextString(m) }
func (*SigningIdentityInfo) ProtoMessage()    {}
func (*SigningIdentityInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1cbdb818a978d24, []int{5}
}

func (m *SigningIdentityInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SigningIdentityInfo.Unmarshal(m, b)
}
func (m *SigningIdentityInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SigningIdentityInfo.Marshal(b, m, deterministic)
}
func (m *SigningIdentityInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SigningIdentityInfo.Merge(m, src)
}
func (m *SigningIdentityInfo) XXX_Size() int {
	return xxx_messageInfo_SigningIdentityInfo.Size(m)
}
func (m *SigningIdentityInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SigningIdentityInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SigningIdentityInfo proto.InternalMessageInfo

func (m *SigningIdentityInfo) GetPublicSigner() []byte {
	if m != nil {
		return m.PublicSigner
	}
	return nil
}

func (m *SigningIdentityInfo) GetPrivateSigner() *KeyInfo {
	if m != nil {
		return m.PrivateSigner
	}
	return nil
}

// KeyInfo represents a (secret) key that is either already stored
// in the bccsp/keystore or key material to be imported to the
// bccsp key-store. In later versions it may contain also a
// keystore identifier
type KeyInfo struct {
	// Identifier of the key inside the default keystore; this for
	// the case of Software BCCSP as well as the HSM BCCSP would be
	// the SKI of the key
	KeyIdentifier string `protobuf:"bytes,1,opt,name=key_identifier,json=keyIdentifier,proto3" json:"key_identifier,omitempty"`
	// KeyMaterial (optional) for the key to be imported; this is
	// properly encoded key bytes, prefixed by the type of the key
	KeyMaterial          []byte   `protobuf:"bytes,2,opt,name=key_material,json=keyMaterial,proto3" json:"key_material,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyInfo) Reset()         { *m = KeyInfo{} }
func (m *KeyInfo) String() string { return proto.CompactTextString(m) }
func (*KeyInfo) ProtoMessage()    {}
func (*KeyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1cbdb818a978d24, []int{6}
}

func (m *KeyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyInfo.Unmarshal(m, b)
}
func (m *KeyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyInfo.Marshal(b, m, deterministic)
}
func (m *KeyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyInfo.Merge(m, src)
}
func (m *KeyInfo) XXX_Size() int {
	return xxx_messageInfo_KeyInfo.Size(m)
}
func (m *KeyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_KeyInfo proto.InternalMessageInfo

func (m *KeyInfo) GetKeyIdentifier() string {
	if m != nil {
		return m.KeyIdentifier
	}
	return ""
}

func (m *KeyInfo) GetKeyMaterial() []byte {
	if m != nil {
		return m.KeyMaterial
	}
	return nil
}

// FabricOUIdentifier represents an organizational unit and
// its related chain of trust identifier.
type FabricOUIdentifier struct {
	// Certificate represents the second certificate in a certification chain.
	// (Notice that the first certificate in a certification chain is supposed
	// to be the certificate of an identity).
	// It must correspond to the certificate of root or intermediate CA
	// recognized by the MSP this message belongs to.
	// Starting from this certificate, a certification chain is computed
	// and bound to the OrganizationUnitIdentifier specified
	Certificate []byte `protobuf:"bytes,1,opt,name=certificate,proto3" json:"certificate,omitempty"`
	// OrganizationUnitIdentifier defines the organizational unit under the
	// MSP identified with MSPIdentifier
	OrganizationalUnitIdentifier string   `protobuf:"bytes,2,opt,name=organizational_unit_identifier,json=organizationalUnitIdentifier,proto3" json:"organizational_unit_identifier,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	XXX_unrecognized             []byte   `json:"-"`
	XXX_sizecache                int32    `json:"-"`
}

func (m *FabricOUIdentifier) Reset()         { *m = FabricOUIdentifier{} }
func (m *FabricOUIdentifier) String() string { return proto.CompactTextString(m) }
func (*FabricOUIdentifier) ProtoMessage()    {}
func (*FabricOUIdentifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1cbdb818a978d24, []int{7}
}

func (m *FabricOUIdentifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FabricOUIdentifier.Unmarshal(m, b)
}
func (m *FabricOUIdentifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FabricOUIdentifier.Marshal(b, m, deterministic)
}
func (m *FabricOUIdentifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FabricOUIdentifier.Merge(m, src)
}
func (m *FabricOUIdentifier) XXX_Size() int {
	return xxx_messageInfo_FabricOUIdentifier.Size(m)
}
func (m *FabricOUIdentifier) XXX_DiscardUnknown() {
	xxx_messageInfo_FabricOUIdentifier.DiscardUnknown(m)
}

var xxx_messageInfo_FabricOUIdentifier proto.InternalMessageInfo

func (m *FabricOUIdentifier) GetCertificate() []byte {
	if m != nil {
		return m.Certificate
	}
	return nil
}

func (m *FabricOUIdentifier) GetOrganizationalUnitIdentifier() string {
	if m != nil {
		return m.OrganizationalUnitIdentifier
	}
	return ""
}

// FabricNodeOUs contains configuration to tell apart clients from peers from orderers
// based on OUs. If NodeOUs recognition is enabled then an msp identity
// that does not contain any of the specified OU will be considered invalid.
type FabricNodeOUs struct {
	// If true then an msp identity that does not contain any of the specified OU will be considered invalid.
	Enable bool `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	// OU Identifier of the clients
	ClientOuIdentifier *FabricOUIdentifier `protobuf:"bytes,2,opt,name=client_ou_identifier,json=clientOuIdentifier,proto3" json:"client_ou_identifier,omitempty"`
	// OU Identifier of the peers
	PeerOuIdentifier     *FabricOUIdentifier `protobuf:"bytes,3,opt,name=peer_ou_identifier,json=peerOuIdentifier,proto3" json:"peer_ou_identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *FabricNodeOUs) Reset()         { *m = FabricNodeOUs{} }
func (m *FabricNodeOUs) String() string { return proto.CompactTextString(m) }
func (*FabricNodeOUs) ProtoMessage()    {}
func (*FabricNodeOUs) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1cbdb818a978d24, []int{8}
}

func (m *FabricNodeOUs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FabricNodeOUs.Unmarshal(m, b)
}
func (m *FabricNodeOUs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FabricNodeOUs.Marshal(b, m, deterministic)
}
func (m *FabricNodeOUs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FabricNodeOUs.Merge(m, src)
}
func (m *FabricNodeOUs) XXX_Size() int {
	return xxx_messageInfo_FabricNodeOUs.Size(m)
}
func (m *FabricNodeOUs) XXX_DiscardUnknown() {
	xxx_messageInfo_FabricNodeOUs.DiscardUnknown(m)
}

var xxx_messageInfo_FabricNodeOUs proto.InternalMessageInfo

func (m *FabricNodeOUs) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *FabricNodeOUs) GetClientOuIdentifier() *FabricOUIdentifier {
	if m != nil {
		return m.ClientOuIdentifier
	}
	return nil
}

func (m *FabricNodeOUs) GetPeerOuIdentifier() *FabricOUIdentifier {
	if m != nil {
		return m.PeerOuIdentifier
	}
	return nil
}

// CLMSPConfig collects all the configuration information for
// an IBPCLA MSP.
type CLMSPConfig struct {
	// Name holds the identifier of the MSP
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// List of root kgc pubs trusted by this MSP
	// they are used upon certificate validation (see
	// comment for IntermediateCerts below)
	KGCPubs [][]byte `protobuf:"bytes,2,rep,name=KGC_pubs,json=KGCPubs,proto3" json:"KGC_pubs,omitempty"`
	// Identity denoting the administrator of this MSP
	Admins [][]byte `protobuf:"bytes,3,rep,name=admins,proto3" json:"admins,omitempty"`
	// Identity revocation list
	RevocationList [][]byte `protobuf:"bytes,4,rep,name=revocation_list,json=revocationList,proto3" json:"revocation_list,omitempty"`
	// SigningIdentity holds information on the signing identity
	// this peer is to use, and which is to be imported by the
	// MSP defined before
	CLSigningIdentity *CLMSPSignerConfig `protobuf:"bytes,5,opt,name=CLSigningIdentity,proto3" json:"CLSigningIdentity,omitempty"`
	// OrganizationalUnitIdentifiers holds one or more
	// fabric organizational unit identifiers that belong to
	// this MSP configuration
	OrganizationalUnitIdentifiers []*FabricOUIdentifier `protobuf:"bytes,6,rep,name=organizational_unit_identifiers,json=organizationalUnitIdentifiers,proto3" json:"organizational_unit_identifiers,omitempty"`
	// FabricCryptoConfig contains the configuration parameters
	// for the cryptographic algorithms used by this MSP
	CryptoConfig *FabricCryptoConfig `protobuf:"bytes,7,opt,name=crypto_config,json=cryptoConfig,proto3" json:"crypto_config,omitempty"`
	// List of TLS root certificates trusted by this MSP.
	// They are returned by GetTLSRootCerts.
	TlsRootCerts [][]byte `protobuf:"bytes,8,rep,name=tls_root_certs,json=tlsRootCerts,proto3" json:"tls_root_certs,omitempty"`
	// List of TLS intermediate certificates trusted by this MSP;
	// They are returned by GetTLSIntermediateCerts.
	TlsIntermediateCerts [][]byte `protobuf:"bytes,9,rep,name=tls_intermediate_certs,json=tlsIntermediateCerts,proto3" json:"tls_intermediate_certs,omitempty"`
	// fabric_node_ous contains the configuration to distinguish clients from peers from orderers
	// based on the OUs.
	FabricNodeOus        *FabricNodeOUs `protobuf:"bytes,10,opt,name=fabric_node_ous,json=fabricNodeOus,proto3" json:"fabric_node_ous,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CLMSPConfig) Reset()         { *m = CLMSPConfig{} }
func (m *CLMSPConfig) String() string { return proto.CompactTextString(m) }
func (*CLMSPConfig) ProtoMessage()    {}
func (*CLMSPConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1cbdb818a978d24, []int{9}
}

func (m *CLMSPConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CLMSPConfig.Unmarshal(m, b)
}
func (m *CLMSPConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CLMSPConfig.Marshal(b, m, deterministic)
}
func (m *CLMSPConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CLMSPConfig.Merge(m, src)
}
func (m *CLMSPConfig) XXX_Size() int {
	return xxx_messageInfo_CLMSPConfig.Size(m)
}
func (m *CLMSPConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CLMSPConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CLMSPConfig proto.InternalMessageInfo

func (m *CLMSPConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CLMSPConfig) GetKGCPubs() [][]byte {
	if m != nil {
		return m.KGCPubs
	}
	return nil
}

func (m *CLMSPConfig) GetAdmins() [][]byte {
	if m != nil {
		return m.Admins
	}
	return nil
}

func (m *CLMSPConfig) GetRevocationList() [][]byte {
	if m != nil {
		return m.RevocationList
	}
	return nil
}

func (m *CLMSPConfig) GetCLSigningIdentity() *CLMSPSignerConfig {
	if m != nil {
		return m.CLSigningIdentity
	}
	return nil
}

func (m *CLMSPConfig) GetOrganizationalUnitIdentifiers() []*FabricOUIdentifier {
	if m != nil {
		return m.OrganizationalUnitIdentifiers
	}
	return nil
}

func (m *CLMSPConfig) GetCryptoConfig() *FabricCryptoConfig {
	if m != nil {
		return m.CryptoConfig
	}
	return nil
}

func (m *CLMSPConfig) GetTlsRootCerts() [][]byte {
	if m != nil {
		return m.TlsRootCerts
	}
	return nil
}

func (m *CLMSPConfig) GetTlsIntermediateCerts() [][]byte {
	if m != nil {
		return m.TlsIntermediateCerts
	}
	return nil
}

func (m *CLMSPConfig) GetFabricNodeOus() *FabricNodeOUs {
	if m != nil {
		return m.FabricNodeOus
	}
	return nil
}

// IdemixMSPSIgnerConfig contains the crypto material to set up an idemix signing identity
type CLMSPSignerConfig struct {
	// sk is the secret key of the default signer, corresponding to credential Cred
	Sk []byte `protobuf:"bytes,1,opt,name=sk,proto3" json:"sk,omitempty"`
	// PA is the vice identity of the signer
	PA []byte `protobuf:"bytes,2,opt,name=PA,proto3" json:"PA,omitempty"`
	// ID is the identity of the signer
	ID                   string   `protobuf:"bytes,3,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CLMSPSignerConfig) Reset()         { *m = CLMSPSignerConfig{} }
func (m *CLMSPSignerConfig) String() string { return proto.CompactTextString(m) }
func (*CLMSPSignerConfig) ProtoMessage()    {}
func (*CLMSPSignerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1cbdb818a978d24, []int{10}
}

func (m *CLMSPSignerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CLMSPSignerConfig.Unmarshal(m, b)
}
func (m *CLMSPSignerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CLMSPSignerConfig.Marshal(b, m, deterministic)
}
func (m *CLMSPSignerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CLMSPSignerConfig.Merge(m, src)
}
func (m *CLMSPSignerConfig) XXX_Size() int {
	return xxx_messageInfo_CLMSPSignerConfig.Size(m)
}
func (m *CLMSPSignerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CLMSPSignerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CLMSPSignerConfig proto.InternalMessageInfo

func (m *CLMSPSignerConfig) GetSk() []byte {
	if m != nil {
		return m.Sk
	}
	return nil
}

func (m *CLMSPSignerConfig) GetPA() []byte {
	if m != nil {
		return m.PA
	}
	return nil
}

func (m *CLMSPSignerConfig) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func init() {
	proto.RegisterType((*MSPConfig)(nil), "msp.MSPConfig")
	proto.RegisterType((*FabricMSPConfig)(nil), "msp.FabricMSPConfig")
	proto.RegisterType((*FabricCryptoConfig)(nil), "msp.FabricCryptoConfig")
	proto.RegisterType((*IdemixMSPConfig)(nil), "msp.IdemixMSPConfig")
	proto.RegisterType((*IdemixMSPSignerConfig)(nil), "msp.IdemixMSPSignerConfig")
	proto.RegisterType((*SigningIdentityInfo)(nil), "msp.SigningIdentityInfo")
	proto.RegisterType((*KeyInfo)(nil), "msp.KeyInfo")
	proto.RegisterType((*FabricOUIdentifier)(nil), "msp.FabricOUIdentifier")
	proto.RegisterType((*FabricNodeOUs)(nil), "msp.FabricNodeOUs")
	proto.RegisterType((*CLMSPConfig)(nil), "msp.CLMSPConfig")
	proto.RegisterType((*CLMSPSignerConfig)(nil), "msp.CLMSPSignerConfig")
}

func init() { proto.RegisterFile("msp_config.proto", fileDescriptor_a1cbdb818a978d24) }

var fileDescriptor_a1cbdb818a978d24 = []byte{
	// 960 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x4f, 0x6f, 0xe3, 0x44,
	0x14, 0x57, 0xe2, 0x24, 0x6d, 0x5e, 0x9c, 0x3f, 0x3b, 0xdb, 0x2d, 0x06, 0xb1, 0xbb, 0xa9, 0x01,
	0x91, 0x0b, 0xa9, 0x94, 0x45, 0x42, 0x42, 0x5c, 0x76, 0x53, 0x76, 0x31, 0x6d, 0x69, 0xe4, 0xaa,
	0x17, 0x2e, 0x96, 0x63, 0x4f, 0x92, 0x51, 0x6c, 0x8f, 0x35, 0x33, 0x5e, 0x11, 0xc4, 0x99, 0x2f,
	0xc0, 0x77, 0xe0, 0xcc, 0x95, 0xaf, 0xc6, 0x09, 0xcd, 0x9f, 0x24, 0x4e, 0x5b, 0x65, 0x5b, 0xc1,
	0x6d, 0xe6, 0xbd, 0xdf, 0x7b, 0x7e, 0xf3, 0x7b, 0xff, 0x0c, 0xbd, 0x94, 0xe7, 0x41, 0x44, 0xb3,
	0x19, 0x99, 0x0f, 0x73, 0x46, 0x05, 0x45, 0x56, 0xca, 0x73, 0xf7, 0x1b, 0x68, 0x5e, 0x5e, 0x4f,
	0xc6, 0x4a, 0x8e, 0x10, 0xd4, 0xc4, 0x2a, 0xc7, 0x4e, 0xa5, 0x5f, 0x19, 0xd4, 0x7d, 0x75, 0x46,
	0xc7, 0xd0, 0xd0, 0x56, 0x4e, 0xb5, 0x5f, 0x19, 0xd8, 0xbe, 0xb9, 0xb9, 0x7f, 0xd5, 0xa0, 0xfb,
	0x36, 0x9c, 0x32, 0x12, 0xed, 0xd8, 0x67, 0x61, 0xaa, 0xed, 0x9b, 0xbe, 0x3a, 0xa3, 0xe7, 0x00,
	0x8c, 0x52, 0x11, 0x44, 0x98, 0x09, 0xee, 0x54, 0xfb, 0xd6, 0xc0, 0xf6, 0x9b, 0x52, 0x32, 0x96,
	0x02, 0xf4, 0x15, 0x20, 0x92, 0x09, 0xcc, 0x52, 0x1c, 0x93, 0x50, 0x60, 0x03, 0xb3, 0x14, 0xec,
	0x49, 0x59, 0xa3, 0xe1, 0xc7, 0xd0, 0x08, 0xe3, 0x94, 0x64, 0xdc, 0xa9, 0x29, 0x88, 0xb9, 0xa1,
	0x2f, 0xa1, 0xcb, 0xf0, 0x7b, 0x1a, 0x85, 0x82, 0xd0, 0x2c, 0x48, 0x08, 0x17, 0x4e, 0x5d, 0x01,
	0x3a, 0x5b, 0xf1, 0x05, 0xe1, 0x02, 0x8d, 0xa1, 0xc7, 0xc9, 0x3c, 0x23, 0xd9, 0x3c, 0x20, 0x31,
	0xce, 0x04, 0x11, 0x2b, 0xa7, 0xd1, 0xaf, 0x0c, 0x5a, 0x23, 0x67, 0x98, 0xf2, 0x7c, 0x78, 0xad,
	0x95, 0x9e, 0xd1, 0x79, 0xd9, 0x8c, 0xfa, 0x5d, 0xbe, 0x2b, 0x44, 0x01, 0xbc, 0xa4, 0x6c, 0x1e,
	0x66, 0xe4, 0x57, 0xe5, 0x38, 0x4c, 0x82, 0x22, 0x23, 0xc2, 0x38, 0x9c, 0x11, 0xcc, 0xb8, 0x73,
	0xd0, 0xb7, 0x06, 0xad, 0xd1, 0x47, 0xca, 0xa7, 0xa6, 0xe9, 0xea, 0xc6, 0xdb, 0xe8, 0xfd, 0xe7,
	0xbb, 0xf6, 0x37, 0x19, 0x11, 0x5b, 0x2d, 0x47, 0xdf, 0x41, 0x3b, 0x62, 0xab, 0x5c, 0x50, 0x93,
	0x31, 0xe7, 0x50, 0x85, 0x58, 0x76, 0x37, 0x56, 0x7a, 0x4d, 0xbc, 0x6f, 0x47, 0xa5, 0x1b, 0xfa,
	0x1c, 0x3a, 0x22, 0xe1, 0x41, 0x89, 0xf6, 0xa6, 0xe2, 0xc2, 0x16, 0x09, 0xf7, 0x37, 0xcc, 0x7f,
	0x0d, 0xc7, 0x12, 0x75, 0x0f, 0xfb, 0xa0, 0xd0, 0x47, 0x22, 0xe1, 0xde, 0x9d, 0x04, 0x7c, 0x0b,
	0xdd, 0x99, 0xfa, 0x7e, 0x90, 0xd1, 0x18, 0x07, 0xb4, 0xe0, 0x4e, 0x4b, 0xc5, 0x86, 0x4a, 0xb1,
	0xfd, 0x44, 0x63, 0x7c, 0x75, 0xc3, 0xfd, 0xf6, 0x6c, 0x7b, 0x2d, 0xb8, 0xfb, 0x47, 0x05, 0xd0,
	0xdd, 0xe0, 0xd1, 0x08, 0x9e, 0x49, 0x82, 0x43, 0x51, 0x30, 0x1c, 0x2c, 0x42, 0xbe, 0x08, 0x66,
	0x61, 0x4a, 0x92, 0x95, 0x29, 0xa3, 0xa7, 0x1b, 0xe5, 0x0f, 0x21, 0x5f, 0xbc, 0x55, 0x2a, 0xe4,
	0xc1, 0xc9, 0x3a, 0x7d, 0x25, 0xda, 0x8d, 0x75, 0x91, 0x45, 0x92, 0x56, 0x55, 0xb0, 0x4d, 0xff,
	0xc5, 0x1a, 0xb8, 0x25, 0x58, 0x39, 0x32, 0x28, 0xf7, 0xcf, 0x0a, 0x74, 0xbd, 0x18, 0xa7, 0xe4,
	0x97, 0xfd, 0x85, 0xdc, 0x03, 0x8b, 0xe4, 0x4b, 0xd3, 0x05, 0xf2, 0x88, 0x46, 0xd0, 0x90, 0xb1,
	0x61, 0xe6, 0x58, 0x8a, 0x82, 0x4f, 0x14, 0x05, 0x1b, 0x5f, 0xd7, 0x4a, 0x67, 0x32, 0x64, 0x90,
	0xe8, 0x33, 0x68, 0x97, 0x0a, 0x35, 0x5f, 0x3a, 0x35, 0xe5, 0xcf, 0xde, 0x0a, 0x27, 0x4b, 0x74,
	0x04, 0x75, 0x9c, 0xd3, 0x68, 0xe1, 0xd4, 0xfb, 0x95, 0x81, 0xe5, 0xeb, 0x8b, 0xfb, 0x7b, 0x15,
	0x9e, 0xdd, 0xeb, 0x5c, 0x86, 0x1b, 0x31, 0x1c, 0xab, 0x70, 0x6d, 0x5f, 0x9d, 0x51, 0x07, 0xaa,
	0x7c, 0x1d, 0x6d, 0x95, 0x2f, 0xd1, 0x19, 0xbc, 0xd8, 0x5f, 0xb3, 0xea, 0x11, 0x4d, 0xff, 0xd3,
	0x7d, 0x95, 0x29, 0xbf, 0xc4, 0x68, 0x82, 0x55, 0xd4, 0x75, 0x5f, 0x9d, 0xe5, 0x93, 0x70, 0xc6,
	0x68, 0x92, 0xa4, 0x38, 0x93, 0x0e, 0x55, 0xd4, 0x4d, 0xdf, 0xde, 0x0a, 0xbd, 0x18, 0xfd, 0x08,
	0x27, 0x32, 0x2c, 0xe9, 0x28, 0x4c, 0x82, 0x12, 0x05, 0x24, 0x9b, 0x51, 0x96, 0xaa, 0xb3, 0x6a,
	0x44, 0xdb, 0x7f, 0xb9, 0x05, 0xfa, 0x1b, 0x9c, 0xb7, 0x85, 0xb9, 0x14, 0x9e, 0xde, 0xd3, 0xa6,
	0x32, 0x8e, 0xbc, 0x98, 0x26, 0x24, 0x0a, 0x4c, 0x56, 0x34, 0x1d, 0xb6, 0x16, 0x6a, 0xc2, 0xd0,
	0x2b, 0xe8, 0xe4, 0x8c, 0xbc, 0x97, 0xc5, 0x6e, 0x50, 0x55, 0x95, 0x3b, 0x5b, 0xe5, 0xee, 0x1c,
	0xeb, 0x8e, 0x6f, 0x1b, 0x8c, 0x36, 0x72, 0xaf, 0xe1, 0xc0, 0x68, 0xd0, 0x17, 0xd0, 0x59, 0xe2,
	0x72, 0xcd, 0x99, 0x1a, 0x69, 0x2f, 0x71, 0xa9, 0xc0, 0xd0, 0x09, 0xd8, 0x12, 0x96, 0x86, 0x02,
	0x33, 0x12, 0x26, 0x26, 0x0f, 0xad, 0x25, 0x5e, 0x5d, 0x1a, 0x91, 0xfb, 0xdb, 0xba, 0x19, 0xca,
	0x83, 0x01, 0xf5, 0xa1, 0x25, 0x9b, 0x90, 0xcc, 0x48, 0x14, 0x0a, 0x6c, 0x9e, 0x50, 0x16, 0x3d,
	0x20, 0x91, 0xd5, 0x0f, 0x27, 0xd2, 0xfd, 0xbb, 0x02, 0xed, 0x9d, 0x66, 0x95, 0xa3, 0x15, 0x67,
	0xe1, 0x34, 0xd1, 0x1f, 0x3d, 0xf4, 0xcd, 0x0d, 0x79, 0x70, 0x14, 0x25, 0x44, 0xa6, 0x96, 0x16,
	0xb7, 0xbf, 0xb2, 0x67, 0xc2, 0x21, 0x6d, 0x74, 0x55, 0x94, 0x1e, 0xf7, 0x3d, 0xa0, 0x1c, 0x63,
	0x76, 0xcb, 0x91, 0xb5, 0xdf, 0x51, 0x4f, 0x9a, 0x94, 0xdd, 0xb8, 0xff, 0x58, 0xd0, 0x1a, 0x5f,
	0xec, 0xef, 0xd6, 0x8f, 0xe1, 0xf0, 0xfc, 0xdd, 0x38, 0xc8, 0x8b, 0xe9, 0x7a, 0xe9, 0x1c, 0x9c,
	0xbf, 0x1b, 0x4f, 0x8a, 0x69, 0x79, 0x87, 0x58, 0x1f, 0xda, 0x21, 0xb5, 0x7b, 0x77, 0xc8, 0x19,
	0x3c, 0x19, 0x5f, 0xdc, 0xaa, 0x40, 0x55, 0xf4, 0xad, 0xd1, 0xb1, 0x7a, 0x85, 0x0a, 0x6e, 0xa7,
	0xfd, 0xef, 0x1a, 0x3c, 0x64, 0x89, 0x34, 0xfe, 0xdf, 0x25, 0x72, 0xf0, 0xdf, 0x96, 0xc8, 0xe1,
	0xa3, 0x96, 0x48, 0xf3, 0x71, 0x4b, 0x04, 0x1e, 0xba, 0x44, 0xc6, 0x92, 0xfc, 0xdb, 0x03, 0x50,
	0x0f, 0xbb, 0xca, 0x66, 0xd8, 0x75, 0xa0, 0x3a, 0x79, 0xbd, 0x1e, 0x7e, 0x93, 0xd7, 0xf2, 0xee,
	0x9d, 0x99, 0x01, 0x57, 0xf5, 0xce, 0xde, 0x04, 0x70, 0x42, 0xd9, 0x7c, 0xb8, 0x58, 0xe5, 0x98,
	0x25, 0x38, 0x9e, 0x63, 0x36, 0xd4, 0x5f, 0xd1, 0xbf, 0x46, 0x5c, 0x86, 0xf1, 0xa6, 0x77, 0xc9,
	0x73, 0xed, 0x7f, 0x12, 0x46, 0xcb, 0x70, 0x8e, 0x7f, 0x1e, 0xcc, 0x89, 0x58, 0x14, 0xd3, 0x61,
	0x44, 0xd3, 0xd3, 0x92, 0xed, 0xa9, 0xb6, 0x3d, 0xd5, 0xb6, 0xa7, 0x29, 0xcf, 0xa7, 0x0d, 0x75,
	0x7e, 0xf5, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x53, 0xf9, 0x83, 0xea, 0x76, 0x09, 0x00, 0x00,
}
