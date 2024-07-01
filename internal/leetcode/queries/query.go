package queries

import "io"

type IQuery interface {
	GetQuery() io.Reader
}
