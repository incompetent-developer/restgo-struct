package storage

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"cloud.google.com/go/storage"
	"github.com/incompetent-developer/restgo-struct/cloudcredential"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

// Connect to cloud storage
func (gcs *Storage) Connect() error {
	/*
		Read credential file
	*/
	credentialJSON, err := ioutil.ReadFile(gcs.CredentialPath)
	if err != nil {
		return err
	}

	/*
		Parse credential
	*/
	var (
		jsonCredential cloudcredential.JSONCredential
	)
	if err := json.Unmarshal(credentialJSON, &jsonCredential); err != nil {
		return err
	}

	gcs.credential = jsonCredential

	/*
		Create background context
	*/
	ctx := context.Background()

	/*
		Credential from JSON
	*/
	creds, err := google.CredentialsFromJSON(
		ctx,
		credentialJSON,
		storage.ScopeFullControl,
	)

	client, err := storage.NewClient(
		ctx,
		option.WithCredentials(creds),
	)
	if err != nil {
		return err
	}

	/*
		Cache
	*/
	gcs.Client = client
	gcs.Ctx = ctx

	return nil
}
