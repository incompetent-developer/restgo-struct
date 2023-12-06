package storage

import (
	"context"

	"cloud.google.com/go/storage"
	"github.com/incompetent-developer/restgo-struct/cloudcredential"
)

type ImageStorage struct {
	Url            string `json:"url"`
	Storage        string `json:"storage"`
	CredentialPath string `json:"credential_path"`
}

// Storage is cloud firestore details
type Storage struct {
	Client            *storage.Client
	Ctx               context.Context
	BucketPublic      string `json:"bucket_public"`
	BucketPublicPath  string `json:"bucket_public_path"`
	BucketPrivate     string `json:"bucket_private"`
	BucketPrivatePath string `json:"bucket_private_path"`
	CredentialPath    string `json:"credential_path"`
	StorageURL        string `json:"storage_url"`
	credential        cloudcredential.JSONCredential
}
