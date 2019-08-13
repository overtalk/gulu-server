package gitlab_test

import (
	"os"
	"testing"

	"gitlab.com/SausageShoot/admin-server/utils/gitlab"
	"gitlab.com/SausageShoot/admin-server/utils/parse"
)

func TestCatcher_Fetch(t *testing.T) {
	token, isExist := os.LookupEnv("GITLAB_TOKEN")
	if !isExist {
		t.Error("`GITLAB_TOKEN` is absent")
		return
	}

	pid, isExist := os.LookupEnv("GITLAB_PID")
	if !isExist {
		t.Error("`GITLAB_PID` is absent")
		return
	}

	catcher := gitlab.Catcher(gitlab.Config{
		Token: token,
		Ref:   "master",
		Pid:   int(parse.Int(pid)),
		Url:   "https://git.xindong.com/api/v4",
	})

	data, err := catcher.Fetch("WeaponList.txt")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(string(data))
}
