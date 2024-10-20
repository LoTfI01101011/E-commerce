package rest

import (
	"encoding/json"
	"net/http"

	"github.com/LoTfI01101011/E-commerce/Api_gateway/api/gRPC"
)

func LoginHundler(w http.ResponseWriter, r *http.Request, user *gRPC.User) {
	//get/validate the data from the request
	var body struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6"`
	}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//send the data to the User service
	res, err := user.LoginUser(body.Email, body.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonErr := map[string]error{"error": err}
		json.NewEncoder(w).Encode(jsonErr)
	}
	//return the to token
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonRes := map[string]string{"Token": res}
	json.NewEncoder(w).Encode(jsonRes)

}
func RegisterHundler(w http.ResponseWriter, r *http.Request, user *gRPC.User) {
	var body struct {
		UserName        string `json:"username" validate:"required"`
		Email           string `json:"email" validate:"required,email"`
		Password        string `json:"password" validate:"required,min=8"`
		ConfirmPassword string `json:"confirm_password" validate:"required,min=8"`
	}
	//get the data from the body
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]string{"response": "Invalid request body"}
		json.NewEncoder(w).Encode(response)
		return
	}
	//check if the passwrod equals the confirm pass
	if body.Password != body.ConfirmPassword {
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]string{"response": "Invalid request"}
		json.NewEncoder(w).Encode(response)
		return
	}
	//send it to the User service and get the response
	response, err := user.RegisterUser(body.UserName, body.Email, body.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]error{"error": err}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	res := map[string]string{"Response": response}
	json.NewEncoder(w).Encode(res)
	//return the response
}
func LogoutHundler(w http.ResponseWriter, r *http.Request, user *gRPC.User) {
	//get the token from the request header
	token := r.Header.Get("Authorization")

	//call the logout function
	res, err := user.LogoutUser(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]error{"error": err}
		json.NewEncoder(w).Encode(response)
		return
	}

	jsonRes := map[string]string{"response": res}
	json.NewEncoder(w).Encode(jsonRes)
}
