package numbers

import "strconv"

// ParseUint8 -
func ParseUint8(source string) (uint8, error) {
	ui64, err := strconv.ParseUint(source, 10, 8)
	if err != nil {
		return 0, err
	}
	return uint8(ui64), nil
}
