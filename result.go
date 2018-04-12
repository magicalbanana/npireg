package npireg

// Result represents the data returned after performing an NPI look up
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
	EnumerationType string `json:"enumeration_type"`
	Basic           struct {
		Status                              string `json:"status"`
		OrganizationName                    string `json:"organization_name"`
		LastName                            string `json:"last_name"`
		FirstName                           string `json:"first_name"`
		MiddleName                          string `json:"middle_name"`
		NamePrefix                          string `json:"name_prefix"`
		NameSuffix                          string `json:"name_suffix"`
		Name                                string `json:"name"`
		SoloProprietor                      string `json:"sole_proprietor"`
		Gender                              string `json:"gender"`
		LastUpdated                         string `json:"last_updated"`
		EnumerationDate                     string `json:"enumeration_date"`
		Credential                          string `json:"credential"`
		DeactivationReasonCode              string `json:"deactivation_reason_code"`
		DeactivationDate                    string `json:"deactivation_date"`
		ReactivationDate                    string `json:"reactivation_date"`
		AuthorizedOfficialLastName          string `json:"authorized_official_last_name"`
		AuthorizedOfficialFirstName         string `json:"authorized_official_first_name"`
		AuthorizedOfficialMiddleName        string `json:"authorized_official_middle_name"`
		AuthorizedOfficialTitleOrPosition   string `json:"authorized_official_title_or_position"`
		AuthorizedOfficialTelephoneNumber   string `json:"authorized_official_telephone_number"`
		AuthorizedOfficialNamePrefix        string `json:"authorized_official_name_prefix"`
		AuthorizedOfficiaNameSuffix         string `json:"authorized_official_name_suffix"`
		AuthorizedOfficialCredential        string `json:"authorized_official_credential"`
		OrganizationSubpart                 string `json:"organizational_subpart"`
		ParentOrganizationLegalBusinessName string `json:"parent_organization_legal_business_name"`
		ParentOrganizationEIN               string `json:"parent_organization_ein"`
	} `json:"basic"`
	OtherNames []struct {
		OrgnizationName string `json:"organization_name"`
		Code            string `json:"code"`
		LastName        string `json:"last_name"`
		FirstName       string `json:"first_name"`
		MiddleName      string `json:"middle_name"`
		Prefix          string `json:"prefix"`
		Suffix          string `json:"suffix"`
		Credential      string `json:"credential"`
		Type            string `json:"type"`
	} `json:"other_names"`
	Taxonomies []struct {
		State   string `json:"state"`
		Code    string `json:"code"`
		Primary bool   `json:"primary"`
		License string `json:"license"`
		Desc    string `json:"desc"`
	} `json:"taxonomies"`
	Addresses []struct {
		Address1    string `json:"address_1"`
		Address2    string `json:"address_2"`
		State       string `json:"state"`
		PostalCode  string `json:"postal_code"`
		CountryCode string `json:"country_code"`
		CountryName string `json:"country_name"`
		AddressType string `json:"address_type"`
		// Refers to whether the address information entered pertains to the provider's
		// Mailing Address or the provider's Practice Location Address. When not
		// specified, the results will contain the providers where either the Mailing
		// Address or the Practice Location Addresses match the entered address
		// information. Valid values are:
		// - LOCATION
		// - MAILING
		AddressPurpose  string `json:"address_purpose"`
		City            string `json:"city"`
		TelephoneNumber string `json:"telephone_number"`
		FaxNumber       string `json:"fax_number"`
	} `json:"addresses"`
	CreatedEpoch int `json:"created_epoch"`
	Identifiers  []struct {
		Code       string `json:"code"`
		Issuer     string `json:"issuer"`
		State      string `json:"state"`
		Identifier string `json:"identifier"`
		Desc       string `json:"desc"`
	} `json:"identifiers"`
	LastUpdatedEpoch int `json:"last_updated_epoch"`
}
