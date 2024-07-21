package http

import (
	"context"
	"net/http"

	"github.com/mat285/auth"
)

type Handlers struct {
	Register http.HandlerFunc
	Login    http.HandlerFunc
	Refresh  http.HandlerFunc
}

type Callbacks struct {
	PreRegister  func(context.Context, auth.Registration) (auth.Registration, error)
	PostRegister func(context.Context, auth.Identity) error
	PostLogin    func(context.Context, auth.AuthValue) error
}

func GetHandlers(m *auth.Manager, callbacks Callbacks) Handlers {
	return Handlers{
		Register: RegisterHandler(m, callbacks.PreRegister, callbacks.PostRegister),
		Login:    LoginHandler(m, callbacks.PostLogin),
	}
}

func RegisterHandler(m *auth.Manager, dp func(context.Context, auth.Registration) (auth.Registration, error), cb func(context.Context, auth.Identity) error) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var in auth.Registration
		err := ReadJSONBody(r, &in)
		if err != nil {
			writeError(rw, err)
			return
		}
		reg, err := dp(r.Context(), in)
		if err != nil {
			writeError(rw, err)
			return
		}

		ident, err := m.Register(r.Context(), reg, reg.Data)
		if err != nil {
			writeError(rw, err)
			return
		}

		err = cb(r.Context(), *ident)
		if err != nil {
			writeError(rw, err)
			return
		}
		WriteJSONBody(rw, ident)
	}
}

func LoginHandler(m *auth.Manager, cb func(context.Context, auth.AuthValue) error) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var reg auth.Creds
		err := ReadJSONBody(r, &reg)
		if err != nil {
			writeError(rw, err)
			return
		}
		_, av, err := m.Login(r.Context(), reg)
		if err != nil {
			writeError(rw, err)
			return
		}

		err = cb(r.Context(), av)
		if err != nil {
			writeError(rw, err)
			return
		}
		WriteJSONBody(rw, av)
	}
}

func writeError(rw http.ResponseWriter, err error) {
	if err == nil {
		return
	}
	rw.WriteHeader(400)
	rw.Write([]byte(err.Error()))
	return
}
