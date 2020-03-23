package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("RequestPath: " + r.URL.Path + ", RequestMethod: " + r.Method)
		t := time.Now()
		next.ServeHTTP(w, r)
		fmt.Printf("Execution time: %s \n", time.Now().Sub(t).String())

	})
}
