package scalar

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"io"
	"strconv"
)

/*
Token:
	model: github.com/d3v-friends/go-graphql/typ.Token
*/

func MarshalToken(b string) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = w.Write([]byte(strconv.Quote(b)))
	})
}

func UnmarshalToken(v interface{}) (string, error) {
	switch t := v.(type) {
	case string:
		return t, nil
	case []byte:
		return string(t), nil
	default:
		var err = fmt.Errorf("invalid Token scalar")
		return "", err
	}
}
