package ojichatslack

import (
	"os"
	"encoding/json"
	"net/http"

	"github.com/nlopes/slack"
	"github.com/greymd/ojichat/generator"
)

var verificationToken = os.Getenv("VERIFICATION_TOKEN")

func generateOjiMsg(name string) (string, error) {
	config := generator.Config{TargetName: name, EmojiNum: 4, PunctiuationLebel: 0}
	result, err := generator.Start(config)
	if err != nil {
		return "", err
	}
	return result, nil
}

func Ojichat(w http.ResponseWriter, r *http.Request) {
	s, err := slack.SlashCommandParse(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !s.ValidateToken(verificationToken) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	switch s.Command {
	case "/ojichat":
		msg, err := generateOjiMsg(s.Text)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		m := &slack.Msg{Text: msg, ResponseType:"in_channel"}
		b, err := json.Marshal(m)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

