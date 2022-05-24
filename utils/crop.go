package utils

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"

	"github.com/muesli/smartcrop"
	"github.com/muesli/smartcrop/nfnt"
)

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func CropImage(path string) error {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("can't open", err)
		return err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		fmt.Println("can't decode", err)
		return err
	}

	analyzer := smartcrop.NewAnalyzer(nfnt.NewDefaultResizer())
	topCrop, err := analyzer.FindBestCrop(img, 250, 250)
	if err != nil {
		fmt.Println("can't find best crop", err)
		return err
	}

	fmt.Printf("Top crop: %+v\n", topCrop)
	croppedimg := img.(SubImager).SubImage(topCrop)

	newCroppedFilePath := filepath.Dir(path) + "/cropped_" + filepath.Base(path)
	cf, err := os.Create(newCroppedFilePath)
	if err != nil {
		fmt.Println("can't create new file", err)
		return err
	}
	cf.Close()
	nf, err := os.OpenFile(newCroppedFilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("can't create new file", err)
		return err
	}
	defer nf.Close()

	opt := jpeg.Options{
		Quality: 90,
	}
	err = jpeg.Encode(nf, croppedimg, &opt)
	if err != nil {
		fmt.Println("can't encode jpeg", err)
		return err
	}

	return nil
}
