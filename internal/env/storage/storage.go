package storage

import (
	"context"

	"github.com/qilin/crm-api/internal/config"
	"gocloud.dev/blob"
	_ "gocloud.dev/blob/fileblob"
	_ "gocloud.dev/blob/memblob"
	_ "gocloud.dev/blob/s3blob"
)

type Env struct {
	Bucket *blob.Bucket
}

func New(ctx context.Context, cfg config.StorageConf) (*Env, error) {
	bucket, err := blob.OpenBucket(ctx, cfg.Bucket)
	if err != nil {
		return nil, err
	}

	return &Env{
		Bucket: bucket,
	}, nil
}
