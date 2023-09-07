package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/pro-assistance/pro-assister/tokenHelper"
)

type Claim string
type Claims []Claim

const (
	ClaimUserID    Claim = "user_id"
	ClaimDomainIDS Claim = "domain_ids"
)

func (item Claim) Inject(r *http.Request, h *tokenHelper.TokenHelper) error {
	d, err := h.ExtractTokenMetadata(r, item.String())
	if err != nil {
		return err
	}
	*r = *r.WithContext(context.WithValue(r.Context(), item, d))
	//ctx = context.WithValue(ctx, claim, d)
	return nil
}

func (items Claims) Inject(r *http.Request, h *tokenHelper.TokenHelper) (err error) {
	for _, claim := range items {
		err = claim.Inject(r, h)
		if err != nil {
			break
		}
	}
	return err
}

func (item Claim) String() string {
	return string(item)
}

func (item Claim) Split() []string {
	return strings.Split(item.String(), ",")
}

func (item Claim) FromContext(ctx context.Context) string {
	return ctx.Value(item.String()).(string)
}

func (item Claim) FromContextSlice(ctx context.Context) []string {
	return strings.Split(ctx.Value(item.String()).(string), ",")
}
