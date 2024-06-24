package imageproc

import (
	"image"

	"gocv.io/x/gocv"
)

// 画像内の長方形を検出
func DetectRectangle(img gocv.Mat) image.Rectangle {
	// ここにOpenCVを使った画像内の長方形検出ロジックを実装
	return image.Rect(50, 50, 200, 200) // 仮の座標
}
