package ctxparser

type ParseResultCtx struct {
	isLiteral bool
}

var LiteralParseResultCtx *ParseResultCtx = &ParseResultCtx{
	isLiteral: true,
}

var StructParseResultCtx *ParseResultCtx = &ParseResultCtx{
	isLiteral: false,
}

var ErrorParseResultCtx *ParseResultCtx = &ParseResultCtx{
	isLiteral: false,
}
