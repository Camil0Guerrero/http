package app_test

import (
	"testing"

	"github.com/Camil0Guerrero/http/internal/app"
	"github.com/Camil0Guerrero/http/tests"
	"github.com/stretchr/testify/require"
)

func TestBasicGET(t *testing.T) {
	t.Run("A basic GET petition", func(t *testing.T) {
		server := tests.CreateServer()
		defer server.Close()

		args := []string{"./http", "../../examples/get.http"}
		err := app.Run(args)
		require.NoError(t, err)
	})
}

func TestBasicPUT(t *testing.T) {
	t.Run("A basic PUT petition", func(t *testing.T) {
		server := tests.CreateServer()
		defer server.Close()

		args := []string{"./http", "../../examples/put.http"}
		err := app.Run(args)
		require.NoError(t, err)
	})
}
