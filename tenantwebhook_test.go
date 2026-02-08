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

func TestTenantWebhookNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Tenants.Webhooks.New(
		context.TODO(),
		"cm6abc123def456",
		ark.TenantWebhookNewParams{
			Name:      "My App Webhook",
			URL:       "https://myapp.com/webhooks/email",
			AllEvents: ark.Bool(true),
			Enabled:   ark.Bool(true),
			Events:    []string{"MessageSent", "MessageDeliveryFailed", "MessageBounced"},
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

func TestTenantWebhookGet(t *testing.T) {
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
	_, err := client.Tenants.Webhooks.Get(
		context.TODO(),
		"123",
		ark.TenantWebhookGetParams{
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

func TestTenantWebhookUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Tenants.Webhooks.Update(
		context.TODO(),
		"123",
		ark.TenantWebhookUpdateParams{
			TenantID:  "cm6abc123def456",
			AllEvents: ark.Bool(true),
			Enabled:   ark.Bool(true),
			Events:    []string{"string"},
			Name:      ark.String("name"),
			URL:       ark.String("https://example.com"),
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

func TestTenantWebhookList(t *testing.T) {
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
	_, err := client.Tenants.Webhooks.List(context.TODO(), "cm6abc123def456")
	if err != nil {
		var apierr *ark.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTenantWebhookDelete(t *testing.T) {
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
	_, err := client.Tenants.Webhooks.Delete(
		context.TODO(),
		"123",
		ark.TenantWebhookDeleteParams{
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

func TestTenantWebhookListDeliveriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Tenants.Webhooks.ListDeliveries(
		context.TODO(),
		"123",
		ark.TenantWebhookListDeliveriesParams{
			TenantID: "cm6abc123def456",
			After:    ark.Int(0),
			Before:   ark.Int(0),
			Event:    ark.TenantWebhookListDeliveriesParamsEventMessageSent,
			Page:     ark.Int(1),
			PerPage:  ark.Int(1),
			Success:  ark.Bool(true),
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

func TestTenantWebhookReplayDelivery(t *testing.T) {
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
	_, err := client.Tenants.Webhooks.ReplayDelivery(
		context.TODO(),
		"whr_abc123def456",
		ark.TenantWebhookReplayDeliveryParams{
			TenantID:  "cm6abc123def456",
			WebhookID: "123",
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

func TestTenantWebhookGetDelivery(t *testing.T) {
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
	_, err := client.Tenants.Webhooks.GetDelivery(
		context.TODO(),
		"whr_abc123def456",
		ark.TenantWebhookGetDeliveryParams{
			TenantID:  "cm6abc123def456",
			WebhookID: "123",
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

func TestTenantWebhookTest(t *testing.T) {
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
	_, err := client.Tenants.Webhooks.Test(
		context.TODO(),
		"123",
		ark.TenantWebhookTestParams{
			TenantID: "cm6abc123def456",
			Event:    ark.TenantWebhookTestParamsEventMessageSent,
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
