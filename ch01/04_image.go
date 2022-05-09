//lissajous 产生随机利萨茹图形的GIF动画
package main

import (
	"bytes"
	"image"
	"image/color"
	"image/gif"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // 画板中的第一种颜色
	blackIndex = 1 // 画板中的下一种颜色
)

func main() {
	//如果是rand.Int，此时获取的随机数，都是重复的一些随机数据
	//下面这种写法使用了当前时间播种伪随机数生成器，可以保证每次随机都是随机的
	rand.Seed(time.Now().UTC().UnixNano())
	//改为直接输出字节文件
	buf := &bytes.Buffer{}
	lissajous(buf)
	if err := ioutil.WriteFile("output.gif", buf.Bytes(), 0666); err != nil {
		panic(err)
	}
}

func lissajous(out io.Writer) {
	const ( //const用于给常量命名
		cycles  = 5     // 完整的x振荡器变化的个数
		res     = 0.001 // 角度分辨率
		size    = 100   // 图象画布包含 [-size..+size]
		nframes = 64    // 动画中的帧数
		delay   = 8     // 以10ms为单位的帧间延迟
	)
	freq := rand.Float64() * 3.0 // y振荡器的相对频率
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference 相位差
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		//append的作用：在原切片末尾添加元素
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // 注意：忽略编码错误
}
