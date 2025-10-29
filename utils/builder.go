package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// RepoBuilder helps create files and dirs without
// checking for errors after every single operation.
type RepoBuilder struct {
	basePath string
	err      error
}

// NewBuilder creates a new RepoBuilder and the base directory.
func NewBuilder(basePath string) *RepoBuilder {
	b := &RepoBuilder{basePath: basePath}

	// Create the root directory for the repo
	err := os.MkdirAll(basePath, 0755)
	if err != nil {
		b.err = fmt.Errorf("failed to create base directory %s: %w", basePath, err)
	}
	return b
}

// checkErr is an internal helper to stop if an error has occurred.
func (b *RepoBuilder) checkErr() bool {
	return b.err != nil
}

// MkDir creates a new directory.
// It stops execution if an error has already occurred.
func (b *RepoBuilder) MkDir(name string) {
	if b.checkErr() {
		return
	}
	// We use MkdirAll to create parent directories if needed.
	err := os.MkdirAll(filepath.Join(b.basePath, name), 0755)
	if err != nil {
		b.err = fmt.Errorf("failed to create dir %s: %w", name, err)
	}
}

// WriteFile creates a new file with the given content.
// It stops execution if an error has already occurred.
func (b *RepoBuilder) WriteFile(name, content string) {
	if b.checkErr() {
		return
	}
	// Ensure the directory for the file exists
	dir := filepath.Dir(name)
	if dir != "." {
		b.MkDir(dir)
		if b.checkErr() {
			return
		}
	}
	// Write the file
	fullPath := filepath.Join(b.basePath, name)
	err := os.WriteFile(fullPath, []byte(strings.TrimSpace(content)), 0644)
	if err != nil {
		b.err = fmt.Errorf("failed to write file %s: %w", name, err)
	}
}

// Err returns the first error encountered by the builder, if any.
func (b *RepoBuilder) Err() error {
	return b.err
}
