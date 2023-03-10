package recovery

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			err := recover()
			if err != nil {
				fmt.Println(err) // May be log this error? Send to sentry?

				jsonBody, _ := json.Marshal(map[string]string{
					"error": "There was an internal server error",
				})

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(jsonBody)
			}

		}()

		next.ServeHTTP(w, r)

	})
}

func EndApp() (w http.ResponseWriter, r *http.Request) {
	fmt.Println("End App")
	message := recover()
	fmt.Println(w, "Terjadi Error", message)

	return
}
