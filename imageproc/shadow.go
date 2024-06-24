package imageproc

import (
	"gocv.io/x/gocv"
)

// RemoveShadows 影の除去
func RemoveShadows(img gocv.Mat) gocv.Mat {
	hsv := gocv.NewMat()
	defer hsv.Close()

	// 画像をHSVに変換
	gocv.CvtColor(img, &hsv, gocv.ColorBGRToHSV)

	// HSVチャンネルに分離
	channels := gocv.Split(hsv)
	defer func() {
		for _, c := range channels {
			c.Close()
		}
	}()

	// 影の除去のために、Vチャンネルを調整
	for y := 0; y < channels[2].Rows(); y++ {
		for x := 0; x < channels[2].Cols(); x++ {
			v := channels[2].GetUCharAt(y, x)
			if v < 128 {
				channels[2].SetUCharAt(y, x, 255)
			}
		}
	}

	// 調整後のチャンネルをマージして、BGRに戻す
	gocv.Merge(channels, &hsv)
	result := gocv.NewMat()
	gocv.CvtColor(hsv, &result, gocv.ColorHSVToBGR)

	return result
}
