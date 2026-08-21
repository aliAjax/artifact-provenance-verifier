package adapter

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
)

type BlobStore struct{ Root string }

func NewBlobStore(root string) *BlobStore { return &BlobStore{Root: root} }
func (b *BlobStore) Put(_ context.Context, key string, data []byte) error {
	p := filepath.Join(b.Root, key)
	if e := os.MkdirAll(filepath.Dir(p), 0750); e != nil {
		return e
	}
	return os.WriteFile(p, data, 0600)
}
func (b *BlobStore) Get(_ context.Context, key string) ([]byte, error) {
	p := filepath.Join(b.Root, key)
	d, e := os.ReadFile(p)
	if e != nil {
		return nil, fmt.Errorf("blob read: %w", e)
	}
	return d, nil
}
func (b *BlobStore) Delete(_ context.Context, key string) error {
	return os.Remove(filepath.Join(b.Root, key))
}
