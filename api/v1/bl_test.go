package apiv1

import (
	"testing"
)

func TestGetPercent(t *testing.T) {
	newValue := GetPercent(2000, 25)
	if newValue != 500 {
		t.Fail()
	}
}

func TestGetZeroPercent(t *testing.T) {
	newValue := GetPercent(2000, 0)
	if newValue != 0 {
		t.Fail()
	}
}

func TestGetNegativePercent(t *testing.T) {
	newValue := GetPercent(2000, -25)
	if newValue != -500 {
		t.Fail()
	}
}
