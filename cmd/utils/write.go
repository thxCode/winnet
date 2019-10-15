package utils

import (
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/thxcode/winnet/pkg/converters"
)

func Output(w io.Writer, obj interface{}) error {
	if obj == nil {
		return nil
	}

	switch t := obj.(type) {
	case []byte:
		if len(t) == 0 {
			return nil
		}
		return fprint(w, converters.UnsafeBytesToString(t))
	case string:
		if len(t) == 0 {
			return nil
		}
		return fprint(w, t)
	default:
		if reflect.TypeOf(obj).Kind() == reflect.Slice {
			sb := &strings.Builder{}
			sl := reflect.ValueOf(obj)
			for i := 0; i < sl.Len(); i++ {
				sb.WriteString(fmt.Sprintln(sl.Index(i)))
			}
			return fprint(w, sb)
		}

		return fprint(w, t)
	}
}

func fprint(w io.Writer, obj interface{}) (err error) {
	_, err = fmt.Fprint(w, obj)
	if err != nil {
		err = fmt.Errorf("failed to output result: %v", err)
	}
	return
}
