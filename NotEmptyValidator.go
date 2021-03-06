package ecms_validator

import (
	. "github.com/extensible-cms/ecms-go-validator/is"
)

type NotEmptyValidatorOptions struct {
	MessageFuncs *MessageFuncs
}

const (
	DefaultEmptyNotAllowedMsg = "Empty values are not allowed."
)

func NewNotEmptyValidatorOptions () NotEmptyValidatorOptions {
	return NotEmptyValidatorOptions{
		MessageFuncs: &MessageFuncs{
			EmptyNotAllowed: func(ops ValidatorOptions, x interface{}) string {
				return DefaultEmptyNotAllowedMsg
			},
		},
	}
}

func NotEmptyValidator (options ValidatorOptions) Validator {
	return func (x interface{}) (bool, []string){
		if Empty(x) {
			return false, []string{GetErrorMessageByKey(options, EmptyNotAllowed, x)}
		}
		return true, nil
	}
}

func NotEmptyValidator1 () Validator {
	return NotEmptyValidator(NewNotEmptyValidatorOptions())
}

func (n NotEmptyValidatorOptions) GetErrorMessageByKey (key int, x interface{}) string {
	return GetErrorMessageByKey(n, key, x)
}

func (n NotEmptyValidatorOptions) GetMessageFuncs () *MessageFuncs {
	return n.MessageFuncs
}

func (n NotEmptyValidatorOptions) GetValueObscurator () ValueObscurator {
	return DefaultValueObscurator
}
