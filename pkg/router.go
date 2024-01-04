package pkg

import (
	"encoding/json"
	"errors"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/lib-module-go/nmodule"
	"github.com/NubeIO/lib-module-go/router"
	"github.com/NubeIO/module-core-system/schema"
	"net/http"
)

var route *router.Router

func (m *Module) CallModule(method nhttp.Method, urlString string, headers http.Header, body []byte) ([]byte, error) {
	mod := (nmodule.Module)(m)
	return route.CallHandler(&mod, method, urlString, headers, body)
}

func InitRouter() {
	route = router.NewRouter()

	route.Handle(nhttp.GET, "/api/networks/schema", GetNetworkSchema)
	route.Handle(nhttp.GET, "/api/devices/schema", GetDeviceSchema)
	route.Handle(nhttp.GET, "/api/points/schema", GetPointSchema)

	route.Handle(nhttp.GET, "/api/system/schedule/store/name/:name", GetSystemScheduleStore)
}

func GetNetworkSchema(m *nmodule.Module, r *router.Request) ([]byte, error) {
	return json.Marshal(schema.GetNetworkSchema())
}

func GetDeviceSchema(m *nmodule.Module, r *router.Request) ([]byte, error) {
	return json.Marshal(schema.GetDeviceSchema())
}

func GetPointSchema(m *nmodule.Module, r *router.Request) ([]byte, error) {
	return json.Marshal(schema.GetPointSchema())
}

func GetSystemScheduleStore(m *nmodule.Module, r *router.Request) ([]byte, error) {
	obj, ok := (*m).(*Module).store.Get(r.PathParams["name"])
	if !ok {
		return nil, errors.New("no schedule exists")
	}
	return json.Marshal(obj)
}
