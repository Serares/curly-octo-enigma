package handlers

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/Serares/curly-octo-enigma/app/services"
	"github.com/Serares/curly-octo-enigma/app/views"
)

type AuthHandler struct {
	Logger      *slog.Logger
	AuthService *services.AuthService
}

func NewAuthHandler(logger *slog.Logger, service *services.AuthService) *AuthHandler {
	return &AuthHandler{
		Logger:      logger.WithGroup("Auth Handler"),
		AuthService: service,
	}
}

func (h *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	switch r.Method {
	case http.MethodGet:
		if path[1] == "login" {
			authUrl := h.AuthService.GenerateAuthUrl()
			viewLogin(w, r, views.LoginProps{
				AuthUrl:      authUrl,
				ErrorMessage: "",
			})
			return
		} else if path[1] == "callback" {
			q := r.URL.Query()
			h.Logger.Info("Got into the callback handler")
			if _, ok := q["code"]; ok {
				code := q["code"][0]
				rawToken, err := h.AuthService.Callback(r.Context(), code)
				if err != nil {
					viewLogin(w, r, views.LoginProps{
						AuthUrl:      "",
						ErrorMessage: err.Error(),
					})
				}
				cookie := http.Cookie{
					Name:  "id_token",
					Value: rawToken,
				}
				http.SetCookie(w, &cookie)
				http.Redirect(w, r, "/questions", http.StatusSeeOther)
			}
		}
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func viewLogin(w http.ResponseWriter, r *http.Request, props views.LoginProps) {
	views.Login(props).Render(r.Context(), w)
}
