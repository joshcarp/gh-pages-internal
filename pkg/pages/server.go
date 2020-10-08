package pages

import (
	"github.com/joshcarp/gop"
	"github.com/joshcarp/gop/gop/cli"
	"github.com/joshcarp/gop/gop/gop_filesystem"
	"github.com/joshcarp/gop/gop/gop_gcs"
	"github.com/joshcarp/gop/gop/retriever/retriever_github"
	"github.com/joshcarp/gop/gop/retriever/retriever_naked"
	"github.com/spf13/afero"
	"os"
)

type Server struct{
	*gop.GopperService
}

func New()Server{
	s, _ := NewGopper(os.Getenv("GCS_BUCKET"), os.Getenv("FS_TYPE"), "GIT_TOKENS")
	return Server{GopperService:s}
}



/* NewGopper returns a GopperService for a config; This Gopper can use an os filesystem, memory filesystem or a gcs bucket*/
func NewGopper(cachelocation, fsType, tokenEnv string) (*gop.GopperService, error) {
	r := gop.GopperService{}
	switch fsType {
	case "os":
		r.Gopper = gop_filesystem.New(afero.NewOsFs(), "")
	case "gcs":
		gcs := gop_gcs.New(cachelocation)
		r.Gopper = &gcs
	}
	token, _ := cli.NewTokenMap(tokenEnv, "")
	r.Retriever = retriever_naked.New(retriever_github.New(token), "HEAD")
	return &r, nil
}
