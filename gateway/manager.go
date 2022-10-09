package gateway

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/FANGDEI/simplebank/ecrypto"
	"github.com/FANGDEI/simplebank/store/cache"
	"github.com/FANGDEI/simplebank/store/local"
	"github.com/go-playground/validator/v10"
	"github.com/iris-contrib/middleware/cors"
	jardiniere "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

type Manager struct {
	handler   *iris.Application
	localer   *local.Manager
	cacher    *cache.Manager
	storer    *local.Store
	cryptoer  *ecrypto.Manager
	tokener   *jardiniere.Middleware
	validator *validator.Validate
}

func New() *Manager {
	return &Manager{
		handler:   iris.New(),
		localer:   C.Localer,
		cacher:    C.Cacher,
		storer:    C.Storer,
		cryptoer:  C.Cryptoer,
		validator: C.Validator,
	}
}

func (m *Manager) Run() error {
	err := m.load()
	if err != nil {
		return err
	}
	return m.handler.Run(iris.Addr(
		fmt.Sprintf("%s:%d", C.Host, C.Port),
	))
}

func (m *Manager) load() (err error) {
	err = m.loadPlugin()
	if err != nil {
		return
	}
	return m.loadRoute()
}

func (m *Manager) loadPlugin() error {
	m.loadToken()
	m.loadCORS()
	return nil
}

func (m *Manager) loadCORS() error {
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "GET", "OPTIONS", "DELETE"},
		MaxAge:           3600,
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	m.handler.UseRouter(crs)
	return nil
}

func (m *Manager) loadRoute() error {
	t := reflect.TypeOf(m)
	for i := 0; i < t.NumMethod(); i++ {
		f := t.Method(i)
		if strings.HasPrefix(f.Name, "Route") &&
			f.Type.NumOut() == 0 &&
			f.Type.NumIn() == 1 {
			log.Println("[GATEWAY] LOAD ROUTE:", f.Name)
			f.Func.Call([]reflect.Value{
				reflect.ValueOf(m),
			})
		}
	}
	return nil
}
