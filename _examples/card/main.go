package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cjaewon/skillgo"
)

func basicCardRouter(w http.ResponseWriter, r *http.Request) {
	var payload skillgo.SkillPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // payload Error -> return EOF
		return
	}

	response := skillgo.SkillResponse{
		Version: "2.0",
	}

	thumbnail := skillgo.ThumbnailType{
		ImageURL: "http://k.kakaocdn.net/dn/83BvP/bl20duRC1Q1/lj3JUcmrzC53YIjNDkqbWK/i_6piz1p.jpg",
	}

	var buttons []skillgo.ButtonType

	buttons = append(buttons, skillgo.Button("열어보기", "message", skillgo.ActionConfig{
		MessageText: "짜잔! 우리가 찾던 보물입니다",
	}))
	buttons = append(buttons, skillgo.Button("구경하기", "webLink", skillgo.ActionConfig{
		WebLinkURL: "https://e.kakao.com/t/hello-ryan",
	}))

	response.Template.Outputs = append(
		response.Template.Outputs,
		skillgo.BasicCard("보물상자", "보물상자 안에는 무엇이 있을까", thumbnail, buttons),
	)

	m, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(m)
	w.WriteHeader(200)

}

func commerceCardRouter(w http.ResponseWriter, r *http.Request) {
	var payload skillgo.SkillPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // payload Error -> return EOF
		return
	}

	response := skillgo.SkillResponse{
		Version: "2.0",
	}

	var thumbnails []skillgo.ThumbnailType
	var buttons []skillgo.ButtonType

	thumbnails = append(thumbnails, skillgo.Thumbnail(
		"http://k.kakaocdn.net/dn/83BvP/bl20duRC1Q1/lj3JUcmrzC53YIjNDkqbWK/i_6piz1p.jpg",
		skillgo.Link("", "", "https://store.kakaofriends.com/kr/products/1542"),
		false,
		0,
		0,
	))

	profile := skillgo.Profile("보물상자 팝니다", "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT4BJ9LU4Ikr_EvZLmijfcjzQKMRCJ2bO3A8SVKNuQ78zu2KOqM")
	buttons = append(buttons,
		skillgo.Button("구매하기", "webLink", skillgo.ActionConfig{
			WebLinkURL: "https://store.kakaofriends.com/kr/products/1542",
		}),
		skillgo.Button("전화하기", "phone", skillgo.ActionConfig{
			PhoneNumber: "354-86-00070",
		}),
		skillgo.Button("공유하기", "share", skillgo.ActionConfig{}),
	)

	response.Template.Outputs = append(
		response.Template.Outputs,
		skillgo.CommerceCard("따끈따끈한 보물 상자 팝니다", 10000, "won", 1000, 0, 0, thumbnails, profile, buttons),
	)

	m, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(m)
	w.WriteHeader(200)
}

func listCardRouter(w http.ResponseWriter, r *http.Request) {
	var payload skillgo.SkillPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // payload Error -> return EOF
		return
	}

	response := skillgo.SkillResponse{
		Version: "2.0",
	}

	var items []skillgo.ListItemType

	items = append(
		items,
		skillgo.ListItem("Kakao i Developers", "새로운 AI의 내일과 일상의 변화", "http://k.kakaocdn.net/dn/APR96/btqqH7zLanY/kD5mIPX7TdD2NAxgP29cC0/1x1.jpg", skillgo.LinkType{
			Web: "https://namu.wiki/w/%EB%9D%BC%EC%9D%B4%EC%96%B8(%EC%B9%B4%EC%B9%B4%EC%98%A4%ED%94%84%EB%A0%8C%EC%A6%88)",
		}),
	)

	response.Template.Outputs = append(
		response.Template.Outputs,
		skillgo.ListCard(
			skillgo.ListItem("카카오 i 디벨로퍼스를 소개합니다", "", "http://k.kakaocdn.net/dn/xsBdT/btqqIzbK4Hc/F39JI8XNVDMP9jPvoVdxl1/2x1.jpg", skillgo.LinkType{}),
			items,
			[]skillgo.ButtonType{},
		),
	)

	m, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(m)
	w.WriteHeader(200)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/basic-card", basicCardRouter)
	mux.HandleFunc("/commerce-card", commerceCardRouter)
	mux.HandleFunc("/list-card", listCardRouter)

	log.Fatal(http.ListenAndServe(":3000", mux))
}
