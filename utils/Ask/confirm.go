package Ask

import (
	"github.com/AlecAivazis/survey/v2"
	log "github.com/sirupsen/logrus"
	"os"
)

const (
	IsYourInfoCorrect  = "您填写的信息是否正确 (Is your info correct?)"
	DoYouWannaContinue = "您想要继续吗 (Do you wanna continue?)"
)

// ForSure is func prompt and ask user to Confirm
// You can easily use it like this:
//
//	 if Ask.ForSure(Ask.IsYourInfoCorrect) {
//		// yes
//		} else {
//	 // no
//	 }
func ForSure(msg string) (name bool) {
	name = false
	prompt := &survey.Confirm{
		Message: msg,
	}
	err := survey.AskOne(prompt, &name)
	if err != nil {
		log.Errorf("user confirm got error %v", err)
		os.Exit(137)
	}
	return
}
