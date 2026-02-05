// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ark_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/ArkHQ-io/ark-go"
	"github.com/ArkHQ-io/ark-go/internal/testutil"
	"github.com/ArkHQ-io/ark-go/option"
)

func TestTenantDomainNew(t *testing.T) {
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
	_, err := client.Tenants.Domains.New(
		context.TODO(),
		"cm6abc123def456",
		ark.TenantDomainNewParams{
			Name: "notifications.myapp.com",
		},
	)
	if err != nil {
		var apierr *ark.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTenantDomainGet(t *testing.T) {
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
	_, err := client.Tenants.Domains.Get(
		context.TODO(),
		"123",
		ark.TenantDomainGetParams{
			TenantID: "cm6abc123def456",
		},
	)
	if err != nil {
		var apierr *ark.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTenantDomainList(t *testing.T) {
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
	_, err := client.Tenants.Domains.List(context.TODO(), "cm6abc123def456")
	if err != nil {
		var apierr *ark.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTenantDomainDelete(t *testing.T) {
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
	_, err := client.Tenants.Domains.Delete(
		context.TODO(),
		"123",
		ark.TenantDomainDeleteParams{
			TenantID: "cm6abc123def456",
		},
	)
	if err != nil {
		var apierr *ark.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTenantDomainVerify(t *testing.T) {
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
	_, err := client.Tenants.Domains.Verify(
		context.TODO(),
		"123",
		ark.TenantDomainVerifyParams{
			TenantID: "cm6abc123def456",
		},
	)
	if err != nil {
		var apierr *ark.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
