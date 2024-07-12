package utils

import (
	"errors"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedpassword), err
}

func ComparePassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func GenerateSlugs(name string) string {
	slug := strings.ReplaceAll(name, " ", "-")
	return strings.ToLower(slug)
}

type IDParams struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func GetIdFromParams(ctx *gin.Context) (uuid.UUID, error) {
	var params IDParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		return uuid.Nil, err
	}

	ID, err := uuid.Parse(params.ID)

	if err != nil {
		return ID, err
	}

	return ID, nil
}

func GetAPIKey(header *http.Header) (string, error) {
	val := header.Get("Authorization")

	if val == "" {
		return "", errors.New("no authentication found")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2 {
		return "", errors.New("invalid authorization")
	}

	if vals[0] != "Bearer" {
		return "", errors.New("invalid first part of authorization")
	}
	return vals[1], nil
}

func GenerateToken(userID uuid.UUID) (string, error) {
	var (
		key []byte
		t   *jwt.Token
		s   string
	)

	key = []byte("secret-key")

	t = jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss": "my-auth-server",
			"sub": "john",
			"foo": 2,
		})
	s, err := t.SignedString(key)

	if err != nil {
		return "", err
	}
	return s, nil
}

func GetIDInRoute(ctx *gin.Context, IDName string) (uuid.UUID, error) {
	stringId, ok := ctx.Params.Get(IDName)

	if !ok {
		return uuid.Nil, errors.New("no product id found")
	}

	id, err := uuid.Parse(stringId)

	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func IsValidUUID(uuid string) bool {
	// Define the regex pattern
	var uuidRegex = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[1-5][0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`)
	return uuidRegex.MatchString(uuid)
}
