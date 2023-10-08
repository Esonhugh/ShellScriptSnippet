package cmd_impel

import (
	"errors"
	"github.com/AlecAivazis/survey/v2"
	"github.com/Esonhugh/ShellScriptSnippet/core/defines"
	"github.com/Esonhugh/ShellScriptSnippet/utils/Ask"
	"github.com/Esonhugh/ShellScriptSnippet/utils/Database"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
	"os"
)

func InitDB() error {
	db, err := Database.Open()
	if err != nil {
		log.Errorf("InitDb with Error", err)
		return err
	}
	err = errors.Join(
		db.AutoMigrate(&defines.ShellSnippet{}),
		db.AutoMigrate(&defines.Config{}),
	)
	if err != nil {
		log.Errorf("DB AutoMigrate Failured %v", err)
	} else {
		Database.DB = db
		log.Debugf("DB process successfully")
	}
	return err
}

func FoundSnippet(name string) (r defines.ShellSnippet, e error) {
	r.Name = name
	result := Database.DB.First(&r)
	if result.Error != nil || result.RowsAffected == 0 {
		log.Errorf("Snippet name: %v not found", name)
		return defines.ShellSnippet{}, errors.Join(errors.New("result not found"), result.Error)
	}
	return r, nil
}

func SaveSnippet(sn defines.ShellSnippet) {
	if Database.DB.Model(&defines.ShellSnippet{}).
		Where("name = ?", sn.Name).
		Updates(&sn).RowsAffected == 0 {
		Database.DB.Create(&sn)
	}
}

func GetAllEnabledSnippet() (rs []defines.ShellSnippet, e error) {
	result := Database.DB.Where("is_enable = true").Find(&rs)
	if result.Error != nil {
		log.Errorf("Error when found enabled records: %v", result.Error)
		return nil, e
	}
	return rs, nil
}

func GetAllSnippets() (rs []defines.ShellSnippet, err error) {
	result := Database.DB.Find(&rs)
	if result.Error != nil {
		log.Errorf("Error when getting all records: %v", result.Error)
		return nil, result.Error
	}
	log.Debugf("retrun all snippets: %v", rs)
	return rs, nil
}

func DisableGormDebugging() {
	Database.DB.Logger.LogMode(logger.Silent)
}

func SelectSnippets() (rs []defines.ShellSnippet) {
	allrs, err := GetAllSnippets()
	if err != nil {
		log.Errorf("Find All Snippets error: %v", err)
		os.Exit(-1)
	}
	res := Ask.ChoiceMultiOfSlice(Ask.ConvertToChoiceAbleSlice(allrs))
	for _, v := range res {
		rs = append(rs, v.(defines.ShellSnippet))
	}
	log.Debugf("Selected shell snippets is %v", rs)
	return
}

func SelectSnippetsReturnStringSlice() (rs []string) {
	allrs, err := GetAllSnippets()
	if err != nil {
		log.Errorf("Find All Snippets error: %v", err)
		os.Exit(-1)
	}
	var Options []string
	for _, v := range allrs {
		Options = append(Options, v.Name)
	}
	prompt := &survey.MultiSelect{
		Message: "Target: ",
		Options: Options,
	}
	var result []string
	err = survey.AskOne(prompt, &result)
	if err != nil {
		log.Errorf("Ask for target failed with error: %v", err)
		os.Exit(137)
	}

	log.Debugf("Selected shell snippets is %v", result)
	return result
}

func SelectSnippet() (r defines.ShellSnippet) {
	allrs, err := GetAllSnippets()
	if err != nil {
		log.Errorf("Find All Snippets error: %v", err)
		os.Exit(-1)
	}
	r = Ask.ChoiceOneOfSlice(Ask.ConvertToChoiceAbleSlice(allrs)).(defines.ShellSnippet)
	log.Debugf("Selected shell snippets is %v", r)
	return
}
