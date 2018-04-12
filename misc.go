package npireg

// NPIType ...
type NPIType interface {
	String() string
}

type npiType string

// String ...
func (n npiType) String() string {
	return string(n)
}

// Individual ...
const Individual npiType = "INDIVIDUAL"

// Organization ...
const Organization npiType = "ORGANIZATION"

// EnumerationType ...
type EnumerationType interface {
	String() string
}

type enumerationType string

// String ...
func (e enumerationType) String() string {
	return string(e)
}

// NPI1 ...
const NPI1 enumerationType = "NPI-1"

// NPI2 ...
const NPI2 enumerationType = "NPI-2"

// AddressPurpose ...
type AddressPurpose interface {
	String() string
}

type addressPurpose string

// String ...
func (a addressPurpose) String() string {
	return string(a)
}

// Location ...
const Location addressPurpose = "LOCATION"

// Mailing ...
const Mailing addressPurpose = "MAILING"
