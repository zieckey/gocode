package tmail
import (
"testing"
)

func TestFindAll(t *testing.T) {
	SendMail("hello, 关于mail问题", "Hi, 这个问题我同意。", "weizili@360.cn", "zieckey@163.com")
}
