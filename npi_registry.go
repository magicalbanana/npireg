package npireg

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// NPIRegistry ...
type NPIRegistry interface {
	Search(ctx context.Context, opt *SearchOpts) (*SearchResult, error)
}

// npiRegistry ...
type npiRegistry struct {
	client  *http.Client
	baseURL *url.URL
}

// NewNPIRegistry ...
func NewNPIRegistry(baseURL string, client *http.Client) (NPIRegistry, error) {
	bURL, err := url.ParseRequestURI(baseURL)
	if err != nil {
		return nil, err
	}
	return &npiRegistry{
		client: client, baseURL: bURL,
	}, nil
}

// SearchPath ...
const SearchPath = "api"

// Search makes a request to the NPI registry to look up the information
// related to the given npiNumber. The second returned data is the raw data
// that is returned from the request
func (n *npiRegistry) Search(ctx context.Context, opt *SearchOpts) (*SearchResult, error) {
	p, err := url.Parse(SearchPath)
	if err != nil {
		return nil, err
	}

	u := n.baseURL
	q := u.Query()
	u = u.ResolveReference(p)
	q = opt.BuildQueryParams(q)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	req.URL.RawQuery = opt.BuildQueryParams(req.URL.Query()).Encode()

	resp, err := n.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			// Do nothing
		}

		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		errx := Error{}
		errx.Message = fmt.Sprintf("NPI Registry return a status of: %v", resp.StatusCode)
		errx.Response.StatusCode = resp.StatusCode
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		errx.Response.Body = string(b)
		return nil, errx
	}

	var buf bytes.Buffer

	tee := io.TeeReader(resp.Body, &buf)
	b1, err := ioutil.ReadAll(tee)
	if err != nil {
		return nil, err
	}

	b2, err := ioutil.ReadAll(&buf)
	if err != nil {
		return nil, err
	}

	searchResult := &SearchResult{}
	unMarshalErr := json.Unmarshal(b2, &searchResult)
	if unMarshalErr != nil {
		return nil, unMarshalErr
	}
	searchResult.Raw = b1

	return searchResult, nil
}

// SearchResult represents the data that is returned from making a request to
// https://npiregistry.cms.hhs.gov/api/
type SearchResult struct {
	ResultCount int      `json:"result_count"`
	Results     []Result `json:"results"`
	Raw         []byte
}

// SearchOpts ...
type SearchOpts struct {
	NPINumber           string
	EnumerationType     EnumerationType
	TaxonomyDescription string
	FirstName           string
	LastName            string
	OrganizationName    string
	AddressPurpose      AddressPurpose
	City                string
	State               string
	PostalCode          string
	CountryCode         string
	Limit               int
	Skip                int
	Pretty              bool
}

// BuildQueryParams builds the query params from the given url.Values
// following the parameters from https://npiregistry.cms.hhs.gov/api/demo
func (s *SearchOpts) BuildQueryParams(qryValues url.Values) url.Values {
	qryValues = buildQueryParams("number", s.NPINumber, qryValues)
	if s.EnumerationType != nil {
		qryValues = buildQueryParams("enumeration_type", s.EnumerationType.String(), qryValues)
	}
	qryValues = buildQueryParams("taxonomy_description", s.TaxonomyDescription, qryValues)
	qryValues = buildQueryParams("first_name", s.FirstName, qryValues)
	qryValues = buildQueryParams("last_name", s.LastName, qryValues)
	qryValues = buildQueryParams("organization_name", s.OrganizationName, qryValues)
	if s.AddressPurpose != nil {
		qryValues = buildQueryParams("address_purpose", s.AddressPurpose.String(), qryValues)
	}
	qryValues = buildQueryParams("city", s.City, qryValues)
	qryValues = buildQueryParams("state", s.State, qryValues)
	qryValues = buildQueryParams("postal_code", s.PostalCode, qryValues)
	qryValues = buildQueryParams("country_code", s.CountryCode, qryValues)
	if s.Limit != 0 {
		qryValues = buildQueryParams("limit", strconv.Itoa(s.Limit), qryValues)
	}
	if s.Skip != 0 {
		qryValues = buildQueryParams("skip", strconv.Itoa(s.Skip), qryValues)
	}
	if s.Pretty {
		qryValues = buildQueryParams("pretty", "on", qryValues)
	}
	return qryValues
}

func buildQueryParams(key, value string, u url.Values) url.Values {
	if value != "" {
		u.Add(key, value)
	}
	return u
}
