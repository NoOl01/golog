package unsafe_conv

import (
	"reflect"
	"testing"
)

func TestStringToBytes(t *testing.T) {
	tests := []struct {
		str  string
		want []byte
	}{
		{str: "hello world", want: []byte("hello world")},
		{str: "привет мир", want: []byte("привет мир")},
		{str: "你好，世界", want: []byte("你好，世界")},
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			got := StringToBytes(test.str)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("StringToBytes(%q) = %v, want %v", test.str, got, test.want)
			}

			if s := string(got); s != test.str {
				t.Errorf("string(got) = %v, want %v", got, test.str)
			}
		})
	}
}
