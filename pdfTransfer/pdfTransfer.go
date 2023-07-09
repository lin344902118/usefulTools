package pdfTransfer

import (
	"errors"
	"fmt"
	"github.com/gen2brain/go-fitz"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// 先读取文件，分为网络读取和本地读取。然后转换为text

// 判断文件是否是pdf
func IsPdfFile(filePath string) bool {
	return IsFileWithSuffix(filePath, ".pdf")
}

// 判断文件后缀是否正确
func IsFileWithSuffix(filePath, suffix string) bool {
	fileExt := filepath.Ext(filepath.Base(filePath))
	if fileExt != suffix {
		return false
	}
	return true
}

// pdf转txt
func PdfToText(doc *fitz.Document) (string, error) {
	// 创建临时文件将转换的text存放到临时文件中
	tmpDir, err := os.MkdirTemp(os.TempDir(), "fitz")
	if err != nil {
		return "", err
	}
	// 将文件转换为text
	for n := 0; n < doc.NumPage(); n++ {
		text, err := doc.Text(n)
		if err != nil {
			return tmpDir, err
		}

		f, err := os.Create(filepath.Join(tmpDir, fmt.Sprintf("%03d.txt", n)))
		if err != nil {
			return tmpDir, err
		}
		_, err = f.WriteString(text)
		if err != nil {
			f.Close()
			return tmpDir, err
		}
		f.Close()
	}
	return tmpDir, nil
}

// 本地文件转换
func PdfToTextFromLocal(filePath string) (string, error) {
	if !IsPdfFile(filePath) {
		return "", errors.New("File format error: Only support pdf.")
	}
	doc, err := fitz.New(filePath)
	if err != nil {
		return "", err
	}
	defer doc.Close()
	return PdfToText(doc)
}

// 网络文件转换
func PdfToTextFromNetwork(fileUrl string) (string, error) {
	if !strings.HasPrefix(fileUrl, "http") {
		return "", errors.New("url format error:Should start with http or https")
	}
	resp, err := http.Get(fileUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	datas, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	doc, err := fitz.NewFromMemory(datas)
	if err != nil {
		return "", err
	}
	defer doc.Close()
	return PdfToText(doc)
}
