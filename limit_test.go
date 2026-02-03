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

func TestLimitGet(t *testing.T) {
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
	_, err := client.Limits.Get(context.TODO())
	if err != nil {
		var apierr *ark.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
