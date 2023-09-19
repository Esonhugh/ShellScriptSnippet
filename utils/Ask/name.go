package Ask

import (
	"github.com/AlecAivazis/survey/v2"
	log "github.com/sirupsen/logrus"
	"os"
)

func ForName(msg string) (name string) {
	prompt := &survey.Input{
		Message: msg,
	}
	err := survey.AskOne(prompt, &name)
	if err != nil {
		log.Errorf("user confirm got error %v", err)
		os.Exit(137)
	}
	return
}
