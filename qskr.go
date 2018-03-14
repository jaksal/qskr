package qskr

import (
	"bytes"
	"fmt"
	"net/url"

	"unicode"

	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
)

// URLEncode euc-kr string queryescape
func URLEncode(str string) string {
	var buf bytes.Buffer
	for _, r := range str {
		switch {
		case unicode.IsNumber(r) ||
			unicode.Is(unicode.Latin, r) ||
			r == '-' ||
			r == '_' ||
			r == '.':
			buf.WriteString(string(r))

		case unicode.IsSpace(r):
			buf.WriteString("+")

		case r == '$' ||
			r == '&' ||
			r == '+' ||
			r == ',' ||
			r == '/' ||
			r == ':' ||
			r == ';' ||
			r == '=' ||
			r == '?' ||
			r == '@':
			buf.WriteString(fmt.Sprintf("%%%s%s", string("0123456789ABCDEF"[byte(r)>>4]), string("0123456789ABCDEF"[byte(r)&15])))

		default:
			rstr, n, err := transform.String(korean.EUCKR.NewEncoder(), string(r))
			if err != nil {
				fmt.Println("euc-kr encode err!", string(r), err)
				return ""
			}
			if n > 1 {
				buf.WriteString(fmt.Sprintf("%%%X%%%X", rstr[0], rstr[1]))
			} else {
				buf.WriteString(fmt.Sprintf("%%%X", rstr[0]))
			}
		}
	}
	return buf.String()
}

// URLDecode euc-kr string queryunescape
func URLDecode(str string) (string, error) {
	dst, err := url.QueryUnescape(str)
	if err != nil {
		return "", err
	}
	result, _, err := transform.String(korean.EUCKR.NewDecoder(), dst)
	if err != nil {
		return "", fmt.Errorf("decode err=%s", err)
	}
	return result, nil
}
