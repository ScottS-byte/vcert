/*
 * Copyright 2023 Venafi, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package domain

import (
	"github.com/Venafi/vcert/v5/pkg/certificate"
	"github.com/Venafi/vcert/v5/pkg/util"
)

// PlaybookRequest Contains data needed to generate a certificate request
// CSR is a PEM-encoded Certificate Signing PlaybookRequest
type PlaybookRequest struct {
	CADN           string                       `yaml:"cadn,omitempty"`
	ChainOption    certificate.ChainOption      `yaml:"chain,omitempty"`
	CsrOrigin      string                       `yaml:"csr,omitempty"`
	CustomFields   []certificate.CustomField    `yaml:"fields,omitempty"`
	DNSNames       []string                     `yaml:"sanDNS,omitempty"`
	EmailAddresses []string                     `yaml:"sanEmail,omitempty"`
	FriendlyName   string                       `yaml:"nickname,omitempty"`
	IPAddresses    []string                     `yaml:"sanIP,omitempty"`
	IssuerHint     util.IssuerHint              `yaml:"issuerHint,omitempty"`
	KeyCurve       certificate.EllipticCurve    `yaml:"keyCurve,omitempty"`
	KeyLength      int                          `yaml:"keySize,omitempty"`
	KeyPassword    string                       `yaml:"-"`
	KeyType        certificate.KeyType          `yaml:"keyType,omitempty"`
	Location       certificate.Location         `yaml:"location,omitempty"`
	OmitSANs       bool                         `yaml:"omitSans,omitempty"`
	Origin         string                       `yaml:"appInfo,omitempty"`
	Subject        Subject                      `yaml:"subject,omitempty"`
	Timeout        int                          `yaml:"timeout,omitempty"`
	UPNs           []string                     `yaml:"sanUPN,omitempty"`
	URIs           []string                     `yaml:"sanURI,omitempty"`
	ExtKeyUsages   certificate.ExtKeyUsageSlice `yaml:"eku,omitempty"`
	ValidDays      string                       `yaml:"validDays,omitempty"`
	Zone           string                       `yaml:"zone,omitempty"`
}
