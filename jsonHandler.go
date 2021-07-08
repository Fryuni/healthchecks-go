package healthchecks

import (
	"encoding/json"
	"net/http"
)

type JsonHandler struct {
	check  HealthCheck
	Pretty bool
}

func NewJsonHandler(check HealthCheck) *JsonHandler {
	if check == nil {
		check = rootNamespace
	}
	if _, ok := check.(HealthCheckNamespace); !ok {
		check = HealthCheckNamespace{"status": check}
	}
	return &JsonHandler{check: check}
}

func (h *JsonHandler) Live() http.Handler {
	return h.handler(func(live, _ bool) bool {
		return live
	})
}

func (h *JsonHandler) Ready() http.Handler {
	return h.handler(func(_, ready bool) bool {
		return ready
	})
}

func (h *JsonHandler) LiveAndReady() http.Handler {
	return h.handler(func(live, ready bool) bool {
		return live && ready
	})
}

func (h *JsonHandler) handler(simplifier func(bool, bool) bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		live, ready, message := h.check.Run()
		ok := simplifier(live, ready)

		if ok {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
		}

		enc := json.NewEncoder(w)
		if h.Pretty {
			enc.SetIndent("", "  ")
		}
		_ = enc.Encode(message)
	})
}
