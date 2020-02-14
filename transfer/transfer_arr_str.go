package transfer

import (
	"bytes"
	"strconv"
	"strings"
	"sync"
)

var (
	bfPool = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer([]byte{})
		},
	}
)

// []int32 以sep合并为字符串
func JoinInt32Slice(array []int32, sep string) (s string) {

	length := len(array)

	if 0 == length {
		return ""
	}

	if 1 == length {
		return strconv.FormatInt(int64(array[0]), 10)
	}
	buf := bfPool.Get().(*bytes.Buffer)
	for _, e := range array {
		buf.WriteString(strconv.FormatInt(int64(e), 10))
		buf.WriteString(sep)
	}

	if buf.Len() > 0 {
		buf.Truncate(buf.Len() - 1)
	}

	s = buf.String()
	buf.Reset()
	bfPool.Put(buf)

	return
}

// int32位十进制的字符串，以sep分割为[]int32
func SplitInt32s(s, sep string) ([]int32, error) {
	if "" == s {
		return nil, nil
	}
	sArr := strings.Split(s, sep)
	iArr := make([]int32, 0, len(sArr))

	for _, es := range sArr {
		ei, err := strconv.ParseInt(es, 10, 32)
		if nil != err {
			return nil, err
		}

		iArr = append(iArr, int32(ei))
	}

	return iArr, nil
}

func JoinInt64Slice(array []int64, sep string) (s string) {
	length := len(array)
	if 0 == length {
		return ""
	}

	if 1 == length {
		return strconv.FormatInt(array[0], 10)
	}

	buf := bfPool.Get().(*bytes.Buffer)

	for _, e := range array {
		buf.WriteString(strconv.FormatInt(e, 10))
		buf.WriteString(sep)
	}

	if buf.Len() > 0 {
		buf.Truncate(buf.Len() - 1)
	}

	s = buf.String()
	buf.Reset()
	bfPool.Put(buf)
	return
}

func SplitInt64s(s, sep string) ([]int64, error) {
	if "" == s {
		return nil, nil
	}

	sArr := strings.Split(s, sep)
	iArr := make([]int64, 0, len(sArr))

	for _, es := range sArr {
		ei, err := strconv.ParseInt(es, 10, 64)
		if nil != err {
			return nil, err
		}
		iArr = append(iArr, ei)
	}

	return iArr, nil
}
