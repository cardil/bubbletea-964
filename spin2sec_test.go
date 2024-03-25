package bubbleteaepolldevnullinput_test

import (
	"bytes"
	"os"
	"os/exec"
	"path"
	"runtime"
	"testing"
)

func TestSpin2SecFailing(t *testing.T) {
	c := spin2secCmd()
	expectToNotFail(t, c)
}

func TestSpin2SecPassing(t *testing.T) {
	c := spin2secCmd()
	buf := bytes.NewBufferString("")
	c.Stdin = buf
	expectToNotFail(t, c)
}

func spin2secCmd() *exec.Cmd {
	c := exec.Command("go", "run", "./cmd/spin2sec")
	c.Dir = dir()
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c
}

func dir() string {
	_, file, _, _ := runtime.Caller(1)
	return path.Dir(file)
}

func expectToNotFail(t *testing.T, c *exec.Cmd) {
	t.Helper()
	if err := c.Run(); err != nil {
		t.Fatal("unexpected error:", err)
	}
}
