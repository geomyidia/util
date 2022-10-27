package searchpath_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/geomyidia/util/pkg/searchpath"
)

func TestFindFileDefaultPaths(t *testing.T) {
	path, err := searchpath.FindFile("reader.go")
	assert.NoError(t, err)
	require.True(t, path != "")
}

func TestFindFileNotFound(t *testing.T) {
	path, err := searchpath.FindFile("no-such-file")
	assert.Error(t, err)
	require.True(t, path == "")
	assert.Equal(t, searchpath.ErrNotFound, err)
}

func TestFindFileStringPaths(t *testing.T) {
	opts := searchpath.WithPathStr("../version:../filesystem")
	path, err := searchpath.FindFile("caller.go", opts)
	assert.NoError(t, err)
	require.True(t, path != "")
	assert.True(t, len(path) > 0)
	path, err = searchpath.FindFile("reader.go", opts)
	assert.NoError(t, err)
	require.True(t, path != "")
	assert.True(t, len(path) > 0)
	path, err = searchpath.FindFile("util.go", opts)
	assert.NoError(t, err)
	require.True(t, path != "")
	assert.True(t, len(path) > 0)
}

func TestFindFileSlicePaths(t *testing.T) {
	opts := searchpath.WithPathSlice([]string{"../version", "../filesystem"})
	path, err := searchpath.FindFile("caller.go", opts)
	assert.NoError(t, err)
	require.True(t, path != "")
	assert.True(t, len(path) > 0)
	path, err = searchpath.FindFile("reader.go", opts)
	assert.NoError(t, err)
	require.True(t, path != "")
	assert.True(t, len(path) > 0)
	path, err = searchpath.FindFile("util.go", opts)
	assert.NoError(t, err)
	require.True(t, path != "")
	assert.True(t, len(path) > 0)
}

func TestFindFilePaths(t *testing.T) {
	opts := searchpath.WithPaths("../version", "../filesystem")
	path, err := searchpath.FindFile("caller.go", opts)
	assert.NoError(t, err)
	require.True(t, path != "")
	assert.True(t, len(path) > 0)
	path, err = searchpath.FindFile("reader.go", opts)
	assert.NoError(t, err)
	require.True(t, path != "")
	assert.True(t, len(path) > 0)
	path, err = searchpath.FindFile("util.go", opts)
	assert.NoError(t, err)
	require.True(t, path != "")
	assert.True(t, len(path) > 0)
}

func TestFindFilePathSep(t *testing.T) {
	opts := []searchpath.Option{
		searchpath.WithSeparator("!"),
		searchpath.WithPathStr("../version!../filesystem"),
	}
	bytes, err := searchpath.FindFile("caller.go", opts...)
	assert.NoError(t, err)
	require.NotNil(t, bytes)
	assert.True(t, len(bytes) > 0)
	bytes, err = searchpath.FindFile("util.go", opts...)
	assert.NoError(t, err)
	require.NotNil(t, bytes)
	assert.True(t, len(bytes) > 0)
}

func TestFindFilePathSepMismatch(t *testing.T) {
	opts := []searchpath.Option{
		searchpath.WithSeparator("!"),
		searchpath.WithPathStr("../version:../filesystem"),
	}
	path, err := searchpath.FindFile("caller.go", opts...)
	assert.Error(t, err)
	assert.Equal(t, path, "")
	assert.Equal(t, searchpath.ErrNotFound, err)
	path, err = searchpath.FindFile("util.go", opts...)
	assert.Error(t, err)
	assert.Equal(t, path, "")
	assert.Equal(t, searchpath.ErrNotFound, err)
}
