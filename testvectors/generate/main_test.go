package main

import (
	"errors"
	"os"
	"testing"

	"github.com/cinode/go-common/cutl"
	"github.com/cinode/go-common/picotestify/require"
)

func TestMainFunction(t *testing.T) {
	const dataDir = "../data"
	const dataDirBak = "../data.bak"

	require.NoError(t, os.Rename(dataDir, dataDirBak))
	require.NoError(t, os.MkdirAll(dataDir, 0o755))

	defer func() {
		require.NoError(t, os.RemoveAll(dataDir))
		require.NoError(t, os.Rename(dataDirBak, dataDir))
	}()

	main()

	fi, err := os.Stat(dataDir)
	require.NoError(t, err)

	require.True(t, fi.IsDir())
}

func TestPanicIfError(t *testing.T) {
	t.Run("no error", func(t *testing.T) {
		cutl.PanicIfError(nil)
	})

	t.Run("error", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Fatal("panic expected")
			} else if rErr, isError := r.(error); !isError {
				t.Fatalf("expected panic with error, got %v", r)
			} else if rErr.Error() != "test" {
				t.Fatalf("expected panic with error 'test', got %v", rErr.Error())
			}
		}()

		cutl.PanicIfError(errors.New("test"))
	})
}
