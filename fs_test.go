package fs_test

import (
	"embed"
	"testing"

	"github.com/go-universal/fs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed *.go
var embeds embed.FS

func TestDir(t *testing.T) {
	fs := fs.NewDir(".")
	invalid := "missing.go"
	valid := "fs.go"

	t.Run("Exists", func(t *testing.T) {
		ok, err := fs.Exists(invalid)
		require.NoError(t, err)
		assert.False(t, ok, "%s should not exist!", invalid)

		ok, err = fs.Exists(valid)
		require.NoError(t, err)
		assert.True(t, ok, "%s should exist!", valid)
	})

	t.Run("Open", func(t *testing.T) {
		_, err := fs.Open(invalid)
		assert.Error(t, err, "%s should not be opened!", invalid)

		file, err := fs.Open(valid)
		require.NoError(t, err)
		require.NotNil(t, file)
		file.Close()
	})

	t.Run("ReadFile", func(t *testing.T) {
		_, err := fs.ReadFile(invalid)
		assert.Error(t, err, "%s should not be read!", invalid)

		content, err := fs.ReadFile(valid)
		require.NoError(t, err)
		assert.NotEmpty(t, content, "%s should not be empty!", valid)
	})

	t.Run("Search", func(t *testing.T) {
		dir := "."
		phrase := `fs_flex`
		ignore := ""
		ext := "go"

		result, err := fs.Search(dir, phrase, ignore, ext)
		require.NoError(t, err)
		assert.NotNil(t, result, "Search should find a result!")
	})

	t.Run("Find", func(t *testing.T) {
		dir := "."
		pattern := ".*_test.*"

		result, err := fs.Find(dir, pattern)
		require.NoError(t, err)
		assert.NotNil(t, result, "Find should find a result!")
	})

	t.Run("Lookup", func(t *testing.T) {
		dir := "."
		pattern := `.*\.go`

		results, err := fs.Lookup(dir, pattern)
		require.NoError(t, err)
		assert.NotEmpty(t, results, "Lookup should find results!")
	})

	t.Run("FS", func(t *testing.T) {
		assert.NotNil(t, fs.FS(), "FS should return a valid fs.FS instance!")
	})

	t.Run("Http", func(t *testing.T) {
		assert.NotNil(t, fs.Http(), "Http should return a valid http.FileSystem instance!")
	})
}

func TestEmbed(t *testing.T) {
	fs := fs.NewEmbed(embeds)
	invalid := "missing.go"
	valid := "fs.go"

	t.Run("Exists", func(t *testing.T) {
		ok, err := fs.Exists(invalid)
		require.NoError(t, err)
		assert.False(t, ok, "%s should not exist!", invalid)

		ok, err = fs.Exists(valid)
		require.NoError(t, err)
		assert.True(t, ok, "%s should exist!", valid)
	})

	t.Run("Open", func(t *testing.T) {
		_, err := fs.Open(invalid)
		assert.Error(t, err, "%s should not be opened!", invalid)

		file, err := fs.Open(valid)
		require.NoError(t, err)
		require.NotNil(t, file)
		file.Close()
	})

	t.Run("ReadFile", func(t *testing.T) {
		_, err := fs.ReadFile(invalid)
		assert.Error(t, err, "%s should not be read!", invalid)

		content, err := fs.ReadFile(valid)
		require.NoError(t, err)
		assert.NotEmpty(t, content, "%s should not be empty!", valid)
	})

	t.Run("Search", func(t *testing.T) {
		dir := "."
		phrase := `flexi`
		ignore := ""
		ext := "go"

		result, err := fs.Search(dir, phrase, ignore, ext)
		require.NoError(t, err)
		assert.NotNil(t, result, "Search should find a result!")
	})

	t.Run("Find", func(t *testing.T) {
		dir := "."
		pattern := ".*test.*"

		result, err := fs.Find(dir, pattern)
		require.NoError(t, err)
		assert.NotNil(t, result, "Find should find a result!")
	})

	t.Run("Lookup", func(t *testing.T) {
		dir := "."
		pattern := `.*\.go`

		results, err := fs.Lookup(dir, pattern)
		require.NoError(t, err)
		assert.NotEmpty(t, results, "Lookup should find results!")
	})

	t.Run("FS", func(t *testing.T) {
		assert.NotNil(t, fs.FS(), "FS should return a valid fs.FS instance!")
	})

	t.Run("Http", func(t *testing.T) {
		assert.NotNil(t, fs.Http(), "Http should return a valid http.FileSystem instance!")
	})
}
