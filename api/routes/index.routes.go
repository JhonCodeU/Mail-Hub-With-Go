package routes

import (
	"github/jhoncodeu/mailbox-masive-go/config"
	"github/jhoncodeu/mailbox-masive-go/src/auth"
	"net/http"
)

var headers = auth.HeaderHttpBasicAuth(config.AuthUser, config.AuthPass)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Api Mailbox Masive v1"))
}
