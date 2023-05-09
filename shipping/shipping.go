package shipping

const (
	VolumetricWeightDivisor5000 = 5000
	VolumetricWeightDivisor6000 = 6000
)

func VolumetricWeightCm(w, h, d, divisor float64) float64 {
	n := w * h * d

	if n == 0 {
		return n
	}

	return n / divisor
}

func VolumetricWeightMm(w, h, d int, divisor float64) float64 {
	return VolumetricWeightCm(float64(w)/10, float64(h)/10, float64(d)/10, divisor)
}
