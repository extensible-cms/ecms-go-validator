package ecms_validator

import "testing"

type NotEmptyValidatorTestCase struct {
	Name string
	Value interface {}
	ExpectedResultBln bool
	ExpectedMessagesLen int
}

func TestNotEmptyValidator(t *testing.T) {
	expectedErrorMessage := "Empty values are not allowed.  Received and empty value."
	validatorOptions := NotEmptyValidatorOptions{
		MessageFuncs: &MessageFuncs{
			EmptyNotAllowed: func(options ValidatorOptions, x interface{}) string {
				return expectedErrorMessage
			},
		},
	}
	validator := NotEmptyValidator(validatorOptions)
	for _, testCase := range []NotEmptyValidatorTestCase{
		{"validate_(nil)", nil, false, 1},
		{"validate_(0)", 0, false, 1},
		{"validate_(false)", false, false, 1},
		{"validate_([]string)", make([]string, 0), false, 1},
		{"validate_([]string{\"hello\"})", []string{"hello"}, true, 0},
		{"validate_(map[string]string{})", make(map[string]string), false, 1},
		{"validate_(map[string]string{\"hello\": \"world\"})", map[string]string{"hello": "world"}, true, 0},
		{"validate_(struct{}{})", struct{}{}, false, 1},
		{"validate_(struct{Name string}{\"hello\"})", struct{Name string}{"hello"}, true, 0},
	}{
		t.Run(testCase.Name, func(t *testing.T) {
			result, messages := validator(testCase.Value)
			messagesLen := len(messages)
			if result != testCase.ExpectedResultBln {
				t.Errorf("Expected %v for `result.Result` but got %v",
					testCase.ExpectedResultBln, result)
			}
			if messagesLen != testCase.ExpectedMessagesLen  {
				t.Errorf("Expected %d messages.  Got %d",
					testCase.ExpectedMessagesLen, messagesLen)
			}
			for _, message := range messages {
				if len(message) == 0 {
					t.Error("Expected non-empty message strings.  " +
						"Received an empty message string.")
				}
				if message != expectedErrorMessage {
					t.Errorf("Received unknown error message %v;  " +
						"\nExpected: %v", message, expectedErrorMessage)
				}
			}
		})
	}
}

func TestNotEmptyValidator1(t *testing.T) {
	expectedErrorMessage := DefaultEmptyNotAllowedMsg
	validator := NotEmptyValidator1()
	for _, testCase := range []NotEmptyValidatorTestCase{
		{"validate_(nil)", nil, false, 1},
		{"validate_(0)", 0, false, 1},
		{"validate_(false)", false, false, 1},
		{"validate_([]string)", make([]string, 0), false, 1},
		{"validate_([]string{\"hello\"})", []string{"hello"}, true, 0},
		{"validate_(map[string]string{})", make(map[string]string), false, 1},
		{"validate_(map[string]string{\"hello\": \"world\"})", map[string]string{"hello": "world"}, true, 0},
		{"validate_(struct{}{})", struct{}{}, false, 1},
		{"validate_(struct{Name string}{\"hello\"})", struct{Name string}{"hello"}, true, 0},
	}{
		t.Run(testCase.Name, func(t *testing.T) {
			result, messages := validator(testCase.Value)
			messagesLen := len(messages)
			if result != testCase.ExpectedResultBln {
				t.Errorf("Expected %v for `result.Result` but got %v",
					testCase.ExpectedResultBln, result)
			}
			if messagesLen != testCase.ExpectedMessagesLen  {
				t.Errorf("Expected %d messages.  Got %d",
					testCase.ExpectedMessagesLen, messagesLen)
			}
			for _, message := range messages {
				if len(message) == 0 {
					t.Error("Expected non-empty message strings.  " +
						"Received an empty message string.")
				}
				if message != expectedErrorMessage {
					t.Errorf("Received unknown error message %v;  " +
						"\nExpected: %v", message, expectedErrorMessage)
				}
			}
		})
	}
}
