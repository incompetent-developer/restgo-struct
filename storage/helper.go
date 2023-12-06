package storage

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"cloud.google.com/go/storage"
)

// ObjectURL is return storage URL
func (gcs Storage) ObjectURL(objAttrs *storage.ObjectAttrs) string {
	return fmt.Sprintf("%s/%s/%s", prefixURL, objAttrs.Bucket, objAttrs.Name)
}

// Upload to google cloud storage
func (gcs Storage) Upload(r io.Reader, contentType, bucket, name string, isPublic bool) (*storage.ObjectHandle, *storage.ObjectAttrs, error) {
	/*
		Bucket
	*/
	bh := gcs.Client.Bucket(bucket)

	/*
		Next check if the bucket exists
	*/
	if _, err := bh.Attrs(gcs.Ctx); err != nil {
		return nil, nil, err
	}

	obj := bh.Object(name)
	w := obj.NewWriter(gcs.Ctx)

	/*
		Content type
	*/
	w.ContentType = contentType

	/*
		Copy
	*/
	if _, err := io.Copy(w, r); err != nil {
		return nil, nil, err
	}
	if err := w.Close(); err != nil {
		return nil, nil, err
	}

	/*
		Public ?
	*/
	if isPublic {
		if err := obj.ACL().Set(gcs.Ctx, storage.AllUsers, storage.RoleReader); err != nil {
			return nil, nil, err
		}
	}

	attrs, err := obj.Attrs(gcs.Ctx)

	return obj, attrs, err
}

// RemoveObject Delete deletes the single specified object.
func (gcs Storage) RemoveObject(bucket, path string) error {
	/*
		Bucket
	*/
	bh := gcs.Client.Bucket(bucket)

	/*
		Delete object
	*/
	if err := bh.Object(path).Delete(gcs.Ctx); err != nil {
		switch err {
		case storage.ErrBucketNotExist:
			return ErrBucketNotExist
		case storage.ErrObjectNotExist:
			return ErrObjectNotExist
		default:
			return err
		}
	}

	return nil
}

// GenerateSignedURL is generate url with own credential
// Who know this url, Can access this file
func (gcs Storage) GenerateSignedURL(bucket, path string, expiresInSec int64) (string, error) {
	return storage.SignedURL(
		bucket,
		path,
		&storage.SignedURLOptions{
			GoogleAccessID: gcs.credential.ClientEmail,
			PrivateKey:     []byte(gcs.credential.PrivateKey),
			Method:         http.MethodGet,
			Expires:        time.Now().Add(time.Duration(expiresInSec) * time.Second),
		},
	)
}

// IsBucketNotExistError will check error equal bucket not exists error
func IsBucketNotExistError(err error) bool {
	return err == storage.ErrBucketNotExist
}

// IsObjectNotExistError will check error equal object not exists error
func IsObjectNotExistError(err error) bool {
	return err == storage.ErrObjectNotExist
}
