package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// MyImage は image.Image インタフェースを実装するカスタム画像型です。
type MyImage struct {
	width, height int
}

// Bounds は画像の境界を返します。
func (m MyImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.width, m.height)
}

// ColorModel は色のモデルを返します。
func (m MyImage) ColorModel() color.Model {
	return color.RGBAModel
}

// At は指定された座標の色を返します。
func (m MyImage) At(x, y int) color.Color {
	// 適切な色の値を計算
	v := uint8((x + y) / 2) // 例として (x+y)/2 を使用
	return color.RGBA{v, v, 255, 255}
}

func main() {
	// 200x200 の MyImage インスタンスを作成し、pic.ShowImage に渡します。
	m := MyImage{width: 200, height: 200}
	pic.ShowImage(m)
}
