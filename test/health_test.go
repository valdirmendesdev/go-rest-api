// +build e2e

package test

import (
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	client := resty.New()

	resp, err := client.R().Get("http://localhost:8080/api/health")
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode())
}