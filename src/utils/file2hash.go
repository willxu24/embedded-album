package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
)

func File2md5(file io.Reader) string {
	md5 := md5.New()
	io.Copy(md5, file)
	return hex.EncodeToString(md5.Sum(nil))
}

func File2Sha256(file io.Reader) string {
	sha256 := sha256.New()
	io.Copy(sha256, file)
	return hex.EncodeToString(sha256.Sum(nil))
}
