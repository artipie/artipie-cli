package artictl

import (
	"context"
	"net/http"

	"github.com/artipie/artipie-cli/internal/config"
	"github.com/artipie/artipie-cli/pkg/artipie"
)

const userAgent = "ArtiCtl"

// ArtiCtl - main conext.
type ArtiCtl struct {
	ctx    context.Context
	Client *artipie.Client
}

// NewArtiCtl context.
func NewArtiCtl(ctx context.Context, ctlCtx *config.CtlContext) (*ArtiCtl, error) {
	client, err := artipie.NewClient(&http.Client{}, ctlCtx.Endpoint, ctlCtx.Auth)
	if err != nil {
		return nil, err
	}
	client.UserAgent = userAgent
	return &ArtiCtl{ctx, client}, nil
}
