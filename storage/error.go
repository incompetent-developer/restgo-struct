package storage

import (
	"cloud.google.com/go/storage"
)

var (
	// ErrBucketNotExist indicates that the bucket does not exist.
	ErrBucketNotExist = storage.ErrBucketNotExist
	// ErrObjectNotExist indicates that the object does not exist.
	ErrObjectNotExist = storage.ErrObjectNotExist
)
