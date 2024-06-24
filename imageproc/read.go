package imageproc

import (
	"bytes"
	"errors"
	"image"
	"image/jpeg"
	"os"

	"github.com/strukturag/libheif/go/heif"
	"gocv.io/x/gocv"
)

// ReadImage 画像ファイルの読み込み（JPEGおよびHEIC対応）
func ReadImage(path string) (gocv.Mat, error) {
	file, err := os.Open(path)
	if err != nil {
		return gocv.Mat{}, err
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(file)

	_, format, err := image.Decode(buf)
	if err != nil {
		return gocv.Mat{}, err
	}

	if format == "heif" || format == "heic" {
		return readHEIC(buf.Bytes())
	} else if format == "jpeg" || format == "jpg" {
		return readJPEG(buf.Bytes())
	}

	return gocv.Mat{}, errors.New("対応していない画像形式です")
}

func readJPEG(data []byte) (gocv.Mat, error) {
	img, err := jpeg.Decode(bytes.NewReader(data))
	if err != nil {
		return gocv.Mat{}, err
	}

	mat, err := gocv.ImageToMatRGB(img)
	if err != nil {
		return gocv.Mat{}, err
	}
	return mat, nil
}

func readHEIC(data []byte) (gocv.Mat, error) {
	c, err := heif.NewContext()

	err = c.ReadFromMemory(data)
	if err != nil {
		return gocv.Mat{}, err
	}
	handle, err := c.GetPrimaryImageHandle()
	if err != nil {
		return gocv.Mat{}, err
	}
	img, err := handle.DecodeImage(heif.ColorspaceRGB, heif.ChromaInterleavedRGBA, nil)

	image, err := img.GetImage()
	mat, err := gocv.ImageToMatRGB(image)
	if err != nil {
		return gocv.Mat{}, err
	}

	return mat, nil
}
