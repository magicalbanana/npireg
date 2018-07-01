package npireg

import "strings"

// Result represents the data returned after performing an NPI look up
// https://npiregistry.cms.hhs.gov/registry/Json-Conversion-Field-Map
type Result struct {
	// The NPI Number is the unique 10-digit National Provider Identifier
	// assigned to the provider.
	Number int `json:"number"`
	// The Read API can be refined to retrieve only Individual Providers or
	// Organizational Providers. When it is not specified, both Type 1 and
	// Type 2 NPIs will be returned. When using the Enumeration Type, it
	// cannot be the only criteria entered. Additional criteria must also be
	// entered as well. Valid values are:
	// - NPI-1: Individual Providers (Type 1) NPIs
	// - NPI-2: Organizational Providers (Type 2) NPIs
	EnumerationType  string       `json:"enumeration_type"`
	Basic            Basic        `json:"basic"`
	OtherNames       []OtherName  `json:"other_names,omitempty"`
	Taxonomies       []Taxonomy   `json:"taxonomies,omitempty"`
	Addresses        []Address    `json:"addresses,omitempty"`
	Identifiers      []Identifier `json:"identifiers,omitempty"`
	CreatedEpoch     int64        `json:"created_epoch"`
	LastUpdatedEpoch int64        `json:"last_updated_epoch"`
}

// Basic ...
type Basic struct {
	ReplacementNPI                      string `json:"replacement_npi,omitempty"`
	EIN                                 string `json:"ein,omitempty"`
	Status                              string `json:"status,omitempty"`
	OrganizationName                    string `json:"organization_name,omitempty"`
	LastName                            string `json:"last_name,omitempty"`
	FirstName                           string `json:"first_name,omitempty"`
	MiddleName                          string `json:"middle_name,omitempty"`
	NamePrefix                          string `json:"name_prefix,omitempty"`
	NameSuffix                          string `json:"name_suffix,omitempty"`
	Name                                string `json:"name,omitempty"`
	SoleProprietor                      string `json:"sole_proprietor,omitempty"`
	Gender                              string `json:"gender,omitempty"`
	LastUpdated                         string `json:"last_updated,omitempty"`
	EnumerationDate                     string `json:"enumeration_date,omitempty"`
	Credential                          string `json:"credential,omitempty"`
	DeactivationReasonCode              string `json:"deactivation_reason_code,omitempty"`
	DeactivationDate                    string `json:"deactivation_date,omitempty"`
	ReactivationDate                    string `json:"reactivation_date,omitempty"`
	AuthorizedOfficialLastName          string `json:"authorized_official_last_name,omitempty"`
	AuthorizedOfficialFirstName         string `json:"authorized_official_first_name,omitempty"`
	AuthorizedOfficialMiddleName        string `json:"authorized_official_middle_name,omitempty"`
	AuthorizedOfficialTitleOrPosition   string `json:"authorized_official_title_or_position,omitempty"`
	AuthorizedOfficialTelephoneNumber   string `json:"authorized_official_telephone_number,omitempty"`
	AuthorizedOfficialNamePrefix        string `json:"authorized_official_name_prefix,omitempty"`
	AuthorizedOfficialNameSuffix        string `json:"authorized_official_name_suffix,omitempty"`
	AuthorizedOfficialCredential        string `json:"authorized_official_credential,omitempty"`
	OrganizationSubpart                 string `json:"organizational_subpart,omitempty"`
	ParentOrganizationLegalBusinessName string `json:"parent_organization_legal_business_name,omitempty"`
	ParentOrganizationEIN               string `json:"parent_organization_ein,omitempty"`
}

// OtherName ...
type OtherName struct {
	OrgnizationName string `json:"organization_name,omitempty"`
	Code            string `json:"code,omitempty"`
	LastName        string `json:"last_name,omitempty"`
	FirstName       string `json:"first_name,omitempty"`
	MiddleName      string `json:"middle_name,omitempty"`
	Prefix          string `json:"prefix,omitempty"`
	Suffix          string `json:"suffix,omitempty"`
	Credential      string `json:"credential,omitempty"`
	Type            string `json:"type,omitempty"`
}

// Taxonomy ...
type Taxonomy struct {
	State         string `json:"state,omitempty"`
	Code          string `json:"code,omitempty"`
	Primary       bool   `json:"primary,omitempty"`
	License       string `json:"license,omitempty"`
	Desc          string `json:"desc,omitempty"`
	TaxonomyGroup string `json:"taxonomy_group,omitempty"`
}

// Address ..
type Address struct {
	Address1    string `json:"address_1,omitempty"`
	Address2    string `json:"address_2,omitempty"`
	State       string `json:"state,omitempty"`
	PostalCode  string `json:"postal_code,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	CountryName string `json:"country_name,omitempty"`
	AddressType string `json:"address_type,omitempty"`
	// Refers to whether the address information entered pertains to the provider's
	// Mailing Address or the provider's Practice Location Address. When not
	// specified, the results will contain the providers where either the Mailing
	// Address or the Practice Location Addresses match the entered address
	// information. Valid values are:
	// - LOCATION
	// - MAILING
	AddressPurpose  string `json:"address_purpose,omitempty"`
	City            string `json:"city,omitempty"`
	TelephoneNumber string `json:"telephone_number,omitempty"`
	FaxNumber       string `json:"fax_number,omitempty"`
}

// FormattedAddress ...
func (a *Address) FormattedAddress() string {
	aa := make([]string, 0)
	aa = append(aa, a.Address1)
	if a.Address2 != "" {
		aa = append(aa, a.Address2)
	}
	aa = append(aa, ",")
	aa = append(aa, a.City)
	aa = append(aa, ",")
	aa = append(aa, a.State)
	aa = append(aa, a.PostalCode)
	aa = append(aa, a.CountryCode)
	return strings.Join(aa, " ")
}

// Identifier ...
type Identifier struct {
	Code       string `json:"code,omitempty"`
	Issuer     string `json:"issuer,omitempty"`
	State      string `json:"state,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Desc       string `json:"desc,omitempty"`
}
