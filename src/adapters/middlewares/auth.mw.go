package middlewares

import (
    "net/http"
    bs "tuliospuri/go-clean/src/interactors"
    pres "tuliospuri/go-clean/src/adapters/presenters"
)

type authMw struct {
    authBs bs.AuthBs
}

func NewAuthMiddleware(authBs bs.AuthBs) authMw {
    return authMw{authBs}
}

func (m authMw) Execute(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        uri := r.
            Method + " " +
            r.URL.Path

        group := "guest"

        token := r.Header.
            Get("Authorization")

        ok, err := m.authBs.
            Authorized(uri, group, token)

        if ok {
            next.ServeHTTP(w, r)
        } else {
            jsonPres := pres.NewJsonPresenter()
            jsonPres.GetOutput(w, r, err)
        }
    })
}
