package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/abhinav0031/scratch/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user %v", err))
		return
	}

	respondWthJson(w, 200, databaseUserToUser(user))
}

func (cfg *apiConfig) handleUsersGet(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWthJson(w, http.StatusOK, databaseUserToUser(user))
}

func (cfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := cfg.DB.GetPostsForUsers(r.Context(), database.GetPostsForUsersParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldn't get posts: %v", err))
		return
	}
	respondWthJson(w, 200, databasePostsToPosts(posts))

}
