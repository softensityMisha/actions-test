package main

import (
	"log"
	"os"
	"strings"

	"github.com/h2non/bimg"
)

const (
	pdfPostfix  = ".pdf"
	jpegPostfix = ".jpeg"
)

func main() {
	content, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}

	for _, entrieContent := range content {
		if !strings.Contains(entrieContent.Name(), pdfPostfix) {
			continue
		}

		outputName := strings.ReplaceAll(entrieContent.Name(), pdfPostfix, jpegPostfix)
		if err := convertPDFToJPG(entrieContent.Name(), outputName); err != nil {
			panic(err)
		}
	}
}

func convertPDFToJPG(inputFileName, outputFileName string) error {
	buf, err := bimg.Read(inputFileName)
	if err != nil {
		return err
	}

	image, err := bimg.NewImage(buf).Convert(bimg.JPEG)
	if err != nil {
		return err
	}

	log.Println(image)
	return bimg.Write(outputFileName, image)
}
