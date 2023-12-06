package cloudpubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
)

// Pubsub is structure for cloud pub/sub
type Pubsub struct {
	Client         *pubsub.Client
	Ctx            context.Context
	Subscription   *pubsub.Subscription
	CredentialPath string `json:"credential_path"`
	CredentialData []byte
}
