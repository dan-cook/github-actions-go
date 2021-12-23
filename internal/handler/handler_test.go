package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github-actions-go/internal/handler"

	"github.com/stretchr/testify/assert"
)

func TestServe(t *testing.T) {
	type expected struct {
		code int
		body string
	}

	tests := []struct {
		name     string
		request  *http.Request
		expected expected
	}{
		{
			name:    "200 ok - happy path",
			request: httptest.NewRequest(http.MethodGet, "/ping", nil),
			expected: expected{
				code: http.StatusOK,
				body: "Pong!",
			},
		},
		{
			name:    "405 method not allowed - error",
			request: httptest.NewRequest(http.MethodPost, "/ping", nil),
			expected: expected{
				code: http.StatusMethodNotAllowed,
			},
		},
	}

	handler := handler.NewPingHandler()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, test.request)

			assert.Equal(t, test.expected.code, rr.Code)
			assert.Equal(t, test.expected.body, rr.Body.String())
		})
	}
}
