package utils

import (
	fr "github.com/Kagami/go-face"
)

func FloatSliceToDescriptor(points []float32) fr.Descriptor {
	var dots fr.Descriptor
	for k, v := range points {
		dots[k] = v
	}
	return dots
}
