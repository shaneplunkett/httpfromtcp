package request

import (
	"fmt"
	"io"
	"strings"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	r, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	s := strings.Split(string(r), "\r\n")
	rl := strings.Split(s[0], " ")
	if len(rl) != 3 {
		return nil, fmt.Errorf("invalid read line length")
	}

	method := rl[0]
	if len(method) == 0 {
		return nil, fmt.Errorf("method does not exist")
	}
	if method != "GET" && method != "POST" {
		return nil, fmt.Errorf("Unsupported method: %s", method)
	}

	target := rl[1]
	if len(target) == 0 {
		return nil, fmt.Errorf("Request Target does not exist")
	}
	if !strings.HasPrefix(target, "/") {
		return nil, fmt.Errorf("Unsupported Target: %s", target)
	}
	if rl[2] != "HTTP/1.1" {
		return nil, fmt.Errorf("Unsupported HTTP Version: %s", rl[2])
	}

	return &Request{
		RequestLine: RequestLine{
			Method:        method,
			RequestTarget: target,
			HttpVersion:   "1.1",
		},
	}, nil

}
