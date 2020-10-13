package parser

import (
	"github.com/jormin/go-study/crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	b, err := ioutil.ReadFile("profile.html")
	if err != nil {
		panic(err)
	}
	result := ParseProfile(string(b))
	basicInfo := result.Items[0].Data.(model.BasicInfo)
	expectID := 1865081142
	expectNickname := "念初"
	if expectID != basicInfo.ID {
		t.Errorf("expect ID: %d; got %d", expectID, basicInfo.ID)
	}
	if expectNickname != basicInfo.Nickname {
		t.Errorf("expect nickname: %s; got %s", expectNickname, basicInfo.Nickname)
	}
}
