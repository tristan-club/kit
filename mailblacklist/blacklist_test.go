package mailblacklist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var blacklist = []string{
	"stimail.cn", "multiscanner.org", "colingpt.pro", "thinkmail.cc", "mail.tmwlsw.com", "tankmail.cn", "veraa.com", "thinkmail.top", "0dg.top", "tipent.com",
}

func initTestEnv() {
	if err := InitEmailBlackListChecker("0.0.0.0:6379", ""); err != nil {
		panic("init blacklist checker error")
	}
}

func TestAddBlacklist(t *testing.T) {
	initTestEnv()
	err := AddEmailBlacklist(blacklist)
	assert.NoError(t, err, "add blacklist error")
}

func TestCheckBlackList(t *testing.T) {
	initTestEnv()
	shouldValidList := []string{"frdco@58email.vip", "0.0.test@gmail.com", "tyl.0309@qq.com", "test@qq.com", "test@outlook.com", "test@hotmail.com", "test@163.com"}
	shouldInvalidList := []string{"test@stimail.cn", "test@multiscanner.org", "test@colingpt.pro", "test@thinkmail.cc", "test@mail.tmwlsw.com", "test@tankmail.cn", "test@veraa.com", "test@thinkmail.top", "test@0dg.top", "test@tipent.com"}

	for _, v := range shouldValidList {
		flag, err := CheckEmailInBlackList(v)
		assert.NoError(t, err, "check email error")
		assert.Equal(t, false, flag, "assert %s valid error", v)
	}

	for _, v := range shouldInvalidList {
		flag, err := CheckEmailInBlackList(v)
		assert.NoError(t, err, "check email error")
		assert.Equal(t, true, flag, "assert %s invalid error", v)
	}

}
