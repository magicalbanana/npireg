package npireg

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSearch_Integration(t *testing.T) {
	c := &http.Client{
		Timeout: time.Second * 5,
	}

	npiReg, err := NewNPIRegistry("https://npiregistry.cms.hhs.gov", c)
	assert.NoError(t, err)

	opts := &SearchOpts{
		OrganizationName: "rush*",
		City:             "Chicago",
	}
	result, err := npiReg.Search(context.Background(), opts)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result.ResultCount)
}
