package searchpath_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/geomyidia/util/pkg/searchpath"
)

func TestReadFileDefaultPaths(t *testing.T) {
	bytes, err := searchpath.ReadFile("reader.go")
	assert.NoError(t, err)
	require.NotNil(t, bytes)
	assert.True(t, len(bytes) > 0)
}

func TestReadFileNotFound(t *testing.T) {
	bytes, err := searchpath.ReadFile("no-such-file")
	assert.Error(t, err)
	assert.Nil(t, bytes)
	assert.Equal(t, searchpath.ErrNotFound, err)
}

func TestReadFileStringPaths(t *testing.T) {
	opts := searchpath.WithPathStr("../version:../filesystem")
	bytes, err := searchpath.ReadFile("caller.go", opts)
	assert.NoError(t, err)
	require.NotNil(t, bytes)
	assert.True(t, len(bytes) > 0)
	bytes, err = searchpath.ReadFile("reader.go", opts)
	assert.NoError(t, err)
	require.NotNil(t, bytes)
	assert.True(t, len(bytes) > 0)
	bytes, err = searchpath.ReadFile("util.go", opts)
	assert.NoError(t, err)
	require.NotNil(t, bytes)
	assert.True(t, len(bytes) > 0)
}

func TestReadFileSlicePaths(t *testing.T) {
	opts := searchpath.WithPathSlice([]string{"../version", "../filesystem"})
	bytes, err := searchpath.ReadFile("caller.go", opts)
	assert.NoError(t, err)
	require.NotNil(t, bytes)
	assert.True(t, len(bytes) > 0)
	bytes, err = searchpath.ReadFile("reader.go", opts)
	assert.NoError(t, err)
	require.NotNil(t, bytes)
	assert.True(t, len(bytes) > 0)
	bytes, err = searchpath.ReadFile("util.go", opts)
	assert.NoError(t, err)
	require.NotNil(t, bytes)
	assert.True(t, len(bytes) > 0)
}

func TestReadFilePaths(t *testing.T) {
	opts := searchpath.WithPaths("../version", "../filesystem")
	bytes, err := searchpath.ReadFile("caller.go", opts)
	assert.NoError(t, err)
	require.NotNil(t, bytes)
	assert.True(t, len(bytes) > 0)
	bytes, err = searchpath.ReadFile("reader.go", opts)
	assert.NoError(t, err)
	require.NotNil(t, bytes)
	assert.True(t, len(bytes) > 0)
	bytes, err = searchpath.ReadFile("util.go", opts)
	assert.NoError(t, err)
	require.NotNil(t, bytes)
	assert.True(t, len(bytes) > 0)
}

func TestReadFilePathSep(t *testing.T) {
	opts := []searchpath.Option{
		searchpath.WithSeparator("!"),
		searchpath.WithPathStr("../version!../filesystem"),
	}
	bytes, err := searchpath.ReadFile("caller.go", opts...)
	assert.NoError(t, err)
	require.NotNil(t, bytes)
	assert.True(t, len(bytes) > 0)
	bytes, err = searchpath.ReadFile("util.go", opts...)
	assert.NoError(t, err)
	require.NotNil(t, bytes)
	assert.True(t, len(bytes) > 0)
	opts = []searchpath.Option{
		searchpath.WithSeparator("!"),
		searchpath.WithPathStr("../version:../filesystem"),
	}
	bytes, err = searchpath.ReadFile("caller.go", opts...)
	assert.Error(t, err)
	assert.Nil(t, bytes)
	assert.Equal(t, searchpath.ErrNotFound, err)
	bytes, err = searchpath.ReadFile("util.go", opts...)
	assert.Error(t, err)
	assert.Nil(t, bytes)
	assert.Equal(t, searchpath.ErrNotFound, err)
}
