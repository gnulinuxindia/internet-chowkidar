package tests

import (
	"context"
	"net/http"
	"strconv"
	"sync"
	"time"

	genapi "github.com/gnulinuxindia/internet-chowkidar/api/gen"
	"github.com/gnulinuxindia/internet-chowkidar/pkg/di"
	"go.uber.org/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var port int = 45820
var lock sync.Mutex

type TestSuiteBase struct {
	suite.Suite
	Ctx    context.Context
	Ctrl   *gomock.Controller
	Client *genapi.Client
	Server http.Server
}

func (s *TestSuiteBase) SetupTest() {
	s.Ctx = context.Background()
	s.Ctrl = gomock.NewController(s.T())

	handlers, err := di.InjectMockHandlers(s.Ctrl)
	if err != nil {
		panic(err)
	}

	server, err := genapi.NewServer(
		handlers,
	)
	if err != nil {
		panic(err)
	}

	/*
	* Ideally, the best way to do this is to not actually start a server
	* and just use `Serve` to handle the request, but the auto-generated
	* client code that gives us type-safe requests does not currently let
	* us do that, so this is the second best way.
	* TODO: figure out a way to not have to listen on a port
	 */
	var localPort int
	lock.Lock()
	localPort = port
	port = port + 1
	lock.Unlock()

	localPortStr := ":" + strconv.Itoa(localPort)

	s.Server = http.Server{
		ReadHeaderTimeout: time.Second,
		Addr:              localPortStr,
		Handler:           server,
	}

	client, err := genapi.NewClient("http://localhost" + localPortStr)
	if err != nil {
		panic(err)
	}
	s.Client = client

	go func() {
		s.Server.ListenAndServe()
	}()
}

func (s *TestSuiteBase) TearDownTest() {
	s.Ctrl.Finish()
	s.Server.Shutdown(s.Ctx)
}

// helpers
func (s *TestSuiteBase) RNil(object any, msgAndArgs ...any) {
	assert.Nil(s.T(), object, msgAndArgs...)
}

func (s *TestSuiteBase) RNotNil(object any, msgAndArgs ...any) {
	assert.NotNil(s.T(), object, msgAndArgs...)
}
