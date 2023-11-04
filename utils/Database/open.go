package Database

import (
	"github.com/Esonhugh/ShellScriptSnippet/utils/File"
	"github.com/glebarez/sqlite"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

const (
	DBLocSuffix = ".config/sss/conf.db"
)

var DBLocation = filepath.Join(os.Getenv("HOME"), DBLocSuffix)

var DB *gorm.DB

func Open() (*gorm.DB, error) {
	home := os.Getenv("HOME")
	if home == "" {
		home = os.TempDir()
	}
	DBLocation = filepath.Join(home, DBLocSuffix)
	if !File.PathExists(DBLocation) {
		log.Debugf("DBlocation %v is not exists", DBLocation)
		err := os.MkdirAll(filepath.Dir(DBLocation), os.ModeDir|0755)
		if err != nil {
			log.Debugf("Mkdir all failure %v", err)
		}
		f, _ := os.Create(DBLocation)
		f.Close()
		log.Debugf("Create DB file success")
	}
	return gorm.Open(sqlite.Open(DBLocation), &gorm.Config{})
}

func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Errorf("Close DB error: %v", err)
	}
	err = sqlDB.Close()
	if err != nil {
		log.Errorf("Close DB error: %v", err)
	}
}
