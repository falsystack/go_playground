package app

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type GoogleUserId struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
}

func init() {
	googleOauthConfig = oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:8080/auth/google/callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	}
}

var googleOauthConfig oauth2.Config

func googleAtuCallback(w http.ResponseWriter, req *http.Request) {
	oauthstate, _ := req.Cookie("oauthstate")

	if req.FormValue("state") != oauthstate.Value {
		log.Printf("Invalid google oauth state cookie : %s state : %s", oauthstate.Value, req.FormValue("state"))
		http.Error(w, fmt.Sprintf("Invalid google oauth state cookie : %s state : %s", oauthstate.Value, req.FormValue("state")), http.StatusInternalServerError)
		return
	}

	data, err := getGoogleUserInfo(req.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Store Id info into Session cookie
	var userInfo GoogleUserId
	err = json.Unmarshal(data, &userInfo)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	session, _ := store.Get(req, "session")
	session.Values["id"] = userInfo.ID
	err = session.Save(req, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, req, "/", http.StatusTemporaryRedirect)
}

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func getGoogleUserInfo(code string) ([]byte, error) {
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("Failed to exchange code with Google: %s", err.Error())
	}

	resp, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("Failed to Get UserInfo %s\n", err.Error())
	}

	return io.ReadAll(resp.Body)

}

func googleLoginHandler(w http.ResponseWriter, req *http.Request) {
	state := generateStateOauthCookie(w)
	url := googleOauthConfig.AuthCodeURL(state)
	http.Redirect(w, req, url, http.StatusTemporaryRedirect)
}

func generateStateOauthCookie(w http.ResponseWriter) string {
	// cookieの有効期限設定
	expiration := time.Now().Add(1 * 24 * time.Hour)

	// 16byteのランダムな文字列を生成する。
	b := make([]byte, 16)
	rand.Read(b) // randomにバイトを埋め込む

	state := base64.URLEncoding.EncodeToString(b)

	cookie := &http.Cookie{
		Name:    "oauthstate",
		Value:   state,
		Expires: expiration,
	}
	http.SetCookie(w, cookie)
	return state
}
