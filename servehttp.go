package gh_pages_internal

import (
	"github.com/joshcarp/gh-pages-internal/pkg/pages"
	"net/http"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request){
	pages.ServeHTTP(w,r)
}
