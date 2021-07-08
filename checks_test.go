package healthchecks

import (
	"reflect"
	"testing"
)

func TestHealthCheckFunc_Run(t *testing.T) {
	sentinelStatus := "sentinel-status"
	var healthCheckFunc = HealthCheckFunc(func() (bool, bool, interface{}) {
		return false, false, sentinelStatus
	})

	if live, ready, status := healthCheckFunc.Run(); live || ready || status != sentinelStatus {
		t.Errorf("registered health check does not match given function")
	}
}

func TestHealthCheckNamespace_Register(t *testing.T) {
	sentinelStatus := "sentinel-status"
	var healthCheckFunc = func() (bool, bool, interface{}) {
		return false, false, sentinelStatus
	}

	name := "my register func"

	namespace := HealthCheckNamespace{}

	namespace.Register(name, HealthCheckFunc(healthCheckFunc))

	if live, ready, status := namespace[name].Run(); live || ready || status != sentinelStatus {
		t.Errorf("registered health check does not match given function")
	}
}

func TestHealthCheckNamespace_RegisterFunc(t *testing.T) {
	sentinelStatus := "sentinel-status"
	var healthCheckFunc = func() (bool, bool, interface{}) {
		return false, false, sentinelStatus
	}

	name := "my register func"

	namespace := HealthCheckNamespace{}

	namespace.RegisterFunc(name, healthCheckFunc)

	if live, ready, status := namespace[name].Run(); live || ready || status != sentinelStatus {
		t.Errorf("registered health check does not match given function")
	}
}

func TestHealthCheckNamespace_Run(t *testing.T) {
	tests := []struct {
		name       string
		h          HealthCheckNamespace
		wantLive   bool
		wantReady  bool
		wantStatus interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLive, gotReady, gotStatus := tt.h.Run()
			if gotLive != tt.wantLive {
				t.Errorf("Run() gotLive = %v, want %v", gotLive, tt.wantLive)
			}
			if gotReady != tt.wantReady {
				t.Errorf("Run() gotReady = %v, want %v", gotReady, tt.wantReady)
			}
			if !reflect.DeepEqual(gotStatus, tt.wantStatus) {
				t.Errorf("Run() gotStatus = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}

func TestRegister(t *testing.T) {
	sentinelStatus := "sentinel-status"
	var healthCheckFunc = func() (bool, bool, interface{}) {
		return false, false, sentinelStatus
	}

	name := "my check func"

	Register(name, HealthCheckFunc(healthCheckFunc))

	if live, ready, status := rootNamespace[name].Run(); live || ready || status != sentinelStatus {
		t.Errorf("registered health check does not match given function")
	}
}

func TestRegisterFunc(t *testing.T) {
	sentinelStatus := "sentinel-status"
	var healthCheckFunc = func() (bool, bool, interface{}) {
		return false, false, sentinelStatus
	}

	name := "my check func"

	RegisterFunc(name, healthCheckFunc)

	if live, ready, status := rootNamespace[name].Run(); live || ready || status != sentinelStatus {
		t.Errorf("registered health check does not match given function")
	}
}
