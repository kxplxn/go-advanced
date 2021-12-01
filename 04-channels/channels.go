package _0204_channels_

import (
	usingChannels "go-advanced/04-channels/01-usingChannels"
	unbufferedChannels "go-advanced/04-channels/02-bufferedAndUnbuffered/01-unbufferedChannels"
	bufferedChannels "go-advanced/04-channels/02-bufferedAndUnbuffered/02-bufferedChannels"
)

var Demos = struct {
	GoroutinesAndChannels func()
	UnbufferedChannels    func()
	BufferedChannels      func()
}{
	GoroutinesAndChannels: usingChannels.Demo,
	UnbufferedChannels:    unbufferedChannels.Demo,
	BufferedChannels:      bufferedChannels.Demo,
}
