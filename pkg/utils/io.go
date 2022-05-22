package utils

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/kanztu/goblog/pkg/server_context"
)

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func ReadFileToString(fname string) string {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		server_context.SrvCtx.Logger.Error(err)
	}

	return string(b)
}

func ReadFileToByte(fname string) ([]byte, error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	return b, nil
}
