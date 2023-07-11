package colors

import "github.com/ferdinandant/happylog/pkg/core"

type ThemeColorScheme struct {
	BgBold   Color
	FgBold   Color
	FgNormal Color
	FgFaint  Color
}

func GetThemeColorScheme(level core.Level) *ThemeColorScheme {
	var bgBold Color
	var fgBold Color
	var fgNormal Color
	var fgFaint Color

	// Multiplex by level
	switch level {
	case core.LevelTrace:
		bgBold = FlagColorBgBoldBrightBlack
		fgBold = FlagColorFgBoldBrightBlack
		fgNormal = FlagColorFgBrightBlack
		fgFaint = FlagColorFgFaintBrightBlack
	case core.LevelDebug:
		bgBold = FlagColorBgBoldBlue
		fgBold = FlagColorFgBoldBlue
		fgNormal = FlagColorFgBlue
		fgFaint = FlagColorFgFaintBlue
	case core.LevelInfo:
		bgBold = FlagColorBgBoldGreen
		fgBold = FlagColorFgBoldGreen
		fgNormal = FlagColorFgGreen
		fgFaint = FlagColorFgFaintGreen
	case core.LevelWarn:
		bgBold = FlagColorBgBoldYellow
		fgBold = FlagColorFgBoldYellow
		fgNormal = FlagColorFgYellow
		fgFaint = FlagColorFgFaintYellow
	case core.LevelError:
		bgBold = FlagColorBgBoldRed
		fgBold = FlagColorFgBoldRed
		fgNormal = FlagColorFgRed
		fgFaint = FlagColorFgFaintRed
	case core.LevelFatal:
		bgBold = FlagColorBgBoldMagenta
		fgBold = FlagColorFgBoldMagenta
		fgNormal = FlagColorFgMagenta
		fgFaint = FlagColorFgFaintMagenta
	}

	// Construct color scheme
	return &ThemeColorScheme{
		BgBold:   bgBold,
		FgBold:   fgBold,
		FgNormal: fgNormal,
		FgFaint:  fgFaint,
	}
}
