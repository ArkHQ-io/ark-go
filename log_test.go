// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ark_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/ArkHQ-io/ark-go"
	"github.com/ArkHQ-io/ark-go/internal/testutil"
	"github.com/ArkHQ-io/ark-go/option"
)

func TestLogGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := ark.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Logs.Get(context.TODO(), "req_V8GGcdWYzgeWIHiI")
	if err != nil {
		var apierr *ark.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLogListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := ark.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Logs.List(context.TODO(), ark.LogListParams{
		CredentialID: ark.String("credentialId"),
		EndDate:      ark.Time(time.Now()),
		Endpoint:     ark.String("endpoint"),
		Page:         ark.Int(1),
		PerPage:      ark.Int(1),
		RequestID:    ark.String("requestId"),
		StartDate:    ark.Time(time.Now()),
		Status:       ark.LogListParamsStatusSuccess,
		StatusCode:   ark.Int(100),
	})
	if err != nil {
		var apierr *ark.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
