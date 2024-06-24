package main

import (
	"os"

	"image-processor/imageproc"

	"gocv.io/x/gocv"
)

func main() {
	// 画像ファイルの読み込み
	img, err := imageproc.ReadImage("input.heic")
	if err != nil {
		panic(err)
	}
	defer img.Close()

	// 画像の切り抜き
	rect := imageproc.DetectRectangle(img)
	croppedImg := img.Region(rect)

	// 水平・垂直の修正
	alignedImg := imageproc.AlignImage(croppedImg)

	// 影の除去
	noShadowImg := imageproc.RemoveShadows(alignedImg)

	// 明るさの調整
	finalImg := imageproc.BrightenImage(noShadowImg)

	// 画像の保存
	outFile, err := os.Create("output.jpg")
	if err != nil {
		panic("画像ファイルを保存できませんでした")
	}
	defer outFile.Close()
	gocv.IMWrite("output.jpg", finalImg)
}
