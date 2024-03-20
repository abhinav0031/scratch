package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/abhinav0031/scratch/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (cnfg *apiConfig) handlerFeedFollowsGet(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := cnfg.DB.GetFeedFollowForUser(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get feed follows")
	}
	respondWthJson(w, http.StatusOK, databaseFeedsFollowToFeedsFollow(feedFollows))
}
func (cnfg *apiConfig) handlerFeedFollowCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId uuid.UUID
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode paramters")
		return
	}

	feedFolow, err := cnfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedId,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed follow")
		return
	}

	respondWthJson(w, http.StatusOK, databaseFeedFollowToFeedFollow(feedFolow))
}
func (cnfg *apiConfig) handlerFeedFollowDelete(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFolowIdStr := chi.URLParam(r, "feedFollowsId")
	feedFlowId, err := uuid.Parse(feedFolowIdStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid feed follows id")
		return
	}

	err = cnfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		ID:     feedFlowId,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't delete feed follow")
		return
	}

	respondWthJson(w, http.StatusOK, struct{}{})
}
