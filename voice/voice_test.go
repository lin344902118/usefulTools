package voice

import (
	"os"
	"testing"
)

// 判断文件是否存在
func IsFileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func TestVoicePlayer_Speak(t *testing.T) {
	vp := NewVoicePlayer("")
	defer func() {
		vp.Release()
		RelaseOle()
	}()
	err := vp.Speak("hello world")
	if err != nil {
		t.Fail()
	}

}

func TestVoicePlayer_SpeakAndSave(t *testing.T) {
	fileName := "test.wav"
	vp := NewVoicePlayer(fileName)
	defer func() {
		vp.Release()
		RelaseOle()
	}()
	err := vp.SpeakAndSave("你好世界")
	if err != nil {
		t.Fail()
	}
	if !IsFileExists(fileName) {
		t.Fail()
	}
}

func TestReadFromDir(t *testing.T) {
	dir := "D:\\temp\\fitz579293743"
	err := ReadFromDir(dir)
	if err != nil {
		t.Fail()
	}
}
