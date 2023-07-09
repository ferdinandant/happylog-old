package core

type Color = string

// Need to use "\033" instead of "\e"
// - https://github.com/golang/go/issues/43337
// - https://www.ing.iac.es/~docs/external/bash/abs-guide/colorizing.html
// - https://i.stack.imgur.com/9UVnC.png
const (
	FlagReset Color = "\033[0m"

	FlagColorBgBoldBrightBlack Color = "\033[1;100m"
	FlagColorBgBoldRed         Color = "\033[1;41m"
	FlagColorBgBoldGreen       Color = "\033[1;42m"
	FlagColorBgBoldYellow      Color = "\033[1;43m"
	FlagColorBgBoldBlue        Color = "\033[1;44m"
	FlagColorBgBoldMagenta     Color = "\033[1;45m"

	FlagColorFgBoldBrightBlack Color = "\033[1;90m"
	FlagColorFgBoldRed         Color = "\033[1;31m"
	FlagColorFgBoldGreen       Color = "\033[1;32m"
	FlagColorFgBoldYellow      Color = "\033[1;33m"
	FlagColorFgBoldBlue        Color = "\033[1;34m"
	FlagColorFgBoldMagenta     Color = "\033[1;35m"

	FlagColorFgFaintBrightBlack Color = "\033[2;90m"
	FlagColorFgFaintBrightWhite Color = "\033[2;97m"
	FlagColorFgFaintRed         Color = "\033[2;31m"
	FlagColorFgFaintGreen       Color = "\033[2;32m"
	FlagColorFgFaintYellow      Color = "\033[2;33m"
	FlagColorFgFaintBlue        Color = "\033[2;34m"
	FlagColorFgFaintMagenta     Color = "\033[2;35m"

	FlagColorFgBrightWhite Color = "\033[97m"
	FlagColorFgBrightBlack Color = "\033[90m"
	FlagColorFgBlack       Color = "\033[30m"
	FlagColorFgRed         Color = "\033[31m"
	FlagColorFgGreen       Color = "\033[32m"
	FlagColorFgYellow      Color = "\033[33m"
	FlagColorFgBlue        Color = "\033[34m"
	FlagColorFgMagenta     Color = "\033[35m"
)
