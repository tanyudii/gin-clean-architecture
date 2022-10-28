package prompt

import (
	"errors"
	"github.com/manifoldco/promptui"
	"github.com/vodeacloud/hr-api/pkg/logger"
)

type PromptContent struct {
	Label    string
	ErrMsg   string
	Mask     rune
	Validate promptui.ValidateFunc
}

func (pc PromptContent) baseValidateString(val string) error {
	if len(val) == 0 {
		return errors.New(pc.ErrMsg)
	}
	return nil
}

func (pc PromptContent) GetString() string {
	validateFunc := pc.baseValidateString

	prompt := promptui.Prompt{
		Label:    pc.Label,
		Validate: validateFunc,
	}

	if pc.Validate != nil {
		prompt.Validate = pc.Validate
	}

	if pc.Mask != 0 {
		prompt.Mask = pc.Mask
	}

	result, err := prompt.Run()
	if err != nil {
		logger.Fatalf(err.Error())
	}
	return result
}
