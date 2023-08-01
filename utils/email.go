package utils

import (
	"fmt"
	emailverifier "github.com/AfterShip/email-verifier"
)

type EmailData struct {
	URL     string
	Name    string
	content string
}

var (
	verifier = emailverifier.NewVerifier()
)

func EmailVerifier(email string) string {

	errMg := ""

	ret, err := verifier.Verify(email)
	if err != nil {
		errMg = "不明なメールアドレスです。"
		return errMg
	}
	if !ret.Syntax.Valid {
		fmt.Println("不明なメールアドレスです。")
		errMg = "不明なメールアドレスです。"
		return errMg
	}

	//fmt.Println("email validation result", ret)

	// check domain suggestion

	if ret.Suggestion != "" {
		errMg = "ドメインに誤りがあります。" + ret.Suggestion + "ではありませんか？"
		return errMg
	}
	//domain check
	verifier = verifier.EnableDomainSuggest()

	return errMg
}
