// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// CertificateAuthority undocumented
type CertificateAuthority struct {
	// Object is the base model of CertificateAuthority
	Object
	// IsRootAuthority undocumented
	IsRootAuthority *bool `json:"isRootAuthority,omitempty"`
	// CertificateRevocationListURL undocumented
	CertificateRevocationListURL *string `json:"certificateRevocationListUrl,omitempty"`
	// DeltaCertificateRevocationListURL undocumented
	DeltaCertificateRevocationListURL *string `json:"deltaCertificateRevocationListUrl,omitempty"`
	// Certificate undocumented
	Certificate *Binary `json:"certificate,omitempty"`
	// Issuer undocumented
	Issuer *string `json:"issuer,omitempty"`
	// IssuerSki undocumented
	IssuerSki *string `json:"issuerSki,omitempty"`
}

// CertificateBasedAuthConfiguration undocumented
type CertificateBasedAuthConfiguration struct {
	// Entity is the base model of CertificateBasedAuthConfiguration
	Entity
	// CertificateAuthorities undocumented
	CertificateAuthorities []CertificateAuthority `json:"certificateAuthorities,omitempty"`
}

// CertificateConnectorSetting undocumented
type CertificateConnectorSetting struct {
	// Object is the base model of CertificateConnectorSetting
	Object
	// Status Certificate connector status
	Status *int `json:"status,omitempty"`
	// CertExpiryTime Certificate expire time
	CertExpiryTime *time.Time `json:"certExpiryTime,omitempty"`
	// EnrollmentError Certificate connector enrollment error
	EnrollmentError *string `json:"enrollmentError,omitempty"`
	// LastConnectorConnectionTime Last time certificate connector connected
	LastConnectorConnectionTime *time.Time `json:"lastConnectorConnectionTime,omitempty"`
	// ConnectorVersion Version of certificate connector
	ConnectorVersion *string `json:"connectorVersion,omitempty"`
	// LastUploadVersion Version of last uploaded certificate connector
	LastUploadVersion *int `json:"lastUploadVersion,omitempty"`
}
