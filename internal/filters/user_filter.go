package filters

import "net/http"

type UserFilter struct {
	Search string
}

func (u *UserFilter) Mapper(r *http.Request) {
	u.Search = r.URL.Query().Get("search")
}
