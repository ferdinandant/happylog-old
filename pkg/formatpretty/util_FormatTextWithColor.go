package formatpretty

func FormatTextWithColor(color Color, str string) string {
	return string(color) + str + string(FlagReset)
}
