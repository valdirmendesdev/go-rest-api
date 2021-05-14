// +build e2e

package test

import (
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetComments(t *testing.T) {
	client := resty.New()
	resp, err := client.R().Get(BASE_URL + "/api/comments")
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode())
}

func TestPostComment(t *testing.T) {
	client := resty.New()
	resp, err := client.R().
		SetBody(`{"Slug":"/", "Author":"123456", "Body":"hello world"}`).
		Post(BASE_URL + "/api/comments")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode())
}