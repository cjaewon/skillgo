package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cjaewon/skillgo"
)

func simpleTextRouter(w http.ResponseWriter, r *http.Request) {
	var payload skillgo.SkillPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // payload Error -> return EOF
		return
	}

	response := skillgo.SkillResponse{
		Version: "2.0",
	}
	response.Template.Outputs = append(response.Template.Outputs, skillgo.SimpleText("간단한 텍스트 요소입니다."))
	m, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(m)
	w.WriteHeader(200)

}

func simpleImageRouter(w http.ResponseWriter, r *http.Request) {
	var payload skillgo.SkillPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // payload Error -> return EOF
		return
	}

	response := skillgo.SkillResponse{
		Version: "2.0",
	}
	response.Template.Outputs = append(
		response.Template.Outputs,
		skillgo.SimpleImage("http://k.kakaocdn.net/dn/83BvP/bl20duRC1Q1/lj3JUcmrzC53YIjNDkqbWK/i_6piz1p.jpg", "보물상자입니다"),
	)

	m, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(m)
	w.WriteHeader(200)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/simple-text", simpleTextRouter)
	mux.HandleFunc("/simple-image", simpleImageRouter)

	log.Fatal(http.ListenAndServe(":3000", mux))
}
