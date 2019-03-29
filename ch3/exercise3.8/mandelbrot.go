package mandelbrot

import (
	"image/color"
	"math"
	"math/big"
	"math/cmplx"
)

func mandelbrot64(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + complex64(z)
		if cmplx.Abs(complex128(v)) > 2 {
			if n > 50 {
				return color.RGBA{100, 0, 0, 255}
			}
			scale := math.Log(float64(n)) / math.Log(float64(iterations))
			return color.RGBA{0, 0, 255 - uint8(scale*255), 255}
		}
	}
	return color.Black
}

func mandelbrot128(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			if n > 50 {
				return color.RGBA{100, 0, 0, 255}
			}
			scale := math.Log(float64(n)) / math.Log(float64(iterations))
			return color.RGBA{0, 0, 255 - uint8(scale*255), 255}
		}
	}
	return color.Black
}

func mandelbrotBigFloat(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	zR := (&big.Float{}).SetFloat64(real(z))
	zI := (&big.Float{}).SetFloat64(imag(z))
	var vR, vI = &big.Float{}, &big.Float{}
	for i := uint8(0); i < iterations; i++ {
		vR2, vI2 := &big.Float{}, &big.Float{}
		vR2.Mul(vR, vR).Sub(vR2, (&big.Float{}).Mul(vI, vI)).Add(vR2, zR)
		vI2.Mul(vR, vI).Mul(vI2, big.NewFloat(2)).Add(vI2, zI)
		vR, vI = vR2, vI2
		squareSum := &big.Float{}
		squareSum.Mul(vR, vR).Add(squareSum, (&big.Float{}).Mul(vI, vI))
		if squareSum.Cmp(big.NewFloat(4)) == 1 {
			if i > 50 {
				return color.RGBA{100, 0, 0, 255}
			}
			scale := math.Log(float64(i)) / math.Log(float64(iterations))
			return color.RGBA{0, 0, 255 - uint8(scale*255), 255}
		}
	}
	return color.Black
}

func mandelbrotBigRat(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	zR := (&big.Rat{}).SetFloat64(real(z))
	zI := (&big.Rat{}).SetFloat64(imag(z))
	var vR, vI = &big.Rat{}, &big.Rat{}
	for i := uint8(0); i < iterations; i++ {
		vR2, vI2 := &big.Rat{}, &big.Rat{}
		vR2.Mul(vR, vR).Sub(vR2, (&big.Rat{}).Mul(vI, vI)).Add(vR2, zR)
		vI2.Mul(vR, vI).Mul(vI2, big.NewRat(2, 1)).Add(vI2, zI)
		vR, vI = vR2, vI2
		squareSum := &big.Rat{}
		squareSum.Mul(vR, vR).Add(squareSum, (&big.Rat{}).Mul(vI, vI))
		if squareSum.Cmp(big.NewRat(4, 1)) == 1 {
			if i > 50 {
				return color.RGBA{100, 0, 0, 255}
			}
			scale := math.Log(float64(i)) / math.Log(float64(iterations))
			return color.RGBA{0, 0, 255 - uint8(scale*255), 255}
		}
	}
	return color.Black
}
