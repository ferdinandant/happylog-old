package colors

func FormatTextWithColor(color Color, str string) string {
	return color + str + FlagReset
}