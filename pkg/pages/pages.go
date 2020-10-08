package pages

import (
	"mime"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/joshcarp/gop"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request){
	New().ServeHTTP(w, r)
}


func (s Server)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		gop.HandleErr(w, err)
	}()
	pth := strings.TrimLeft(r.URL.Path, "/")
	pth = strings.TrimLeft(pth, os.Getenv("BASE_PATH"))
	pth = strings.TrimRight(pth, "/")
	pth = path.Join(os.Getenv("BASE"), pth)
	if r.Method == "DELETE"{
		err = s.Cache(pth+"/", nil)
		return
	}
	switch path.Ext(pth){
	case "":
		pth += "/index.html"
	}
	w.Header().Add("Content-Type", mime.TypeByExtension(path.Ext(pth)))
	res, _, err := s.Retrieve(pth+"@HEAD")
	if err != nil {
		return
	}
	_, err = w.Write(res)
}
