package conditional

// Int returns val1 if condition, otherwise val2
func Int(condition bool, val1, val2 int) int {
	if condition {
		return val1
	}
	return val2
}

// Int8 returns val1 if condition, otherwise int8
func Int8(condition bool, val1, val2 int8) int8 {
	if condition {
		return val1
	}
	return val2
}

// Int16 returns val1 if condition, otherwise val2
func Int16(condition bool, val1, val2 int16) int16 {
	if condition {
		return val1
	}
	return val2
}

// Int32 returns val1 if condition, otherwise val2
func Int32(condition bool, val1, val2 int32) int32 {
	if condition {
		return val1
	}
	return val2
}

// Int64 returns val1 if condition, otherwise val2
func Int64(condition bool, val1, val2 int64) int64 {
	if condition {
		return val1
	}
	return val2
}

// UInt returns val1 if condition, otherwise val2
func UInt(condition bool, val1, val2 uint) uint {
	if condition {
		return val1
	}
	return val2
}

// UInt8 returns val1 if condition, otherwise int8
func UInt8(condition bool, val1, val2 uint8) uint8 {
	if condition {
		return val1
	}
	return val2
}

// UInt16 returns val1 if condition, otherwise val2
func UInt16(condition bool, val1, val2 uint16) uint16 {
	if condition {
		return val1
	}
	return val2
}

// UInt32 returns val1 if condition, otherwise val2
func UInt32(condition bool, val1, val2 uint32) uint32 {
	if condition {
		return val1
	}
	return val2
}

// UInt64 returns val1 if condition, otherwise val2
func UInt64(condition bool, val1, val2 uint64) uint64 {
	if condition {
		return val1
	}
	return val2
}

// String returns val1 if condition, otherwise val2
func String(condition bool, val1, val2 string) string {
	if condition {
		return val1
	}
	return val2
}

// Interface returns val1 if condition, otherwise val2
func Interface(condition bool, val1, val2 interface{}) interface{} {
	if condition {
		return val1
	}
	return val2
}

// Rune returns val1 if condition, otherwise val2
func Rune(condition bool, val1, val2 rune) rune {
	if condition {
		return val1
	}
	return val2
}

// Byte returns val1 if condition, otherwise val2
func Byte(condition bool, val1, val2 byte) byte {
	if condition {
		return val1
	}
	return val2
}

// Bytes returns val1 if condition, otherwise val2
func Bytes(condition bool, val1, val2 []byte) []byte {
	if condition {
		return val1
	}
	return val2
}

// Float32 returns val1 if condition, otherwise val2
func Float32(condition bool, val1, val2 float32) float32 {
	if condition {
		return val1
	}
	return val2
}

// Float64 returns val1 if condition, otherwise val2
func Float64(condition bool, val1, val2 float64) float64 {
	if condition {
		return val1
	}
	return val2
}

// Float returns val1 if condition, otherwise val2
// synonym to FLoat64
func Float(condition bool, val1, val2 float64) float64 {
	return Float64(condition, val1, val2)
}

// Complex64 returns val1 if condition, otherwise val2
func Complex64(condition bool, val1, val2 complex64) complex64 {
	if condition {
		return val1
	}
	return val2
}

// Complex128 returns val1 if condition, otherwise val2
func Complex128(condition bool, val1, val2 complex128) complex128 {
	if condition {
		return val1
	}
	return val2
}
