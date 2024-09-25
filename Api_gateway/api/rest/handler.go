package rest

import (
	"encoding/json"
	"net/http"
)

type loginAndRegisterResponse struct {
	Token string `json:"token"`
}

func LoginHundler(w http.ResponseWriter, r *http.Request) {
	//get/validate the data from the request
	//send the data to the User service
	//return the to token

}
func RegisterHundler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(loginAndRegisterResponse{
		Token: "ldsjflsjflsdjflksjf",
	})
	return
}
