package healthchecks

type HealthCheck interface {
	Run() (live, ready bool, status interface{})
}

type HealthCheckFunc func() (live, ready bool, status string)

func (h HealthCheckFunc) Run() (live, ready bool, status string) {
	return h()
}

type HealthCheckNamespace map[string]HealthCheck

func (h HealthCheckNamespace) Run() (live, ready bool, status interface{}) {
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

var rootNamespace = HealthCheckNamespace{}

func Register(name string, check HealthCheck) {
	rootNamespace.Register(name, check)
}
