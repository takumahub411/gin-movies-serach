package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/teris-io/shortid"
	"golang.org/x/crypto/bcrypt"
	"movies-search-go/initializers"
	"movies-search-go/models"
	"movies-search-go/utils"
	"net/http"
	"path/filepath"
	"time"
)

func SignupGetHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

func SignupPostHandler(c *gin.Context) {

	var user models.User

	user.Username = c.PostForm("name")
	user.Email = c.PostForm("email")
	user.Password = c.PostForm("password")
	PassConf := c.PostForm("passwordConf")

	errMg := ValidationUser(&user)
	if errMg != nil {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{
			"error": errMg,
		})
		return
	}

	//email domain check
	errMgDomain := utils.EmailVerifier(user.Email)
	if errMgDomain != "" {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{
			"error": errMgDomain,
		})
		return
	}

	// image receive
	Photo, err := c.FormFile("image")

	if err != nil {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{
			"error": "画像を設定してください。",
		})
	}
	filename := filepath.Base(Photo.Filename)
	user.Photo = filename
	//save images to imagesFolder
	if err := c.SaveUploadedFile(Photo, "./images/"+filename); err != nil {
		c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
		return
	}

	// short id create
	sid, _ := shortid.New(1, shortid.DefaultABC, 2342)
	userId, _ := sid.Generate()

	// email address exist
	initializers.DB.First(&user, "email = ?", user.Email)

	if user.ID != "" {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{
			"error": "メールアドレスが既に登録されています、他のメールアドレスを入力してください。",
		})
		return
	}

	// passwordConf

	if user.Password != PassConf {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{
			"error": "確認パスワードが間違っています",
		})
		return
	}

	//hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{
			"error": "パスワード入力をし直してください。",
		})
		return
	}

	user = models.User{
		ID:          userId,
		Username:    user.Username,
		Email:       user.Email,
		Password:    string(hash),
		Photo:       user.Photo,
		EmailVerify: false,

		CreatedAt: time.Now(),
	}

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{
			"error": "もう一度やり直してください。",
		})
		return
	}

	//send mail verify
	name := "登録認証"
	emailAddress := "greeceikitai@gmail.com"
	appPassword := "oqoxdcrowkznjzhq"
	sender := utils.NewGmailSend(name, emailAddress, appPassword)
	to := []string{user.Email}
	link := "localhost:8001/emailVerify/" + user.Username + "/" + user.ID

	fmt.Println(link)
	subject := "メール認証"
	content :=
		`<html>
			<h1>クリックして登録の認証を行ってください。</h1>
			<a href="` + link + `">こちらをクリック</a>
		</html>`

	err = sender.SendEmail(subject, content, to, nil, nil)

	c.Redirect(http.StatusFound, "/success_signup")

	//session, _ := initializers.Store.Get(c.Request, "session")
	//
	//session.Values["user"] = &user
	//session.Save(c.Request, c.Writer)
}

func LoginGetHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func SignupCompGetHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "success_signup.html", nil)
}

func LoginPostHandler(c *gin.Context) {
	var user models.User

	user.Email = c.PostForm("email")
	LoginPassword := c.PostForm("password")

	initializers.DB.First(&user, "email = ?", user.Email)

	//email verify check
	if user.EmailVerify == false {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": "メール認証が完了していません。登録に際にお送りしたメールのリンクをクリックして認証を行ってください。",
		})
		return
	}

	if user.ID == "" {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": "メールアドレスもしくはパスワードが間違っています。",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(LoginPassword))

	if err != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": "メールアドレスもしくはパスワードが間違っています。",
		})
		return
	}

	//generate jwt token
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	//	"sub": user.ID,
	//	"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	//})
	//
	//tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	//
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"error": "failed to create token",
	//	})
	//}
	//// send it back
	//c.SetSameSite(http.SameSiteLaxMode)
	//c.SetCookie("Auth", tokenString, 3600*24*30, "", "", false, true)
	//c.JSON(http.StatusOK, gin.H{})

	//session
	session, _ := initializers.Store.Get(c.Request, initializers.SessionName)
	userData, _ := session.Values["userId"]

	if userData != nil {
		session.Values["userId"] = ""
	}

	session.Values["userId"] = user.ID
	session.Values["username"] = user.Username

	session.Save(c.Request, c.Writer)

	if user.Role == 1 {
		c.Redirect(http.StatusMovedPermanently, "/manege")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/movies")
	}

}

func ValidationUser(u *models.User) (message []string) {

	//go validate check
	validate := validator.New() //create instance
	errMg := validate.Struct(u) //return ng

	var errorMessages []string

	if errMg != nil {
		for _, err := range errMg.(validator.ValidationErrors) {

			var errorMessage string
			inputName := err.Field() //バリデーションでNGになった変数名を取得

			switch inputName {
			case "Username":
				errorMessage = "名前は必須項目です。"
			case "Password":
				var typ = err.Tag()
				switch typ {
				case "required":
					errorMessage = "パスワードは必須項目です。"
				case "min=8,max=20":
					errorMessage = "パスワードは8文字以上20文字以内で設定してください"
				}
			case "Email":
				var typ = err.Tag()
				fmt.Println(typ)
				switch typ {
				case "required":
					errorMessage = "メールアドレスは必須項目です。"
				case "email":
					errorMessage = "メールアドレスの形式に誤りがあります"
				}
			}
			errorMessages = append(errorMessages, errorMessage)
		}
	}

	return errorMessages
}

