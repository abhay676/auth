package controllers

import (
	"encoding/json"
	"errors"
	"github.com/abhay676/auth/api/models"
	"github.com/abhay676/auth/api/responses"
	"github.com/abhay676/auth/api/utils"
	"io/ioutil"
	"net/http"
)

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = user.Validate("signup")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// check for existing user
	res, err := user.FindByEmail(s.DB, user.Email)
	if res != nil {
		responses.ERROR(w, http.StatusAlreadyReported, errors.New("email already exits"))
		return
	}

	// encrypt password
	user.Password, err = models.Hash(user.Password)

	// linking user-agent and ip with user
	user.UserAgent = r.UserAgent()
	user.IP = r.RemoteAddr

	//Generate UniqueID
	user.UserID = utils.GenerateID()

	// Generate token
	user.Token, err = utils.GenerateJWTToken(user.Email, user.UserID)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	// save new user
	newUser, err := user.Save(s.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, newUser)
}
