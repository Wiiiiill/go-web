package web

import (
	"net/http"

	gin "github.com/Wiiiiill/go-web"
)

const ReleaseMode = "release"
const DebugMode = "debug"

type Context struct {
	UserID uint64
	*gin.Context
}

func MakeContext(oc *gin.Context) *Context {
	temp := &Context{0, oc}
	return temp
}

type App struct {
	*gin.Engine
}

func CreateApp() *App {
	router := gin.Default()
	temp := &App{router}
	return temp
}

func SetMode(mode string) {
	gin.SetMode(mode)
}

type Data interface{}
type Callback func(ctx *Context) (Data, error)

// POST is a shortcut for router.Handle("POST", path, handlers).
func (group *App) POST(relativePath string, handlers Callback) {
	group.Handle(http.MethodPost, relativePath, func(ctx *gin.Context) {
		handlers(MakeContext(ctx))
	})
}

// GET is a shortcut for router.Handle("GET", path, handlers).
func (group *App) GET(relativePath string, handlers Callback) {
	group.Handle(http.MethodGet, relativePath, func(ctx *gin.Context) {
		handlers(MakeContext(ctx))
	})
}

// DELETE is a shortcut for router.Handle("DELETE", path, handlers).
func (group *App) DELETE(relativePath string, handlers Callback) {
	group.Handle(http.MethodDelete, relativePath, func(ctx *gin.Context) {
		handlers(MakeContext(ctx))
	})
}

// PATCH is a shortcut for router.Handle("PATCH", path, handlers).
func (group *App) PATCH(relativePath string, handlers Callback) {
	group.Handle(http.MethodPatch, relativePath, func(ctx *gin.Context) {
		handlers(MakeContext(ctx))
	})
}

// PUT is a shortcut for router.Handle("PUT", path, handlers).
func (group *App) PUT(relativePath string, handlers Callback) {
	group.Handle(http.MethodPut, relativePath, func(ctx *gin.Context) {
		handlers(MakeContext(ctx))
	})
}

// OPTIONS is a shortcut for router.Handle("OPTIONS", path, handlers).
func (group *App) OPTIONS(relativePath string, handlers Callback) {
	group.Handle(http.MethodOptions, relativePath, func(ctx *gin.Context) {
		handlers(MakeContext(ctx))
	})
}

// HEAD is a shortcut for router.Handle("HEAD", path, handlers).
func (group *App) HEAD(relativePath string, handlers Callback) {
	group.Handle(http.MethodHead, relativePath, func(ctx *gin.Context) {
		handlers(MakeContext(ctx))
	})
}