func EmailVerifyGetHandler(c *gin.Context) {

	username := c.Param("username")
	userId := c.Param("id")

	result := GetUsername(username)
	if result == true {
		initializers.DB.Model(&models.User{}).Where("id = ?", userId).Update("EmailVerify", true)
	}
	c.HTML(http.StatusOK, "email_verifyComp.html", nil)

}

func GetUsername(name string) (uExist bool) {
	var user models.User
	uExist = true

	initializers.DB.First(&user, "username = ?", name)

	if user.Email == "" {
		return false
	}
	return uExist

}

func MypageGetHandler(c *gin.Context) {
	session, _ := initializers.Store.Get(c.Request, initializers.SessionName)
	id := session.Values["userId"]
	var user []*models.User

	initializers.DB.First(&user, "id = ?", id)

	c.HTML(http.StatusOK, "mypage.html", gin.H{
		"data": &user,
	})
}

func PasswordChangeGetHandler(c *gin.Context) {

	c.HTML(http.StatusOK, "passwordChange.html", nil)
}
func PasswordChangePostHandler(c *gin.Context) {
	var user models.User

	// get data from id
	session, _ := initializers.Store.Get(c.Request, initializers.SessionName)
	id := session.Values["userId"]
	//get post request
	old := c.PostForm("old_password")

	initializers.DB.First(&user, "id = ?", id)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(old))

	if err != nil {
		c.HTML(http.StatusBadRequest, "passwordChange.html", gin.H{
			"error": "パスワードが間違っています。",
		})
		return
	}

	newPass := c.PostForm("new_password")
	newPassConf := c.PostForm("new_passwordConf")
	fmt.Println(newPass)
	fmt.Println(newPassConf)
	if newPass != newPassConf {
		c.HTML(http.StatusBadRequest, "passwordChange.html", gin.H{
			"error": "新しいパスワードの入力を再度お願いします。",
		})
		return
	}
	//hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(newPass), 12)
	if err != nil {
		c.HTML(http.StatusBadRequest, "passwordChange.html", gin.H{
			"error": "パスワード入力をし直してください。",
		})
		return
	}
	user.Password = string(hash)
	initializers.DB.Save(&user)

	c.HTML(http.StatusFound, "login.html", gin.H{
		"message": "パスワードをリセットしました。ログインしなおしてください",
	})
	c.Redirect(http.StatusFound, "/login")

	//
	//initializers.DB.First(&user, "id = ?", id)
	//
	//c.HTML(http.StatusOK, "mypage.html", gin.H{
	//	"data": &user,
	//})
}

func EmailCheckGetHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "emailCheck.html", nil)
}
func SuccessPasswordResetPostHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "success_passwordReset.html", nil)
}
func EmailCheckPostHandler(c *gin.Context) {
	var user models.User

	//get post
	email := c.PostForm("email")

	// email address exist
	initializers.DB.First(&user, "email = ?", email)

	if user.ID == "" {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{
			"error": "入力されたメールアドレスは登録されていません。",
		})
		return
	}

	//send mail reset password
	//send mail verify
	name := "リセットパスワードURL"
	emailAddress := "greeceikitai@gmail.com"
	appPassword := "oqoxdcrowkznjzhq"
	sender := utils.NewGmailSend(name, emailAddress, appPassword)
	to := []string{user.Email}
	link := "localhost:8001/passwordReset/" + user.Username + "/" + user.ID
	fmt.Println()

	subject := "メール認証"
	content :=
		`<html>
			<h1>クリックしてリンクを開き再設定用のパスワードの登録を行ってください。</h1>
			<a href=` + link + `>こちらをクリック</a>
			<p>クリックできない場合はこちらをＷＥＢで検索してください。
			` + link + `
			</p>
		</html>`

	err := sender.SendEmail(subject, content, to, nil, nil)
	if err != nil {
		return
	}

	c.Redirect(http.StatusFound, "/success_sendReset")
}

func PasswordResetGetHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "passwordReset.html", nil)
}
func PasswordResetPostHandler(c *gin.Context) {

	username := c.Param("username")
	userId := c.Param("id")

	newPass := c.PostForm("password")
	newPassConf := c.PostForm("passwordConf")

	if newPass == "" || newPassConf == "" {
		c.HTML(http.StatusBadRequest, "passwordReset.html", gin.H{
			"error": "パスワードを入力してください",
		})
		return
	}

	if newPass != newPassConf {
		c.HTML(http.StatusBadRequest, "passwordReset.html", gin.H{
			"error": "確認用パスワードがあっていません",
		})
		return
	}
	result := GetUsername(username)

	if result == true {
		//hash password
		hash, err := bcrypt.GenerateFromPassword([]byte(newPass), 12)
		if err != nil {
			c.HTML(http.StatusBadRequest, "passwordReset.html", gin.H{
				"error": "パスワード入力をし直してください。",
			})
			return
		}
		initializers.DB.Model(&models.User{}).Where("id = ?", userId).Update("password", hash)
	}
	c.HTML(http.StatusOK, "email_verifyComp.html", nil)

}
