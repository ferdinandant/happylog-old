package formatpretty

// Need to use "\033" instead of "\e"
// - https://github.com/golang/go/issues/43337
// - https://www.ing.iac.es/~docs/external/bash/abs-guide/colorizing.html
// - https://i.stack.imgur.com/9UVnC.png
const (
	FlagReset = "\033[0m"

	FlagColorBgBoldBrightBlack = "\033[1;100m"
	FlagColorBgBoldRed         = "\033[1;41m"
	FlagColorBgBoldGreen       = "\033[1;42m"
	FlagColorBgBoldYellow      = "\033[1;43m"
	FlagColorBgBoldBlue        = "\033[1;44m"
	FlagColorBgBoldMagenta     = "\033[1;45m"

	FlagColorFgBoldBrightBlack = "\033[1;90m"
	FlagColorFgBoldRed         = "\033[1;31m"
	FlagColorFgBoldGreen       = "\033[1;32m"
	FlagColorFgBoldYellow      = "\033[1;33m"
	FlagColorFgBoldBlue        = "\033[1;34m"
	FlagColorFgBoldMagenta     = "\033[1;35m"

	FlagColorFgFaintBrightBlack = "\033[2;90m"
	FlagColorFgFaintBrightWhite = "\033[2;97m"
	FlagColorFgFaintRed         = "\033[2;31m"
	FlagColorFgFaintGreen       = "\033[2;32m"
	FlagColorFgFaintYellow      = "\033[2;33m"
	FlagColorFgFaintBlue        = "\033[2;34m"
	FlagColorFgFaintMagenta     = "\033[2;35m"

	FlagColorFgBrightWhite = "\033[97m"
	FlagColorFgBrightBlack = "\033[90m"
	FlagColorFgBlack       = "\033[30m"
	FlagColorFgRed         = "\033[31m"
	FlagColorFgGreen       = "\033[32m"
	FlagColorFgYellow      = "\033[33m"
	FlagColorFgBlue        = "\033[34m"
	FlagColorFgMagenta     = "\033[35m"
)
