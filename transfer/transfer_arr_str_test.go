package transfer

import (
	"testing"
)

var testInt32AndStringData = map[string][]int32{
	"23,34,56,7": []int32{23, 34, 56, 7},
	"":           []int32{},
	"22":         []int32{22},
	"a":          []int32{96},
	"14,":        []int32{14},
}

func TestJoinInt32Slice(t *testing.T) {
	for str, arr := range testInt32AndStringData {
		int32String := JoinInt32Slice(arr, ",")
		if str != int32String {
			t.Errorf("%s Failed, Expect %s, But %s", str, str, int32String)
		}
	}
}

func TestSplitInt32s(t *testing.T) {
	for str, arr := range testInt32AndStringData {
		int32Slice, err := SplitInt32s(str, ",")
		if nil != err {
			t.Fatal(err)
		}

		if len(arr) != len(int32Slice) {
			t.Errorf("%s Failed, Except length %d, But %d", str, len(arr), len(int32Slice))
		}
		for i, a := range arr {
			if a != int32Slice[i] {
				t.Errorf("%s Failed, the %d-th Except %d, But %d", str, i, a, int32Slice[i])
				break
			}
		}
	}
}

/*nums64 := []int64{33, 55, 66, 77}

int64String := JoinInt64Slice(nums64, "-")
fmt.Println(int64String)
fmt.Println(SplitInt64s(int64String, "-"))*/
