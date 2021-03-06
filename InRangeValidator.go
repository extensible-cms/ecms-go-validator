package ecms_validator

import (
	"fmt"
	. "github.com/extensible-cms/ecms-go-validator/is"
	"reflect"
)

var DefaultInRangeMessageFuncs = MessageFuncs{
	NotAValidType: func(options ValidatorOptions, x interface{}) string {
		return fmt.Sprintf("%v is not a validatable numeric type.", x)
	},
}

type IntValidatorOptions struct {
	MessageFuncs *MessageFuncs
	Min              int64
	Max              int64
	Inclusive        bool
}

type FloatValidatorOptions struct {
	MessageFuncs *MessageFuncs
	Min              float64
	Max              float64
	Inclusive        bool
}

func NewIntRangeValidatorOptions() IntValidatorOptions {
	return IntValidatorOptions{
		MessageFuncs: &MessageFuncs{
			NotWithinRange: func(options ValidatorOptions, x interface{}) string {
				ops := options.(IntValidatorOptions)
				return fmt.Sprintf("%v is not within range %d and %d.", x, ops.Min, ops.Max)
			},
			NotAValidType: DefaultInRangeMessageFuncs[NotAValidType],
		},
		Min: 0,
		Max: 0,
		Inclusive: true,
	}
}

func NewFloatRangeValidatorOptions() FloatValidatorOptions {
	return FloatValidatorOptions{
		MessageFuncs: &MessageFuncs{
			NotWithinRange: func(options ValidatorOptions, x interface{}) string {
				ops := options.(FloatValidatorOptions)
				return fmt.Sprintf("%v is not within range %f and %f", x, ops.Min, ops.Max)
			},
			NotAValidType: DefaultInRangeMessageFuncs[NotAValidType],
		},
		Min: 0.0,
		Max: 0.0,
		Inclusive: true,
	}
}

func IntRangeValidator(options ValidatorOptions) Validator {
	ops := options.(IntValidatorOptions)
	return func(x interface{}) (bool, []string) {
		var intToCheck int64
		rv := reflect.ValueOf(x)
		switch rv.Kind() {
		case reflect.Invalid:
			return false, []string{ops.GetErrorMessageByKey(NotAValidType, x)}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			intToCheck = rv.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			intToCheck = int64(rv.Uint())
		default:
			return false, []string{ops.GetErrorMessageByKey(NotAValidType, x)}
		}
		if !IntWithinRange(ops.Min, ops.Max, intToCheck) {
			return false, []string{ops.GetErrorMessageByKey(NotWithinRange, x)}
		}
		return true, nil
	}
}

func FloatRangeValidator(options ValidatorOptions) Validator {
	ops := options.(FloatValidatorOptions)
	return func(x interface{}) (bool, []string) {
		rv := reflect.ValueOf(x)
		var floatToCheck float64
		switch rv.Kind() {
		case reflect.Invalid:
			return false, []string{ops.GetErrorMessageByKey(NotAValidType, x)}
		case reflect.Float32, reflect.Float64:
			floatToCheck = rv.Float()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			floatToCheck = float64(rv.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			floatToCheck = float64(rv.Uint())
		default:
			return false, []string{ops.GetErrorMessageByKey(NotAValidType, x)}
		}
		if !FloatWithinRange(ops.Min, ops.Max, floatToCheck) {
			return false, []string{ops.GetErrorMessageByKey(NotWithinRange, x)}
		}
		return true, nil
	}
}

func (n IntValidatorOptions) GetErrorMessageByKey(key int, x interface{}) string {
	return GetErrorMessageByKey(n, key, x)
}

func (n IntValidatorOptions) GetMessageFuncs() *MessageFuncs {
	return n.MessageFuncs
}

func (n IntValidatorOptions) GetValueObscurator() ValueObscurator {
	return DefaultValueObscurator
}

func (n FloatValidatorOptions) GetErrorMessageByKey(key int, x interface{}) string {
	return GetErrorMessageByKey(n, key, x)
}

func (n FloatValidatorOptions) GetMessageFuncs() *MessageFuncs {
	return n.MessageFuncs
}

func (n FloatValidatorOptions) GetValueObscurator() ValueObscurator {
	return DefaultValueObscurator
}
