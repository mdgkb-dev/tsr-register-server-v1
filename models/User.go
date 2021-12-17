package models

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v7"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	bun.BaseModel `bun:"users,alias:users"`
	ID                     uuid.UUID                `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Login                  string                   `json:"login"`
	Password               string                   `json:"password"`
	RegisterPropertyToUser RegisterPropertiesToUser `bun:"rel:has-many" json:"registerPropertyToUser"`
	//
	RegistersUsers RegistersUsers `bun:"rel:has-many" json:"registersUsers"`
	RegistersUsersForDelete []uuid.UUID `bun:"-" json:"registersUsersForDelete"`
}

type Users []*User


func (item *User) SetIdForChildren() {
	for i := range item.RegistersUsers {
		item.RegistersUsers[i].UserID = item.ID
	}
}


type RegisterPropertyToUser struct {
	bun.BaseModel      `bun:"register_property_to_user,alias:register_property_to_user"`
	ID                 uuid.UUID         `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	RegisterProperty   *RegisterProperty `bun:"rel:belongs-to" json:"registerProperty"`
	RegisterPropertyID uuid.UUID         `bun:"type:uuid" json:"registerPropertyId"`
	User               *User             `bun:"rel:belongs-to" json:"user"`
	UserID             uuid.UUID         `bun:"type:uuid" json:"userId"`
}

type RegisterPropertiesToUser []*RegisterPropertyToUser

func (i *User) GenerateHashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(i.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	i.Password = string(hash)
	return nil
}

func (i *User) CompareWithHashPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(i.Password), []byte(password)) == nil
}

func (u *User) CreateToken() (*TokenDetails, error) {
	return CreateToken(u.ID.String())
}

func GetUserID(c *gin.Context) (*uuid.UUID, error) {
	metaData, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		return nil, err
	}
	userId, err := uuid.Parse(metaData.UserId)
	if err != nil {
		return nil, err
	}
	return &userId, err
}

func CreateToken(userid string) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.AccessUuid = uuid.NewString()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = td.AccessUuid + "++" + userid

	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	//Creating Refresh Token
	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf") //this should be in an env file
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return td, nil
}

func CreateAuth(userid string, td *TokenDetails, client *redis.Client) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := client.Set(td.AccessUuid, userid, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := client.Set(td.RefreshUuid, userid, rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

func DeleteAuth(givenUuid string, client *redis.Client) (int64, error) {
	deleted, err := client.Del(givenUuid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}

func DeleteTokens(authD *AccessDetails, client *redis.Client) error {
	//get the refresh uuid
	refreshUuid := fmt.Sprintf("%s++%s", authD.AccessUuid, authD.UserId)
	//delete access token
	deletedAt, err := client.Del(authD.AccessUuid).Result()
	if err != nil {
		return err
	}
	//delete refresh token
	deletedRt, err := client.Del(refreshUuid).Result()
	if err != nil {
		return err
	}
	//When the record is deleted, the return value is 1
	if deletedAt != 1 || deletedRt != 1 {
		return errors.New("something went wrong")
	}
	return nil
}

func ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		userId, ok := claims["user_id"].(string)
		if !ok {
			return nil, err
		}
		return &AccessDetails{
			AccessUuid: accessUuid,
			UserId:     userId,
		}, nil
	}
	return nil, err
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("jdnfksdmfksd"), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("token")
	//strArr := strings.Split(bearToken, " ")
	//if len(strArr) == 2 {
	//	return strArr[1]
	//}
	return bearToken
}
