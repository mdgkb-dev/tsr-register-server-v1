package models

import (
	"context"
	"net/http"
	"strings"

	"github.com/pro-assistance/pro-assister/tokenHelper"
	"github.com/uptrace/bun"

	"github.com/google/uuid"
)

type User struct {
	bun.BaseModel `bun:"users,alias:users"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Login         string        `json:"login"`

	RegisterPropertyToUser  RegisterPropertiesToUser `bun:"rel:has-many" json:"registerPropertyToUser"`
	RegistersUsers          RegistersUsers           `bun:"rel:has-many" json:"registersUsers"`
	RegistersUsersForDelete []uuid.UUID              `bun:"-" json:"registersUsersForDelete"`

	UsersDomains UsersDomains `bun:"rel:has-many" json:"-"`

	UserAccountID uuid.NullUUID `bun:"type:uuid" json:"userAccountId"`
	UserAccount   *UserAccount  `bun:"rel:belongs-to" json:"userAccount"`
}

type Users []*User

type UsersWithCount struct {
	Users Users `json:"users"`
	Count int   `json:"count"`
}

func (item *User) SetIDForChildren() {
	for i := range item.RegistersUsers {
		item.RegistersUsers[i].UserID = item.ID.UUID
	}
}

type RegisterPropertyToUser struct {
	bun.BaseModel      `bun:"register_property_to_user,alias:register_property_to_user"`
	ID                 uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	RegisterProperty   *Domain   `bun:"rel:belongs-to" json:"registerProperty"`
	RegisterPropertyID uuid.UUID `bun:"type:uuid" json:"registerPropertyId"`
	User               *User     `bun:"rel:belongs-to" json:"user"`
	UserID             uuid.UUID `bun:"type:uuid" json:"userId"`
}

type RegisterPropertiesToUser []*RegisterPropertyToUser

type Claims string

func (c Claims) String() string {
	return string(c)
}

func (c Claims) Split() []string {
	return strings.Split(c.String(), ",")
}

func (c Claims) FromContext(ctx context.Context) string {
	return ctx.Value(c.String()).(string)
}

func (c Claims) FromContextSlice(ctx context.Context) []string {
	return strings.Split(ctx.Value(c.String()).(string), ",")
}

const (
	ClaimUserID    Claims = "user_id"
	ClaimDomainIDS Claims = "domain_ids"
)

func (item *User) SetJWTClaimsMap(claims map[string]interface{}) {
	claims[ClaimUserID.String()] = item.ID.UUID
	domainIds := make([]string, len(item.UsersDomains))
	for i := range item.UsersDomains {
		domainIds[i] = item.UsersDomains[i].DomainID.UUID.String()
	}
	claims[ClaimDomainIDS.String()] = strings.Join(domainIds, ",")
}

func (item User) InjectClaims(r *http.Request, h *tokenHelper.TokenHelper) (ctx context.Context, err error) {
	ctx = context.Background()
	for _, claim := range []Claims{ClaimUserID, ClaimDomainIDS} {
		ctx, err = item.InjectClaim(r, h, claim, ctx)
		if err != nil {
			break
		}
	}
	return ctx, err
}

func (item User) InjectClaim(r *http.Request, h *tokenHelper.TokenHelper, claim Claims, ctx context.Context) (context.Context, error) {
	d, err := h.ExtractTokenMetadata(r, claim.String())
	if err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, claim.String(), d)
	return ctx, nil
}
