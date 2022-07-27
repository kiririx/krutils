package copyx

import "testing"

func TestDeepCopy(t *testing.T) {
	type Src struct {
		A int
		B string
		C []int
		D map[string]int
		E struct {
			F int
			G string
			H []int
			I map[string]int
		}
	}
	type Dest struct {
		A int
		B string
		C []int
		D map[string]int
		E struct {
			F int
			G string
			H []int
			I map[string]int
		}
	}
	src := Src{
		A: 1,
		B: "2",
		C: []int{3, 4, 5},
		D: map[string]int{
			"6": 7,
			"8": 9,
		},
		E: struct {
			F int
			G string
			H []int
			I map[string]int
		}{
			F: 11,
			G: "12",
			H: []int{13, 14, 15},
			I: map[string]int{
				"16": 17,
				"18": 19,
			},
		},
	}
	dest := Dest{}
	err := DeepCopy(&src, &dest)
	if err != nil {
		t.Error(err)
	}
	t.Log(src)
	t.Log(dest)
}
