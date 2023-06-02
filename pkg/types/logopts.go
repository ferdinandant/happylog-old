package types

import "time"

type LogOpts = struct {
	Level   *Level
	AppName *string
	Now     *time.Time
	Msg     *string
	Ctx     *interface{}
}
