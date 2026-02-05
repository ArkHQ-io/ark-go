// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ark

import (
	"github.com/ArkHQ-io/ark-go/option"
)

// PlatformService contains methods and other services that help with interacting
// with the ark API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPlatformService] method instead.
type PlatformService struct {
	Options  []option.RequestOption
	Webhooks PlatformWebhookService
}

// NewPlatformService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPlatformService(opts ...option.RequestOption) (r PlatformService) {
	r = PlatformService{}
	r.Options = opts
	r.Webhooks = NewPlatformWebhookService(opts...)
	return
}
