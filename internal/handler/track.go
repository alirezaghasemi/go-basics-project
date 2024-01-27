package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/alirezaghasemi/go-basics-project/api"
)

type TrackRepository interface {
	Upsert(ctx context.Context, track *api.Track) (err error)
	Get(ctx context.Context) (ts []api.Track, err error)
	GetById(ctx context.Context) (t api.Track, err error)
}

type TrackHandler struct {
	repo TrackRepository
}

func NewTrackHandler(repo TrackRepository) TrackHandler {
	return TrackHandler{
		repo: repo,
	}
}

func (h TrackHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

	} else if r.Method == "POST" {
		h.post(w, r)
	}
}

func (h TrackHandler) post(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	track := api.Track{}
	if err := json.NewDecoder(r.Body).Decode(&track); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if err := h.repo.Upsert(r.Context(), &track); err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("succeed"))
}

func (h TrackHandler) get(w http.ResponseWriter, r *http.Request) {
 vals := r.URL.Query()
 if vals.Has("track_id") {
	
 }
}