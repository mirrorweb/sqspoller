package sqspoller

import "sync/atomic"

// resetRunState sets the running and shuttingDown status values back to 0.
func (p *Poller) resetRunState() {
	atomic.SwapInt64(&p.running, 0)
	atomic.SwapInt64(&p.shuttingDown, 0)
}

// isRunning checks the running state of the poller
func (p *Poller) isRunning() bool {
	return atomic.LoadInt64(&p.running) == 1
}

// checkAndSetRunningStatus is called at the start of the Run method to check
// whether the poller is already running. If it is, the function returns the
// ErrAlreadyRunning error, else it sets the running status value to 1.
func (p *Poller) checkAndSetRunningStatus() error {
	if ok := atomic.CompareAndSwapInt64(&p.running, 0, 1); !ok {
		return ErrAlreadyRunning
	}
	return nil
}

// checkAndSetShuttingDownStatus is called at the start of any shutdown method
// to check whether the poller is already in the process of shutting down. If it
// is, the function returns the ErrAlreadyShuttingDown error, else it sets the
// shuttingDown value to 1.
func (p *Poller) checkAndSetShuttingDownStatus() error {
	if ok := atomic.CompareAndSwapInt64(&p.shuttingDown, 0, 1); !ok {
		return ErrAlreadyShuttingDown
	}
	return nil

}
