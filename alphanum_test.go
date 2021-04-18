package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var validIDStrings = []struct {
	comment     string
	inputString string
	inputLength int
	output      bool
}{
	{
		comment:     "valid string",
		inputString: "fVrASaSRrpaWUYFx2w9l2VmOPMg9WAQRE7ODJqrtV7gXGSDhIr4cVicZ8EpLqWgU",
		inputLength: 64,
		output:      true,
	},
	{
		comment:     "empty string",
		inputString: "",
		inputLength: 64,
		output:      false,
	},
	{
		comment:     "invalid below 48 char / = 47",
		inputString: "/VrASaSRrpaWUYFx2w9l2VmOPMg9WAQRE7ODJqrtV7gXGSDhIr4cVicZ8EpLqWgU",
		inputLength: 64,
		output:      false,
	},
	{
		comment:     "invalid above 122 char { = 123",
		inputString: "{VrASaSRrpaWUYFx2w9l2VmOPMg9WAQRE7ODJqrtV7gXGSDhIr4cVicZ8EpLqWgU",
		inputLength: 64,
		output:      false,
	},
	{
		comment:     "invalid above 57 char : = 58",
		inputString: ":VrASaSRrpaWUYFx2w9l2VmOPMg9WAQRE7ODJqrtV7gXGSDhIr4cVicZ8EpLqWgU",
		inputLength: 64,
		output:      false,
	},
	{
		comment:     "invalid below 65 char @ = 64",
		inputString: "@VrASaSRrpaWUYFx2w9l2VmOPMg9WAQRE7ODJqrtV7gXGSDhIr4cVicZ8EpLqWgU",
		inputLength: 64,
		output:      false,
	},
	{
		comment:     "invalid above 90 char [ = 91",
		inputString: "[VrASaSRrpaWUYFx2w9l2VmOPMg9WAQRE7ODJqrtV7gXGSDhIr4cVicZ8EpLqWgU",
		inputLength: 64,
		output:      false,
	},
	{
		comment:     "invalid below 97 char ` = 96",
		inputString: "`VrASaSRrpaWUYFx2w9l2VmOPMg9WAQRE7ODJqrtV7gXGSDhIr4cVicZ8EpLqWgU",
		inputLength: 64,
		output:      false,
	},
}

func TestIsValidAlphaNum(t *testing.T) {

	for _, item := range validIDStrings {
		assert.Equal(t, item.output, IsValidAlphaNum(item.inputString, item.inputLength), item.comment)
	}
}
