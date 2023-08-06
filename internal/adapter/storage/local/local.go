package localstorageadapter

import (
	"context"
	"io"
	"os"
	"path"
	"path/filepath"

	storageadapter "github.com/wahyudibo/go-todo-api/internal/adapter/storage"
	"github.com/wahyudibo/go-todo-api/internal/config"
)

var _ storageadapter.StorageAdapter = (*localStorageAdapter)(nil)

type localStorageAdapter struct {
	config *config.Config
}

func New(cfg *config.Config) (storageadapter.StorageAdapter, error) {
	return &localStorageAdapter{config: cfg}, nil
}

func (sa *localStorageAdapter) Upload(_ context.Context, directory, objectKey string, body io.Reader) (string, error) {
	dirPath := filepath.Join(sa.config.StoragePath, directory)

	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		return "", err
	}

	fullPath := filepath.Join(dirPath, objectKey)

	f, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = io.Copy(f, body)
	if err != nil {
		return "", err
	}

	return path.Join(directory, objectKey), nil
}

func (sa *localStorageAdapter) GenerateDownloadURL(ctx context.Context, objectKey string) (string, error) {
	return path.Join(sa.config.LocalStorageDownloadPrefixUrl, objectKey), nil
}
