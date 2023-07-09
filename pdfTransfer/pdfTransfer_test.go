package pdfTransfer

import (
	"fmt"
	"testing"
)

func TestPdfToTextFromLocal(t *testing.T) {
	filePath := "../test.pdf"
	dir, err := PdfToTextFromLocal(filePath)
	if err != nil {
		fmt.Println("pdf to text from local error", err.Error())
		t.Fail()
	}
	if dir == "" {
		fmt.Println("dir is empty")
		t.Fail()
	}
	fmt.Println(dir)
}

func TestPdfToTextFromNetwork(t *testing.T) {
	fileUrl := "http://192.168.2.121/%E6%96%B0%E7%89%88%E4%B8%80%E5%88%86%E9%92%9F%E7%BB%8F%E7%90%86%E4%BA%BA.pdf"
	dir, err := PdfToTextFromNetwork(fileUrl)
	if err != nil {
		fmt.Println("pdf to text from network error", err.Error())
		t.Fail()
	}
	if dir == "" {
		fmt.Println("dir is empty")
		t.Fail()
	}
	fmt.Println(dir)
}
