// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ark_test

import (
	"context"
	"os"
	"testing"

	"github.com/ArkHQ-io/ark-go"
	"github.com/ArkHQ-io/ark-go/internal/testutil"
	"github.com/ArkHQ-io/ark-go/option"
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
	response, err := client.Emails.Send(context.TODO(), ark.EmailSendParams{
		From:    "hello@yourdomain.com",
		Subject: "Hello World",
		To:      []string{"user@example.com"},
		HTML:    ark.String("<h1>Welcome!</h1>"),
		Metadata: map[string]string{
			"user_id":  "usr_123456",
			"campaign": "onboarding",
		},
		Tag: ark.String("welcome"),
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", response.Data)
}
