package colors

import "github.com/ferdinandant/happylog/pkg/core"

type ColorScheme struct {
	BgBold   Color
	FgBold   Color
	FgNormal Color
	FgFaint  Color
}

var ColorSchemeMap = map[core.Level]*ColorScheme{
	core.LevelTrace: {
		BgBold:   FlagColorBgBoldBrightBlack,
		FgBold:   FlagColorFgBoldBrightBlack,
		FgNormal: FlagColorFgBrightBlack,
		FgFaint:  FlagColorFgFaintBrightBlack,
	},
	core.LevelDebug: {
		BgBold:   FlagColorBgBoldBlue,
		FgBold:   FlagColorFgBoldBlue,
		FgNormal: FlagColorFgBlue,
		FgFaint:  FlagColorFgFaintBlue,
	},
	core.LevelInfo: {
		BgBold:   FlagColorBgBoldGreen,
		FgBold:   FlagColorFgBoldGreen,
		FgNormal: FlagColorFgGreen,
		FgFaint:  FlagColorFgFaintGreen,
	},
	core.LevelWarn: {
		BgBold:   FlagColorBgBoldYellow,
		FgBold:   FlagColorFgBoldYellow,
		FgNormal: FlagColorFgYellow,
		FgFaint:  FlagColorFgFaintYellow,
	},
	core.LevelError: {
		BgBold:   FlagColorBgBoldRed,
		FgBold:   FlagColorFgBoldRed,
		FgNormal: FlagColorFgRed,
		FgFaint:  FlagColorFgFaintRed,
	},
	core.LevelFatal: {
		BgBold:   FlagColorBgBoldMagenta,
		FgBold:   FlagColorFgBoldMagenta,
		FgNormal: FlagColorFgMagenta,
		FgFaint:  FlagColorFgFaintMagenta,
	},
}
