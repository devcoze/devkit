package managers

import (
	"os"
	"os/signal"

	"github.com/devcoze/devkit/shutdown"
)

// POSIX signal

const Name = "PosixSignalManager"

// PosixSignalManager implements ShutdownManager interface that is added
// to GracefulShutdown. Initialize with NewPosixSignalManager.
type PosixSignalManager struct {
	signals []os.Signal
}

func NewPosixSignalManager(signals ...os.Signal) *PosixSignalManager {
	if len(signals) == 0 {
		signals = []os.Signal{os.Interrupt, os.Kill}
	}
	return &PosixSignalManager{
		signals: signals,
	}
}

func (m *PosixSignalManager) GetName() string {
	return Name
}

func (m *PosixSignalManager) Start(gs shutdown.GSInterface) error {
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, m.signals...)

		<-c
		// Signal received, start shutdown
		gs.StartShutdown(m)
	}()
	return nil
}

func (m *PosixSignalManager) ShutdownStart() error {
	return nil
}

func (m *PosixSignalManager) ShutdownFinish() error {
	return nil
}
