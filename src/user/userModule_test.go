package user_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/damocles217/user_service/src/user"
	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	r := user.CreateServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/test/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "data", w.Body.String())

}
