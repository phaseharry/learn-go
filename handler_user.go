package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/phaseharry/learn-go/internal/database"
)

/*
Making this function a method. The signature will stay the same since it needs to be
to be compatible with Go standard http library. Making this function a method however,
will give it access to additional data it might not normally have access to. In this case,
since handlerCreateUser is a method part of apiCfg, we have access to its queries.
*/
func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON payload: %v", err))
		return
	}
	currentTimeUtc := time.Now().UTC()

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: currentTimeUtc,
		UpdatedAt: currentTimeUtc,
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}
	respondWithJson(w, 200, user)
}
