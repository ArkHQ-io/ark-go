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

func TestManualPagination(t *testing.T) {
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
	page, err := client.Emails.List(context.TODO(), ark.EmailListParams{
		Page:    ark.Int(1),
		PerPage: ark.Int(10),
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	for _, email := range page.Data {
		t.Logf("%+v\n", email.ID)
	}
	// The mock server isn't going to give us real pagination
	page, err = page.GetNextPage()
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if page != nil {
		for _, email := range page.Data {
			t.Logf("%+v\n", email.ID)
		}
	}
}
