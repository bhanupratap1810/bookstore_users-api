package services

import (
	"encoding/json"
	"fmt"
	"github.com/bhanupratap1810/bookstore_users-api/constants"
	"github.com/kataras/jwt"
	"os"
	"time"
)

type JWTService interface {
	GenerateToken(userId int64, role string) []byte
	ValidateToken(tokenString string) (*LoginJwtClaims, error)
}

type jwtService struct {
	secretKey []byte
}

type LoginJwtClaims struct {
	UserId int64  `json:"user_id"`
	Role   string `json:"role"`
	jwt.Claims
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
	}
}

func getSecretKey() []byte {
	secret := os.Getenv(constants.JwtSecret)
	if secret == "" {
		secret = constants.DefaultJwtSecret
	}
	sec := []byte(secret)
	return sec
}

func (j *jwtService) GenerateToken(userId int64, role string) []byte {
	loginJwtClaims := LoginJwtClaims{
		UserId: userId,
		Role:   role,
		Claims: jwt.Claims{
			Expiry: int64(15 * time.Minute),
			Issuer: "Library Service",
		},
	}
	token, err := jwt.Sign(jwt.HS256, j.secretKey, loginJwtClaims)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return token
}

//todo return of below func
func (j *jwtService) ValidateToken(tokenString string) (*LoginJwtClaims, error) {
	token := []byte(tokenString)
	verifiedToken, err := jwt.Verify(jwt.HS256, j.secretKey, token)
	if err != nil {
		return nil, err
	}

	var claims map[string]interface{}
	err = verifiedToken.Claims(&claims)
	//todo learn how to use this
	//expiry := verifiedToken.StandardClaims.Expiry
	if err != nil {
		return nil, err
	}
	jwtClaims := &LoginJwtClaims{}
	if val, ok := claims[constants.UserIdKey]; ok {
		jwtClaims.UserId, _ = val.(json.Number).Int64()
	}
	if val, ok := claims[constants.RoleKey]; ok {
		jwtClaims.Role = val.(string)
	}
	return jwtClaims, nil
}

// Keep it secret.
//var sharedKey = []byte("sercrethatmaycontainch@r$32chars")
//
//func main() {
//	// Generate a token:
//	myClaims := map[string]interface{}{
//		"foo": "bar",
//	}
//	token, _:= jwt.Sign(jwt.HS256, sharedKey, myClaims, jwt.MaxAge(15 * time.Minute))
//
//	// Verify and extract claims from a token:
//	verifiedToken, _ := jwt.Verify(jwt.HS256, sharedKey, token)
//
//	var claims map[string]interface{}
//	_ = verifiedToken.Claims(&claims)
//}
//type JWTService interface {
//	GenerateToken(name string, admin bool) string
//	ValidateToken(tokenString string) (*jwt.Token, error)
//}
//
//// jwtCustomClaims are custom claims extending default ones.
//type jwtCustomClaims struct {
//	Name  string `json:"name"`
//	Admin bool   `json:"admin"`
//	jwt.StandardClaims
//}
//
//type jwtService struct {
//	secretKey string
//	issuer    string
//}
//
//func NewJWTService() JWTService {
//	return &jwtService{
//		secretKey: getSecretKey(),
//		issuer:    "pragmaticreviews.com",
//	}
//}
//
//func getSecretKey() string {
//	secret := os.Getenv("JWT_SECRET")
//	if secret == "" {
//		secret = "secret"
//	}
//	return secret
//}
//
//func (jwtSrv *jwtService) GenerateToken(username string, admin bool) string {
//
//	// Set custom and standard claims
//	claims := &jwtCustomClaims{
//		username,
//		admin,
//		jwt.StandardClaims{
//			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
//			Issuer:    jwtSrv.issuer,
//			IssuedAt:  time.Now().Unix(),
//		},
//	}
//
//	// Create token with claims
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//
//	// Generate encoded token using the secret signing key
//	t, err := token.SignedString([]byte(jwtSrv.secretKey))
//	if err != nil {
//		panic(err)
//	}
//	return t
//}
//
//func (jwtSrv *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
//	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		// Signing method validation
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
//		}
//		// Return the secret signing key
//		return []byte(jwtSrv.secretKey), nil
//	})
//}

//func NewLoginController(loginService service.LoginService,
//	jWtService service.JWTService) LoginController {
//	return &loginController{
//		loginService: loginService,
//		jWtService:   jWtService,
//	}
//}

//type loginController struct {
//	loginService service.LoginService
//	jWtService   service.JWTService
//}
