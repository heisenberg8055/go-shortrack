package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/heisenberg8055/gotiny/internal/postgres"
	redis_client "github.com/heisenberg8055/gotiny/internal/redis-client"
	template "github.com/heisenberg8055/gotiny/internal/templates"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type InputURL struct {
	LongURL string `json:"longurl"`
}

type OutputURL struct {
	ShortURL string `json:"longurl"`
}

type Count struct {
	Urlcount int64 `json:"urlcount"`
}

func AddURL(w http.ResponseWriter, r *http.Request, postClient *pgxpool.Pool, redisClient *redis.Client, logger *slog.Logger) {

	// Method Check
	if r.Method != "POST" {
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
			redis_client.RedisSet(redisClient, longUrl, shortURL)
			template.RenderHomeShortLink(w, shortURL)
			return
		}
		for postgres.ValidateHash(postClient, shortURL) {
			shortURL = convertToHash(longUrl)
		}
		redis_client.RedisSet(redisClient, longUrl, shortURL)
		err := postgres.InsertData(postClient, longUrl, shortURL)
		if err != nil {
			http.Error(w, fmt.Sprintf("Unable to write to DB pool: %v", err), http.StatusInternalServerError)
			return
		}
		template.RenderHomeShortLink(w, shortURL)
		return
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func GetURL(w http.ResponseWriter, r *http.Request, postClient *pgxpool.Pool, redisClient *redis.Client, logger *slog.Logger) {
	shortURL := r.PathValue("shortUrl")
	if shortURL == "" {
		http.Error(w, "Wrong Request", http.StatusNotFound)
		return
	}
	longURL := ""
	longURL = redis_client.RedisGet(redisClient, shortURL)
	if longURL != "" {
		err := postgres.IncrementCount(postClient, shortURL)
		if err != nil {
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
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, longURL, http.StatusMovedPermanently)
		return
	}
	http.Error(w, "Bad Short URL", http.StatusBadRequest)
}

func validateURL(longUrl string) (bool, error) {
	_, err := url.ParseRequestURI(longUrl)
	return err == nil, err
}

func GetCount(w http.ResponseWriter, r *http.Request, conn *pgxpool.Pool, logger *slog.Logger) {
	shortURL := r.PathValue("shorturl")
	count := postgres.GetCount(conn, shortURL)
	if count == -1 {
		http.Error(w, "Unknown URL", http.StatusBadRequest)
		return
	}
	var response Count
	w.Header().Set("Content-Type", "application/json")
	response.Urlcount = count
	data, _ := json.Marshal(response)
	w.Write(data)
}

func Home(w http.ResponseWriter, r *http.Request, logger *slog.Logger) {
	template.RenderHome(w)
}
