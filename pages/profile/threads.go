package profile

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// GetThreadsByUser shows all forum threads of a particular user.
func GetThreadsByUser(ctx *aero.Context) string {
	nick := ctx.Get("nick")
	viewUser, err := arn.GetUserByNick(nick)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "User not found", err)
	}

	threads := viewUser.Threads()
	arn.SortThreadsLatestFirst(threads)

	return ctx.HTML(components.ProfileThreads(threads, viewUser, utils.GetUser(ctx)))
}
