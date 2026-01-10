package shutdown

import "sync"

// Callback shutdown 回调
type Callback interface {
	OnShutdown(string) error
}

// Hook is a helper type, so you can easily provide anonymous functions
type Hook func(string) error

func (f Hook) OnShutdown(shutdownManager string) error {
	return f(shutdownManager)
}

// ErrHandler 错误处理
type ErrHandler interface {
	OnErr(err error)
}

type ErrFunc func(err error)

func (f ErrFunc) OnErr(err error) {
	f(err)
}

type Manager interface {
	GetName() string
	Start(gs GSInterface) error
	ShutdownStart() error
	ShutdownFinish() error
}

type GSInterface interface {
	StartShutdown(sm Manager)
	ReportErr(err error)
	AddCallback(callback Callback)
}

type GracefulShutdown struct {
	callbacks  []Callback
	managers   []Manager
	errHandler ErrHandler
}

func New() *GracefulShutdown {
	return &GracefulShutdown{
		callbacks: make([]Callback, 0, 10),
		managers:  make([]Manager, 0, 3),
	}
}

func (gs *GracefulShutdown) AddManager(manager Manager) {
	gs.managers = append(gs.managers, manager)
}

func (gs *GracefulShutdown) AddCallback(callback Callback) {
	gs.callbacks = append(gs.callbacks, callback)
}

func (gs *GracefulShutdown) SetErrHandler(handler ErrHandler) {
	gs.errHandler = handler
}

func (gs *GracefulShutdown) StartShutdown(sm Manager) {
	gs.ReportErr(sm.ShutdownStart())

	var wg sync.WaitGroup
	for _, callback := range gs.callbacks {
		wg.Add(1)
		go func(callback Callback) {
			defer wg.Done()
			gs.ReportErr(callback.OnShutdown(sm.GetName()))
		}(callback)
	}
	wg.Wait()

	gs.ReportErr(sm.ShutdownFinish())
}

func (gs *GracefulShutdown) ReportErr(err error) {
	if err != nil && gs.errHandler != nil {
		gs.errHandler.OnErr(err)
	}
}
