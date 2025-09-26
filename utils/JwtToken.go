package utils

import (
	"github.com/Alifarid0011/questionnaire-back-end/constant"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JwtToken interface {
	GenerateAccessToken(Expiry int64, uid primitive.ObjectID, AttrMap AttributeMap) (string, error)
	GenerateRefreshToken(Expiry int64, uid primitive.ObjectID) (string, error)
	ParseToken(tokenStr string) (*CustomClaims, error)
}

// AttributeMap key = object:["act":"attr"], value = attribute (e.g., uid, *)
type AttributeMap = map[string]map[string]string
type CustomClaims struct {
	UID       string             `json:"uid"`
	TokenType constant.TokenType `json:"type"`
	AttrMap   AttributeMap       `json:"attr_map"` // key = object["act":"attr"], value = attribute (e.g., uid, *)
	jwt.RegisteredClaims
}

func (c *CustomClaims) ParseAttr(obj, action string) (string, bool) {
	if value, exist := c.AttrMap[obj][action]; exist {
		return value, exist
	} else if value, exist = c.AttrMap[obj]["*"]; exist {
		return value, exist
	}
	return "*", false
}
