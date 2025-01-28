package template

import (
	"html/template"
	"log/slog"
	"net/http"

	log_middleware "github.com/heisenberg8055/gotiny/internal/log"
)

func RenderHome(w http.ResponseWriter, r *http.Request, logger *slog.Logger) {
	tp, err := template.ParseFiles("static/index.html")
	if err != nil {
		log_middleware.LogError(log_middleware.Response{Method: r.Method, Url: r.URL.Path, Status: http.StatusInternalServerError, Message: err.Error()}, logger, "Failed to render Home")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tp.Execute(w, nil)
}

func RenderHomeShortLink(w http.ResponseWriter, shorturl string, r *http.Request, logger *slog.Logger) {
	t, err := template.ParseFiles("static/index.html")
	if err != nil {
		log_middleware.LogError(log_middleware.Response{Method: r.Method, Url: r.URL.Path, Status: http.StatusInternalServerError, Message: err.Error()}, logger, "Failed to render Home with Short URL: "+shorturl)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, shorturl)
}

func RenderHomeError(w http.ResponseWriter, r *http.Request, logger *slog.Logger) {
	t, err := template.ParseFiles("static/404.html")
	if err != nil {
		log_middleware.LogError(log_middleware.Response{Method: r.Method, Url: r.URL.Path, Status: http.StatusInternalServerError, Message: err.Error()}, logger, "Failed to render Home with Invalid Short URL")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	t.Execute(w, nil)
}

func RenderHomeCount(w http.ResponseWriter, shortURL string, count int64, r *http.Request, logger *slog.Logger) {
	t, err := template.ParseFiles("static/count.html")
	if err != nil {
		log_middleware.LogError(log_middleware.Response{Method: r.Method, Url: r.URL.Path, Status: http.StatusInternalServerError, Message: err.Error()}, logger, "Failed to render Home with Short URL Count: "+shortURL)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, count)
}
