package tmail
import (
"testing"
)

func TestFindAll(t *testing.T) {
	SendMail("hello, 关于mail问题", "Hi, 这个问题我同意。", "weizili@360.cn", "zieckey@163.com") //TODO 收到的邮件有乱码
	SendMail("about the student", "I agree.", "weizili@360.cn", "zieckey@163.com")
}
