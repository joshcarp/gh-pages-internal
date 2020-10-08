package pages

import (
	"github.com/joshcarp/gop"
	"os"
)

type Server struct{
	*gop.GopperService
}

func New()Server{
	s, _ := gop.NewGopper(os.Getenv("GCS_BUCKET"), os.Getenv("FS_TYPE"), "GIT_TOKENS")
	return Server{GopperService:s}
}

