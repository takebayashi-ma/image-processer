package imageproc

import (
	"gocv.io/x/gocv"
)

// BrightenImage 明るさの調整
func BrightenImage(img gocv.Mat) gocv.Mat {
	// 明るさを調整するために、画像に一定の値を加算する
	brightImg := gocv.NewMat()
	gocv.AddWeighted(img, 1.2, gocv.NewMatWithSize(img.Rows(), img.Cols(), img.Type()), 0, 50, &brightImg) // 1.2倍して50を加算
	return brightImg
}
