package util

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math/rand"
	"path/filepath"
	"strings"
	"time"
)

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var videoExts = [...]string{".mp4", ".mov"}
var audioExts = [...]string{".mp3"}
var imageExts = [...]string{".png", ".jpeg", ".jpg", ".gif"}

func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func Sha1(data []byte) string {
	_sha1 := sha1.New()
	_sha1.Write(data)
	return hex.EncodeToString(_sha1.Sum([]byte("")))
}

// 生成width长度验证码
func VerifyCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		_, _ = fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

func CheckVideoExt(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	for _, item := range videoExts {
		if ext == item {
			return true
		}
	}
	return false
}

func CheckAppUploadFileExt(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	for _, item := range videoExts {
		if ext == item {
			return true
		}
	}
	for _, item := range audioExts {
		if ext == item {
			return true
		}
	}
	for _, item := range imageExts {
		if ext == item {
			return true
		}
	}
	return false
}

func JudgeAppUploadFileType(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	for _, item := range videoExts {
		if ext == item {
			return "video"
		}
	}
	for _, item := range audioExts {
		if ext == item {
			return "audio"
		}
	}
	return "image"
}

func CheckImage(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	for _, item := range imageExts {
		if ext == item {
			return true
		}
	}
	return false
}

// time.Now().Format("2006-01-02 15:04:05")
// "2020-04-09T18:16:07+08:00" -> format
// format := "2006-01-02 15:04"
func TimeFormat(t time.Time) string {
	format := "2006-01-02 15:04"
	return t.Format(format)
}

// 获取分页
func GetLimit(count, skip, pageSize int) int {
	limit := 0
	left := count - skip
	if left >= pageSize {
		limit = pageSize
	} else {
		if left <= 0 {
			return 0
		}
		return left
	}
	return limit
}
