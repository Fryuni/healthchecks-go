package healthchecks

type HealthCheck interface {
	Run() (live, ready bool, status interface{})
}

type HealthCheckFunc func() (live, ready bool, status interface{})

func (h HealthCheckFunc) Run() (live, ready bool, status interface{}) {
	return h()
}

type HealthCheckNamespace map[string]HealthCheck

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

func (h HealthCheckNamespace) Register(name string, check HealthCheck) {
	h[name] = check
}

func (h HealthCheckNamespace) RegisterFunc(name string, check HealthCheckFunc) {
	h[name] = check
}

var rootNamespace = HealthCheckNamespace{}

func Register(name string, check HealthCheck) {
	rootNamespace.Register(name, check)
}

func RegisterFunc(name string, check HealthCheckFunc) {
	rootNamespace.RegisterFunc(name, check)
}
