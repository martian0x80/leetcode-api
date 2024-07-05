package queries

import "io"

// Yes, I named it IQuery, because I am still a C# dev at heart

type IQuery interface {
	GetQuery() io.Reader
}
