// Package probe stores basic probes that are used to check services health
package probe

import (
	"fmt"
	"time"
)

// ProberConfig holds prober configuration, submitted through Init methods.
type ProberConfig struct {
	Target  string
	Options map[string]interface{}

	Warning time.Duration
	Fatal   time.Duration
}

// Prober is the base interface that each probe must implement.
type Prober interface {
	Init(ProberConfig) error
	Probe() (status Status, message string)
}

// Status represents the current status of a monitored service.
type Status string

// These constants represent the different available statuses of a service.
const (
	StatusUnknown Status = ""
	StatusWarning Status = "WARNING"
	StatusError   Status = "ERROR"
	StatusOK      Status = "OK"
)

const defaultConnectErrorMsg = "Unable to connect"

// EvaluateDuration is a shortcut for warning duration checks.
// It returns a message containing the duration, and a OK or a WARNING status
// depending on the provided warning duration.
func EvaluateDuration(duration time.Duration, warning time.Duration) (status Status, message string) {
	if duration >= warning {
		status = StatusWarning
	} else {
		status = StatusOK
	}

	message = fmt.Sprintf("%d ms", duration.Nanoseconds()/1000000)

	return
}
