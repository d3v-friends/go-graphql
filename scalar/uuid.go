package scalar

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/d3v-friends/go-tools/typ"
	"io"
	"strconv"
)

/*
UUID:
	model: github.com/d3v-friends/go-graphql/scalar.UUID
*/

func MarshalUUID(b typ.UUID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = w.Write([]byte(strconv.Quote(b.String())))
	})
}

func UnmarshalUUID(v interface{}) (typ.UUID, error) {
	switch t := v.(type) {
	case string:
		return typ.UUID(t), nil
	case []byte:
		return typ.UUID(t), nil
	default:
		var err = fmt.Errorf("invalid UUID scalar")
		return "", err
	}
}
