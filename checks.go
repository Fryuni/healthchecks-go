package healthchecks

// HealthCheck defines an interface for executing a health check. The Run method returns booleans indicating 'live' and 'ready' status, and a status payload.
type HealthCheck interface {
	Run() (live, ready bool, status interface{})
}

// HealthCheckFunc is an adapter to allow the use of ordinary functions as health checks.
// It implements the HealthCheck interface.
type HealthCheckFunc func() (live, ready bool, status interface{})

// Run executes the health check function and returns its results.
func (h HealthCheckFunc) Run() (live, ready bool, status interface{}) {
	return h()
}

// HealthCheckNamespace is a collection of named health checks.
type HealthCheckNamespace map[string]HealthCheck

// Run executes all registered health checks in the namespace.
// It aggregates the live and ready statuses over all checks and returns a map of individual statuses.
func (h HealthCheckNamespace) Run() (live, ready bool, status interface{}) {
	if len(h) == 0 {
		return false, false, "no health checks defined, that is unhealthy"
	}
	live = true
	ready = true
	multiStatus := map[string]interface{}{}
	for name, check := range h {
		subLive, subReady, subStatus := check.Run()
		live = live && subLive
		ready = ready && subReady
		multiStatus[name] = subStatus
	}
	return live, ready, multiStatus
}

// Register adds a new health check with the given name to the namespace.
func (h HealthCheckNamespace) Register(name string, check HealthCheck) {
	h[name] = check
}

// RegisterFunc adds a new health check function with the given name to the namespace.
func (h HealthCheckNamespace) RegisterFunc(name string, check HealthCheckFunc) {
	h[name] = check
}

var rootNamespace = HealthCheckNamespace{}

// Register registers a health check in the root namespace.
func Register(name string, check HealthCheck) {
	rootNamespace.Register(name, check)
}

// RegisterFunc registers a health check function in the root namespace.
func RegisterFunc(name string, check HealthCheckFunc) {
	rootNamespace.RegisterFunc(name, check)
}
