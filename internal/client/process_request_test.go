package client_test

import (
	"net/http"
	"testing"

	"github.com/Camil0Guerrero/http/internal/client"
	"github.com/Camil0Guerrero/http/tests"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/require"
)

func TestBasicGetFlow(t *testing.T) {
	server := tests.CreateServer()
	defer server.Close()

	t.Run("should return 200 OK", func(t *testing.T) {
		url := server.URL + "/success"
		res, err := client.ProcessRequest("GET " + url)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, "everything is ok", res.Body)
	})

	t.Run("should return 404 Not Found", func(t *testing.T) {
		url := server.URL + "/not-found"
		res, err := client.ProcessRequest("GET " + url)
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
}

func TestBasicPutFlow(t *testing.T) {
	server := tests.CreateServer()
	defer server.Close()

	t.Run("should return a 201 Created", func(t *testing.T) {
		url := server.URL + "/create"

		res, err := client.ProcessRequest("PUT " + url + "\nContent-Type: application/json\n\n{\"name\": \"test\"}")
		require.NoError(t, err)

		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})

	t.Run("should return a 404 Not Found", func(t *testing.T) {
		url := server.URL + "/not-found"

		res, err := client.ProcessRequest("PUT " + url + "\nContent-Type: application/json\n\n{\"name\": \"test\"}")
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
}
