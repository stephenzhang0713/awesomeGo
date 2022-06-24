package main

import (
	"context"
	"github.com/google/uuid"
	"net/http"
)

// HelloWorld hello world handler
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	msgID := ""
	if m := r.Context().Value("msgId"); m != nil {
		if value, ok := m.(string); ok {
			msgID = value
		}
	}
	w.Header().Add("msgId", msgID)
	w.Write([]byte("Hello, world"))
}

func injectMsgID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msgID := uuid.New().String()
		ctx := context.WithValue(r.Context(), "msgId", msgID)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}

func main() {
	helloWorldHandler := http.HandlerFunc(HelloWorld)
	http.Handle("/welcome", injectMsgID(helloWorldHandler))
	http.ListenAndServe("localhost:8080", nil)
}
