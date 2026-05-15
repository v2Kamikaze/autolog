package controllers

import (
	"net/http"

	"github.com/gorilla/schema"
	"github.com/v2code/autolog/internal/persistence"
)

type AuthController struct {
	sessionPersistence persistence.SessionPersistence
	decoder            *schema.Decoder
}

func NewAuthController(sessionPersistence persistence.SessionPersistence) *AuthController {
	decoder := schema.NewDecoder()

	return &AuthController{decoder: decoder, sessionPersistence: sessionPersistence}
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {

}
