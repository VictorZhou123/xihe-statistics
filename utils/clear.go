package utils

import "unsafe"

func ClearStringMemory(s string) {
	if len(s) <= 1 {
		return
	}

	bs := *(*[]byte)(unsafe.Pointer(&s))
	for i := 0; i < len(bs); i++ {
		bs[i] = 0
	}
}
