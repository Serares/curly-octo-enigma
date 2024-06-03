package middleware

import "net/http"

type MiddelwareCfgs func(*Middleware)

type Middleware struct {
	Next     http.Handler
	Secure   bool
	HTTPOnly bool
}

func NewMiddleware(next http.Handler, cfgs ...MiddelwareCfgs) http.Handler {
	m := Middleware{
		Next: next,
	}

	for _, cfg := range cfgs {
		cfg(&m)
	}

	return m
}

func WithSecure(s bool) MiddelwareCfgs {
	return func(m *Middleware) {
		m.Secure = s
	}
}

func CheckAuth(r *http.Request) (id string) {
	cookie, err := r.Cookie("id_token")
	if err != nil {
		return
	}

	return cookie.Value
}

func (mw Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := CheckAuth(r)
	if id == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	mw.Next.ServeHTTP(w, r)
}
