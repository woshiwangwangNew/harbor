package filter

import (
	"context"
	beegoctx "github.com/astaxie/beego/context"
	"github.com/goharbor/harbor/src/common/utils/log"
	"github.com/goharbor/harbor/src/core/config"
	"net/http"
)

// SessionReqKey is the key in the context of a request to mark the request carries session when reaching the backend
const SessionReqKey ContextValueKey = "harbor_with_session_req"

// SessionCheck is a filter to mark the requests that carries a session id, it has to be registered as
// "beego.BeforeStatic" because beego will modify the request after execution of these filters, all requests will
// appear to have a session id cookie.
func SessionCheck(ctx *beegoctx.Context) {
	req := ctx.Request
	_, err := req.Cookie(config.SessionCookieName)
	if err == nil {
		ctx.Request = req.WithContext(context.WithValue(req.Context(), SessionReqKey, true))
		log.Debugf("Mark the request as with-session: %s %s", req.Method, req.URL.RawPath)
	}
}

// ReqCarriesSession verifies if the request carries session when
func ReqCarriesSession(req *http.Request) bool {
	r, ok := req.Context().Value(SessionReqKey).(bool)
	return ok && r
}
