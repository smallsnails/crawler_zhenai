package defs

type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
	Flag     string
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
