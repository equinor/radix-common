package http

import (
	"github.com/stretchr/testify/assert"
	netHeader "net/http"
	"testing"
)

type headerAcceptScenario struct {
	expectedAcceptSpecs []AcceptSpec
	acceptString        []string
	expectedError       string
}

func Test_GetAccess(t *testing.T) {
	scenarios := []headerAcceptScenario{
		{
			acceptString: []string{"text/html"},
			expectedAcceptSpecs: []AcceptSpec{
				{Value: "text/html", Q: 1.0},
			},
		},
		{
			acceptString: []string{"text/html;q=1.0"},
			expectedAcceptSpecs: []AcceptSpec{
				{Value: "text/html", Q: 1.0},
			},
		},
		{
			acceptString: []string{"text/html;q=0.9, application/json"},
			expectedAcceptSpecs: []AcceptSpec{
				{Value: "text/html", Q: 0.9},
				{Value: "application/json", Q: 1.0},
			},
		},
		{
			acceptString: []string{"text/plain,     text/html ;q=0.9, application/json ; q=0.3"},
			expectedAcceptSpecs: []AcceptSpec{
				{Value: "text/plain", Q: 1.0},
				{Value: "text/html", Q: 0.9},
				{Value: "application/json", Q: 0.3},
			},
		},
		{
			acceptString: []string{"text/plain,  image/*; q=0.9, image/png ; q=0.3, */*; q=0.1"},
			expectedAcceptSpecs: []AcceptSpec{
				{Value: "text/plain", Q: 1.0},
				{Value: "image/*", Q: 0.9},
				{Value: "image/png", Q: 0.3},
				{Value: "*/*", Q: 0.1},
			},
		},
		{
			acceptString:        []string{"text/plain q=t.f,     text/html ;q=0.9, application/json ; q=0.3"},
			expectedAcceptSpecs: []AcceptSpec{},
			expectedError:       "invalid element in header Accept: 'text/plain q=t.f'",
		},
		{
			acceptString:        []string{"q=1.0"},
			expectedAcceptSpecs: []AcceptSpec{},
		},
		{
			acceptString:        []string{"q=1.0, q=0.5"},
			expectedAcceptSpecs: []AcceptSpec{},
		},
	}
	t.Run("Get Accept", func(t *testing.T) {
		t.Parallel()
		for _, scenario := range scenarios {
			header := netHeader.Header{
				"Accept": scenario.acceptString,
			}
			acceptList, err := GetAccepts(header)
			assert.Equal(t, len(scenario.expectedAcceptSpecs), len(acceptList))
			errorMessage := getErrorMessage(err)
			assert.Equal(t, scenario.expectedError, errorMessage)
			for i, acceptSpec := range acceptList {
				assert.Equal(t, scenario.expectedAcceptSpecs[i].Value, acceptSpec.Value)
				assert.Equal(t, scenario.expectedAcceptSpecs[i].Q, acceptSpec.Q)
			}
		}
	})
}

func getErrorMessage(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
