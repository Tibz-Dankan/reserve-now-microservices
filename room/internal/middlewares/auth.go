package middlewares

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Tibz-Dankan/reserve-now-microservices/room/internal/services"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		URL := "http://localhost:8000/api/v1/auth/verify"
		authorizationHeader := r.Header.Get("Authorization")

		req, err := http.NewRequest(http.MethodPost, URL, nil)
		if err != nil {
			services.AppError(err.Error(), 500, w)
			return
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", authorizationHeader)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			services.AppError(err.Error(), 500, w)
			return
		}

		type Response struct {
			Message string `json:"message"`
			Status  string `json:"status"`
		}

		if res.StatusCode != http.StatusOK {
			rBody, err := io.ReadAll(res.Body)
			if err != nil {

				services.AppError(err.Error(), 500, w)
				return
			}

			response := Response{}
			json.NewDecoder(strings.NewReader(string(rBody))).Decode(&response)
			services.AppError(response.Message, res.StatusCode, w)
			return
		}

		fmt.Printf("auth-service: status code: %d\n", res.StatusCode)
		resBody, _ := io.ReadAll(res.Body)
		fmt.Printf("auth-service: response body: %s\n", resBody)

		next.ServeHTTP(w, r)
	})
}
