package provider

import (
	"fmt"
	"os"
	"path/filepath"
)

// FileClient provides an interface for CRUD operations on a file within a specific working directory.
type FileClient struct {
	workDir string
}

// NewFileClient creates and returns a new FileClient, ensuring the working directory exists.
func NewFileClient(workDir string) (*FileClient, error) {
	// Ensure the working directory exists, creating it if necessary.
	if err := os.MkdirAll(workDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create working directory %s: %w", workDir, err)
	}

	return &FileClient{
		workDir: workDir,
	}, nil
}

// resolvePath joins the working directory with the given file path to create a full path.
func (c *FileClient) resolvePath(path string) string {
	return filepath.Join(c.workDir, path)
}

// Create writes content to a new file within the working directory.
func (c *FileClient) Create(path, content string) error {
	fullPath := c.resolvePath(path)
	return os.WriteFile(fullPath, []byte(content), 0644)
}

// Read reads the content of an existing file within the working directory.
func (c *FileClient) Read(path string) (string, error) {
	fullPath := c.resolvePath(path)
	data, err := os.ReadFile(fullPath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Update overwrites an existing file with new content within the working directory.
func (c *FileClient) Update(path, newContent string) error {
	fullPath := c.resolvePath(path)
	return os.WriteFile(fullPath, []byte(newContent), 0644)
}

// Delete removes a file from the working directory.
func (c *FileClient) Delete(path string) error {
	fullPath := c.resolvePath(path)
	// We handle the case where the file might not exist to prevent a "no such file or directory" error.
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return nil // File doesn't exist, so there's nothing to delete.
	}
	return os.Remove(fullPath)
}
