// Suluvir streaming server
// Copyright (C) 2017  Jannis Fink
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, version 3.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package intnl

import (
	"encoding/json"
	"github.com/suluvir/server/auth"
	"github.com/suluvir/server/logging"
	"github.com/suluvir/server/schema"
	"github.com/suluvir/server/web/handler/api"
	"github.com/suluvir/server/web/httpHelpers"
	"github.com/uber-go/zap"
	"net/http"
)

type createUser struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	PasswordRepeat string `json:"password_repeat"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var payload createUser
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		logging.GetLogger().Error("error dezerializing request body", zap.Error(err))
		api.SendJsonError(w, http.StatusBadRequest, "error dezerializing request body")
		return
	}

	if payload.Password != payload.PasswordRepeat {
		api.SendJsonError(w, http.StatusBadRequest, "passwords do not match")
		return
	}

	user := auth.CreateUser(payload.Username, payload.Email, payload.Password)
	schema.GetDatabase().Create(&user)

	httpHelpers.ServeJsonWithoutCache(w, &user)
}