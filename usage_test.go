// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ark_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/ark-go"
	"github.com/stainless-sdks/ark-go/internal/testutil"
	"github.com/stainless-sdks/ark-go/option"
)

func TestUsage(t *testing.T) {
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
	t.Skip("Prism tests are disabled")
	sendEmail, err := client.Emails.Send(context.TODO(), ark.EmailSendParams{
		From:    "Security <security@myapp.com>",
		Subject: "Reset your password",
		To:      []string{"user@example.com"},
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", sendEmail.Data)
}
