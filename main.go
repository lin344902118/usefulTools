package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"usefulTools/pdfTransfer"
	"usefulTools/voice"
)

func pdfFromLocal(pdfFile string, deleteAfter bool) {
	dir, err := pdfTransfer.PdfToTextFromLocal(pdfFile)
	if err != nil {
		fmt.Println("pdf to text from local failed", err.Error())
		return
	}
	fmt.Println("tmp dir", dir)
	err = voice.ReadFromDir(dir)
	if err != nil {
		fmt.Println("read from dir failed", err.Error())
		return
	}
	if deleteAfter {
		err = os.RemoveAll(dir)
		if err != nil {
			fmt.Println("remove all failed", err.Error())
			return
		}
	}
}

func pdfFromNetwork(pdfUrl string, deleteAfter bool) {
	dir, err := pdfTransfer.PdfToTextFromNetwork(pdfUrl)
	if err != nil {
		fmt.Println("pdf to text from network failed", err.Error())
		return
	}
	fmt.Println("tmp dir", dir)
	err = voice.ReadFromDir(dir)
	if err != nil {
		fmt.Println("read from dir failed", err.Error())
		return
	}
	if deleteAfter {
		err = os.RemoveAll(dir)
		if err != nil {
			fmt.Println("remove all failed", err.Error())
			return
		}
	}
}

func main() {
	var pdfLocation string
	flag.StringVar(&pdfLocation, "location", "", "pdf地址")
	var deleteAfter bool
	flag.BoolVar(&deleteAfter, "delete", true, "生成后是否删除")
	flag.Parse()
	if pdfLocation != "" {
		fmt.Println("pdf location", pdfLocation)
		if strings.HasPrefix(pdfLocation, "http") {
			pdfFromNetwork(pdfLocation, deleteAfter)
		} else {
			pdfFromLocal(pdfLocation, deleteAfter)
		}
	}
}
