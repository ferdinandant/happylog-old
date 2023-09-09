package ctxparser

type ParseResultCtx struct {
	isLiteral              bool
	isAllDescendantLiteral bool
}

var LiteralParseResultCtx *ParseResultCtx = &ParseResultCtx{
	isLiteral:              true,
	isAllDescendantLiteral: true,
}

var ErrorParseResultCtx *ParseResultCtx = &ParseResultCtx{
	isLiteral:              true,
	isAllDescendantLiteral: false,
}
