package middelware

import "net/http"

// check for the Authorization header

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

func CheckAuthHeader(r *http.Request) (auth string) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return ""
	}

	return token
}

func (mw Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// token := CheckAuthHeader(r) // always checking for token
	// if token == "" {
	// 	http.Redirect(w, r, "/login", http.StatusSeeOther)
	// 	return
	// }
	mw.Next.ServeHTTP(w, r)
}
