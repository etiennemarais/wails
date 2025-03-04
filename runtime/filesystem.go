package runtime

import "os"

// FileSystem exposes file system utilities to the runtime
type FileSystem struct{}

// NewFileSystem creates a new FileSystem struct
func NewFileSystem() *FileSystem {
	return &FileSystem{}
}

// HomeDir returns the user's home directory
func (r *FileSystem) HomeDir() (string, error) {
	return os.UserHomeDir()
}

// UserCacheDir returns the user's application cache directory
func (r *FileSystem) UserCacheDir() (string, error) {
	return os.UserCacheDir()
}
