package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/heisenberg8055/gotiny/internal/postgres"
	redis_client "github.com/heisenberg8055/gotiny/internal/redis-client"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type InputURL struct {
	LongURL string `json:"longurl"`
}

func AddURL(w http.ResponseWriter, r *http.Request, postClient *pgxpool.Pool, redisClient *redis.Client) {
	if r.Method != "POST" {
		http.Error(w, "Wrong Api Call Method", http.StatusMethodNotAllowed)
		return
	}
	ct := r.Header.Get("Content-Type")
	if ct != "" {
		mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
		if mediaType != "application/json" {
			msg := "Content-type is not applciation/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}

	//restricts body to 1MB
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var currURL InputURL

	err := dec.Decode(&currURL)

	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			http.Error(w, msg, http.StatusBadRequest)
			return
		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := "Request body contains badly-formed JSON"
			http.Error(w, msg, http.StatusBadRequest)
			return
		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			http.Error(w, msg, http.StatusBadRequest)
			return
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			http.Error(w, msg, http.StatusBadRequest)
			return
		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			http.Error(w, msg, http.StatusBadRequest)
			return
		case err.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			http.Error(w, msg, http.StatusRequestEntityTooLarge)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	if ok, err := validateURL(currURL.LongURL); ok {
		longUrl := currURL.LongURL
		var shortURL string = ""
		shortURL = redis_client.RedisGet(redisClient, longUrl)
		if shortURL != "" {
			w.Write([]byte(shortURL))
			w.WriteHeader(http.StatusNotModified)
			return
		}
		shortURL = postgres.FetchShortUrl(postClient, longUrl)
		if shortURL != "" {
			redis_client.RedisSet(redisClient, longUrl, shortURL)
			w.Write([]byte(shortURL))
			w.WriteHeader(http.StatusNotModified)
			return
		}
		for postgres.ValidateHash(postClient, shortURL) {
			shortURL = convertToHash(longUrl)
		}
		redis_client.RedisSet(redisClient, longUrl, shortURL)
		postgres.InsertData(postClient, longUrl, shortURL)
		w.Write([]byte(shortURL))
		w.WriteHeader(http.StatusCreated)
		return
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func GetURL(w http.ResponseWriter, r *http.Request) {
	shortURL := r.PathValue("shortUrl")
	if shortURL == "" {
		http.Error(w, "Wrong Request", http.StatusNotFound)
		return
	}

}

func validateURL(longUrl string) (bool, error) {
	_, err := url.ParseRequestURI(longUrl)
	return err == nil, err
}
