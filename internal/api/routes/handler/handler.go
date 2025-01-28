package handlers

import (
	"log/slog"
	"net/http"
	"net/url"

	log_middleware "github.com/heisenberg8055/gotiny/internal/log"
	"github.com/heisenberg8055/gotiny/internal/postgres"
	redis_client "github.com/heisenberg8055/gotiny/internal/redis-client"
	template "github.com/heisenberg8055/gotiny/internal/templates"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type InputURL struct {
	LongURL string `json:"longurl"`
}

func AddURL(w http.ResponseWriter, r *http.Request, postClient *pgxpool.Pool, redisClient *redis.Client, logger *slog.Logger) {

	// Method Check
	if r.Method != "POST" {
		log_middleware.LogError(log_middleware.Response{Method: r.Method, Url: r.URL.Path, Status: http.StatusMethodNotAllowed, Message: "Wrong Method Call"}, logger, "Failed POST CALL")
		http.Error(w, "Wrong Api Call Method", http.StatusMethodNotAllowed)
		return
	}

	currURL := InputURL{
		LongURL: r.FormValue("longurl"),
	}

	if ok, err := validateURL(currURL.LongURL); ok {
		longUrl := currURL.LongURL
		shortURL := ""
		shortURL = postgres.FetchShortUrl(postClient, longUrl)
		if shortURL != "" {
			redis_client.RedisSet(r, redisClient, longUrl, shortURL, logger)
			template.RenderHomeShortLink(w, r.Host+"/"+shortURL, r, logger)
			return
		}
		for postgres.ValidateHash(postClient, shortURL) {
			shortURL = convertToHash(longUrl)
		}
		redis_client.RedisSet(r, redisClient, longUrl, shortURL, logger)
		err := postgres.InsertData(postClient, longUrl, shortURL)
		if err != nil {
			log_middleware.LogError(log_middleware.Response{Method: r.Method, Url: r.URL.Path, Status: http.StatusInternalServerError, Message: err.Error()}, logger, "Unable to write to DB pool")
			return
		}
		template.RenderHomeShortLink(w, r.Host+"/"+shortURL, r, logger)
		return
	} else {
		log_middleware.LogError(log_middleware.Response{Method: r.Method, Url: r.URL.Path, Status: http.StatusBadRequest, Message: err.Error()}, logger, "Bad LongURL Input")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func GetURL(w http.ResponseWriter, r *http.Request, postClient *pgxpool.Pool, redisClient *redis.Client, logger *slog.Logger) {
	shortURL := r.PathValue("shortUrl")
	longURL := ""
	longURL = redis_client.RedisGet(redisClient, shortURL)
	if longURL != "" {
		err := postgres.IncrementCount(postClient, shortURL)
		if err != nil {
			log_middleware.LogError(log_middleware.Response{Method: r.Method, Url: r.URL.Path, Status: http.StatusBadRequest, Message: err.Error()}, logger, "Failed to Update Count for shortURL: "+shortURL)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, longURL, http.StatusMovedPermanently)
		return
	}
	longURL = postgres.FetchLongUrl(postClient, shortURL)
	if longURL != "" {
		err := postgres.IncrementCount(postClient, shortURL)
		if err != nil {
			log_middleware.LogError(log_middleware.Response{Method: r.Method, Url: r.URL.Path, Status: http.StatusBadRequest, Message: err.Error()}, logger, "Failed to Update Count for shortURL: "+shortURL)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, longURL, http.StatusMovedPermanently)
		return
	}
	template.RenderHomeError(w, r, logger)
}

func validateURL(longUrl string) (bool, error) {
	_, err := url.ParseRequestURI(longUrl)
	return err == nil, err
}

func GetCount(w http.ResponseWriter, r *http.Request, conn *pgxpool.Pool, logger *slog.Logger) {
	shortURL := r.FormValue("shorturl")
	count := postgres.GetCount(conn, shortURL)
	if count == -1 {
		log_middleware.LogWarn(log_middleware.Response{Method: r.Method, Url: r.URL.Path, Status: http.StatusBadRequest, Message: "Bad Short URL"}, logger, "Short URL is Not Registered")
		http.Error(w, "Unknown URL", http.StatusBadRequest)
		template.RenderHomeCount(w, shortURL, count, r, logger)
		return
	}
	template.RenderHomeCount(w, shortURL, count, r, logger)
}

func Home(w http.ResponseWriter, r *http.Request, logger *slog.Logger) {
	template.RenderHome(w, r, logger)
}
