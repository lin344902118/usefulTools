package voice

import (
	"errors"
	"fmt"
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"os"
	"path/filepath"
	"strings"
	"usefulTools/pdfTransfer"
)

const (
	DEFAULT_VOICE_RATE   = -1
	DEFAULT_VOICE_VOLUME = 200
)

// 文字转语音播放器
type VoicePlayer struct {
	voice        *ole.IDispatch
	fileStream   *ole.IDispatch
	saveFileName string
	rate         int
	volume       int
}

// 初始化播放器
func NewVoicePlayer(fileName string) *VoicePlayer {
	unknown, _ := oleutil.CreateObject("SAPI.SpVoice")
	voice, _ := unknown.QueryInterface(ole.IID_IDispatch)
	var ff *ole.IDispatch
	if fileName != "" {
		saveFile, _ := oleutil.CreateObject("SAPI.SpFileStream")
		ff, _ = saveFile.QueryInterface(ole.IID_IDispatch)
	}
	return &VoicePlayer{
		voice:        voice,
		fileStream:   ff,
		saveFileName: fileName,
		rate:         DEFAULT_VOICE_RATE,
		volume:       DEFAULT_VOICE_VOLUME,
	}
}

// 设置速度
func (p *VoicePlayer) SetVoiceRate(rate int) {
	p.rate = rate
}

// 设置音量
func (p *VoicePlayer) SetVoiceVolume(volume int) {
	p.volume = volume
}

// 设置保存文件名称
func (p *VoicePlayer) SetSaveFileName(saveFileName string) {
	p.saveFileName = saveFileName
}

// 获取声音对象
func (p *VoicePlayer) GetVoice() *ole.IDispatch {
	return p.voice
}

// 获取文件流对象
func (p *VoicePlayer) GetFileStream() *ole.IDispatch {
	return p.fileStream
}

// 获取声音播放速率
func (p *VoicePlayer) GetVoiceRate() int {
	return p.rate
}

// 获取声音音量
func (p *VoicePlayer) GetVoiceVolume() int {
	return p.volume
}

// 获取保存文件名称
func (p *VoicePlayer) GetSaveFileName() string {
	return p.saveFileName
}

// 文字转语音
func (p *VoicePlayer) Speak(text string) error {
	if text == "" {
		return nil
	}
	if p.voice == nil {
		return errors.New("声音未初始化")
	}
	oleutil.PutProperty(p.voice, "Rate", p.rate)
	oleutil.PutProperty(p.voice, "Volume", p.volume)
	text = strings.ReplaceAll(text, " ", "")
	datas := strings.Split(text, ".")
	for _, data := range datas {
		oleutil.CallMethod(p.voice, "Speak", data)
	}
	oleutil.CallMethod(p.voice, "WaitUntilDone", 1000000)
	return nil
}

// 文字转语音并保存
func (p *VoicePlayer) SpeakAndSave(text string) error {
	if p.fileStream == nil {
		return errors.New("文件流未初始化")
	}
	oleutil.CallMethod(p.fileStream, "Open", p.saveFileName, 3, true)
	oleutil.PutPropertyRef(p.voice, "AudioOutputStream", p.fileStream)
	return p.Speak(text)
}

func (p *VoicePlayer) Release() {
	if p.voice != nil {
		p.voice.Release()
	}
	if p.fileStream != nil {
		oleutil.CallMethod(p.fileStream, "Close")
		p.fileStream.Release()
	}
}

func init() {
	ole.CoInitialize(0)
}

func RelaseOle() {
	ole.CoUninitialize()
}

func ReadFromDir(dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	vp := NewVoicePlayer("")
	defer func() {
		vp.Release()
		RelaseOle()
	}()
	for _, f := range files {
		if !pdfTransfer.IsFileWithSuffix(f.Name(), ".txt") {
			fmt.Println("file format error")
			continue
		}
		fileName := strings.TrimSuffix(f.Name(), ".txt")
		err := vp.Speak("正在阅读" + fileName)
		if err != nil {
			fmt.Println("speak error", err.Error())
			continue
		}
		filePath := filepath.Join(dir, f.Name())
		datas, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("read file error", err.Error())
			continue
		}
		err = vp.Speak(string(datas))
		if err != nil {
			fmt.Println("speak error", err.Error())
			continue
		}
		err = vp.Speak(fileName + "阅读完毕，阅读下一页")
		if err != nil {
			fmt.Println("speak error", err.Error())
			continue
		}
	}
	return nil
}
