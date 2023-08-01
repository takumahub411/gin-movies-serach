package utils

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSendEmail(t *testing.T) {
	name := "emailTest"
	emailAddress := "送る側のアドレス"
	appPassword := "oqoxdcrowkznjzhq"
	sender := NewGmailSend(name, emailAddress, appPassword)

	subject := "this is a test"
	content := `
	<h1>hello !</h1>
	<p>test send message</p>
	`
	to := []string{"メールアドレス"}

	err := sender.SendEmail(subject, content, to, nil, nil)
	require.NoError(t, err)
}
