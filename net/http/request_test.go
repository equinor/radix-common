package http

import (
	"net/http"
	"testing"

	"github.com/equinor/radix-common/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_GetImpersonationFromHeader(t *testing.T) {
	sut := GetImpersonationFromHeader

	// Request without Impersonate-* headers
	expected := models.Impersonation{}
	actual, err := sut(&http.Request{})
	require.NoError(t, err)
	assert.Equal(t, expected, actual)

	// Request without both Impersonate-* headers should return error
	_, err = sut(&http.Request{Header: http.Header{"Impersonate-User": []string{"any-user"}}})
	assert.Error(t, err)
	_, err = sut(&http.Request{Header: http.Header{"Impersonate-Group": []string{"any-group"}}})
	assert.Error(t, err)

	// Request with both Impersonate-* headers should succeed
	expected = models.Impersonation{User: "any-user", Groups: []string{"group1", "group2", "group3"}}
	actual, err = sut(&http.Request{Header: http.Header{"Impersonate-User": []string{"any-user"}, "Impersonate-Group": []string{"group1 , group2,group3"}}})
	require.NoError(t, err)
	assert.Equal(t, expected, actual)
}
