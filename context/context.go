package context

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err:= store.Fetch(r.Context())
		if err != nil {
			log.Println("error")
			return
		}
		fmt.Fprintf(w, data)
	}
}
