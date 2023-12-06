package cloudpubsub

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"cloud.google.com/go/pubsub"
	"github.com/incompetent-developer/restgo-struct/cloudcredential"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

// Connect is connect to firestore server
func (ps *Pubsub) Connect() error {
	/*
		Check credential path
	*/
	if ps.CredentialPath != "" {
		/*
			Read credential file
		*/
		credentialJSON, err := ioutil.ReadFile(ps.CredentialPath)
		if err != nil {
			return err
		}

		ps.CredentialData = credentialJSON

	}

	/*
		Parse credential JSON
	*/
	var (
		jsonCredential cloudcredential.JSONCredential
	)

	if err := json.Unmarshal(ps.CredentialData, &jsonCredential); err != nil {
		return err
	}

	/*
		Create background context
	*/
	ctx := context.Background()

	/*
		Credential from JSON
	*/
	creds, err := google.CredentialsFromJSON(
		ctx,
		ps.CredentialData,
		pubsub.ScopePubSub,
	)

	client, err := pubsub.NewClient(
		ctx,
		jsonCredential.ProjectID,
		option.WithCredentials(creds),
	)
	if err != nil {
		return err
	}

	ps.Client = client
	ps.Ctx = ctx

	return nil
}
