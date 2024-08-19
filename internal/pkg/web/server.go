package web

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync/atomic"
	"time"
)

// ServerConfig represents configuration for Server.
type ServerConfig struct {
	Listen            string        `config:"listen"`
	DrainInterval     time.Duration `yaml:"drainInterval"`
	Profile           bool          `yaml:"profile"`
	ReadTimeout       time.Duration `yaml:"readTimeout"`
	ReadHeaderTimeout time.Duration `yaml:"readHeaderTimeout"`
	WriteTimeout      time.Duration `yaml:"writeTimeout"`
	IdleTimeout       time.Duration `yaml:"idleTimeout"`
	Env               string        `yaml:"env"`
}

// Server is an interface for web http server.
type Server interface {
	Run(ctx context.Context) error
	Shutdown(ctx context.Context) error
	Router() gin.IRouter
	Ready() bool
}

var _ Server = (*BaseServer)(nil)

// BaseServer is a default implementation of Server interface.
type BaseServer struct {
	engine     *gin.Engine
	httpServer *http.Server

	config ServerConfig

	isNotReady int32
}

// NewServer returns new *BaseServer.
func NewServer(config ServerConfig, handler *gin.Engine) *BaseServer {
	s := &BaseServer{
		engine: handler,
		config: config,
	}

	s.httpServer = &http.Server{
		Addr:              config.Listen,
		Handler:           s.engine.Handler(),
		ReadTimeout:       config.ReadTimeout,
		ReadHeaderTimeout: config.ReadHeaderTimeout,
		WriteTimeout:      config.WriteTimeout,
		IdleTimeout:       config.IdleTimeout,
	}

	return s
}

func (s *BaseServer) Run(ctx context.Context) error {
	s.Router().GET("/live", func(_ *gin.Context) {})
	s.Router().GET("/ping", s.getPing)

	go func() {
		for {
			<-ctx.Done()
			return
		}
	}()

	return s.httpServer.ListenAndServe()
}

func (s *BaseServer) Shutdown(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, s.config.DrainInterval)
	defer cancel()

	return s.httpServer.Shutdown(ctx)
}

func (s *BaseServer) Router() gin.IRouter {
	return s.engine
}

func (s *BaseServer) Ready() bool {
	return atomic.LoadInt32(&s.isNotReady) == 0
}

func (s *BaseServer) getPing(ctx *gin.Context) {
	if s.Ready() {
		_, _ = ctx.Writer.Write([]byte("pong"))
	} else {
		http.Error(ctx.Writer, "server cannot accept requests", http.StatusTeapot)
	}
}
