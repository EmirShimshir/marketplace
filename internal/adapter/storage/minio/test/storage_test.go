package storage

import (
	"bytes"
	"context"
	storage "github.com/EmirShimshir/marketplace/internal/adapter/storage/minio"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestObjectStorage(t *testing.T) {
	ctx := context.Background()
	container, err := newMinioContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// Clean up the container after the test is complete
	t.Cleanup(func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	})

	url, err := container.ConnectionString(ctx)
	if err != nil {
		t.Fatal(err)
	}
	minioConfig.Host = url

	minioClient, err := newMinioClient(url)
	if err != nil {
		t.Fatal(err)
	}
	store := storage.NewObjectStorage(minioClient, &minioConfig)

	_, path, _, ok := runtime.Caller(0)
	require.Equal(t, ok, true)
	filesPath := filepath.Dir(path) + "/files"

	t.Run("test save png file", func(t *testing.T) {
		testFilename := "test.png"
		data, err := os.ReadFile(filepath.Join(filesPath, testFilename))
		if err != nil {
			t.Errorf("failed to read file %s: %s", testFilename, err)
		}

		fileUrl, err := store.SaveFile(ctx, domain.File{
			Name:   testFilename,
			Path:   "user",
			Reader: bytes.NewReader(data),
		})
		if err != nil {
			t.Errorf("failed to save file to minio: %v", err)
		}

		resp, err := http.Get(fileUrl.String())
		if err != nil {
			t.Errorf("failed to download saved file: %v", err)
		}
		defer resp.Body.Close()

		savedData, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("failed to read response data: %s", err)
		}
		require.Equal(t, reflect.DeepEqual(data, savedData), true)
	})
}
