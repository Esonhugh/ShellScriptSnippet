package File

import (
	"io"
	"os"
)

// Copy file from src to dst
func Copy(src string, dst string) (err error) {
	// Open the source file for reading
	s, err := os.Open(src)
	if err != nil {
		return
	}
	defer s.Close()
	d, err := os.Create(dst)
	if err != nil {
		return
	}
	defer d.Close()
	_, err = io.Copy(d, s)
	return
}
