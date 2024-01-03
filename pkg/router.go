package pkg

import (
	"github.com/NubeIO/lib-module-go/nhttp"
	"net/http"
)

func (m *Module) CallModule(method nhttp.Method, urlString string, headers http.Header, body []byte) ([]byte, error) {
	// TODO implement me
	panic("implement me")
}
