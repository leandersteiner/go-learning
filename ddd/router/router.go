package router

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserHandler interface {
	IsUserSubscriptionActive(ctx context.Context, userId string) bool
}

type UserActiveResponse struct {
	IsActive bool
}

func router(u UserHandler) {
	r := chi.NewRouter()
	r.Get("/user/{userID}/subscription/active", func(w http.ResponseWriter, r *http.Request) {
		uID := chi.URLParam(r, "userID")
		if uID == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		isActive := u.IsUserSubscriptionActive(r.Context(), uID)

		b, err := json.Marshal(UserActiveResponse{IsActive: isActive})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = w.Write(b)
		if err != nil {
			fmt.Println(err)
			return
		}
	})
}
