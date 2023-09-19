package core

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestFilename(t *testing.T) {
	fn := "/Base/Helloworld.php"
	t.Log("Filename :", fn)
	fileName := filepath.Base(fn)
	baseName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	t.Log(baseName)
}
