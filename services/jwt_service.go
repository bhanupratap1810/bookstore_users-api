package services

import (
	"github.com/kataras/jwt"
	"os"
	"time"
)

type JWTService interface {
	GenerateToken(userId int64, role string) []byte
	ValidateToken(tokenString string) error
}

type jwtService struct {
	secretKey []byte
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
	}
}

func getSecretKey() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret"
	}
	sec:=[]byte(secret)
	return sec
}

func (j *jwtService) GenerateToken(userId int64, role string) []byte {
	myClaims := map[string]interface{}{
		"user_id": userId,
		"user_role": role,
	}
	token, _ := jwt.Sign(jwt.HS256, j.secretKey, myClaims, jwt.MaxAge(15 * time.Minute))
	return token
}

func (j *jwtService) ValidateToken(tokenString string) error {
	token:=[]byte(tokenString)
	verifiedToken, err := jwt.Verify(jwt.HS256, j.secretKey, token)
	if err!=nil {
		return err
	}

	var claims map[string]interface{}
	err = verifiedToken.Claims(&claims)
	if err!=nil {
		return err
	}
	return nil
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

