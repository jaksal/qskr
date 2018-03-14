package qskr

import (
	"testing"
)

func TestConvert(t *testing.T) {

	// result string get from http://code.cside.com/3rdpage/us/url/converter.html
	cases := []struct {
		in  string
		out string
	}{
		{
			"가나다 abc ABC 123 ~!@#$%^&*()_+`-=[]{}|;':,./<>?",
			"%B0%A1%B3%AA%B4%D9+abc+ABC+123+%7E%21%40%23%24%25%5E%26%2A%28%29_%2B%60-%3D%5B%5D%7B%7D%7C%3B%27%3A%2C.%2F%3C%3E%3F",
		},
	}

	for _, c := range cases {
		result := URLEncode(c.in)
		if result != c.out {
			t.Errorf("src=%s\nresult1=%s\nresult2=%s\n", c.in, c.out, result)
		}
		result2, err := URLDecode(result)
		if err != nil {
			t.Error("decode err", err)
		} else {
			if result2 != c.in {
				t.Error("decode err2", result2, c.in)
			}
		}

	}
}
