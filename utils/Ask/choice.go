package Ask

import (
	"github.com/AlecAivazis/survey/v2"
	log "github.com/sirupsen/logrus"
	"os"
	"sss/utils/Compare"
)

type ChoiceAble interface {
	UniqName() string
}
type ChoiceAbleSlice []ChoiceAble

func ConvertToChoiceAbleSlice[T ChoiceAble](cs []T) ChoiceAbleSlice {
	var result ChoiceAbleSlice
	for _, v := range cs {
		result = append(result, v)
	}
	return result
}

func ChoiceOneOfSlice(choice ChoiceAbleSlice) ChoiceAble {
	var Options []string
	for _, v := range choice {
		Options = append(Options, v.UniqName())
	}
	prompt := &survey.Select{
		Message: "Target: ",
		Options: Options,
	}
	var result string
	err := survey.AskOne(prompt, &result)
	for _, v := range choice {
		if v.UniqName() == result {
			return v
		}
	}
	if err != nil {
		log.Errorf("user select returns error %v", err)
		os.Exit(137)
	}
	return nil
}

func ChoiceMultiOfSlice(choices ChoiceAbleSlice) ChoiceAbleSlice {
	var Options []string
	for _, v := range choices {
		Options = append(Options, v.UniqName())
	}
	prompt := &survey.MultiSelect{
		Message: "Target: ",
		Options: Options,
	}
	var result []string
	err := survey.AskOne(prompt, &result)

	var resultSlice []ChoiceAble
	for _, v := range choices {
		// ready for select data
		if Compare.IsStringInStringArray(v.UniqName(), result) {
			resultSlice = append(resultSlice, v)
		}
	}
	if err != nil {
		log.Errorf("user select returns error %v", err)
		os.Exit(137)
	}
	return resultSlice
}
