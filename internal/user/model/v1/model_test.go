package v1

import (
	"testing"

	_ "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/errors"
)

func TestLogin(t *testing.T) {
	/*
		testCases := []*TestCase{
			testCaseScannerLoginWithSuccess(),
			testCaseScannerLoginWithError(),
		}

		newApp := app.NewAppMock()

		for _, test := range testCases {
			test.Log(t)

			// streaming
			streaming := v1Streaming.NewStreamingMock()
			test.Streaming.Setup(streaming)

			model := NewModel(newApp, nil)
			result, err := model.Login(test.Arguments[0].(domain.IContext), test.Arguments[1].(string), test.Arguments[2].(string))

			assert.Equal(t, test.Expected[1] == nil, err == nil)    // check nil error
			assert.Equal(t, test.Expected[0] == nil, result == nil) // check nil result
			if test.Expected[0] != nil {
				assert.Equal(t, test.Expected[0], result) // check result object
			}
		}

	*/
}
