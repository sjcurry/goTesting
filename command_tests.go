package main

import(
	"testing"
	cmds "github.com/sjcurry/packages/commands"
}

func TestCcmdLeft(t *testing.T) {
	res, err := cmds.C_cmd([]string{"1", "3", "b"}, "aaaaaa")
	if err == nil {
		if res != "bbbaaa" {
			t.Errorf("TestCcmdLeft Expected 'bbbaaa';got %s", res)
		}
	}
}

