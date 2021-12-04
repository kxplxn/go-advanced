package _0204_channels_

import (
	usingChannels "go-advanced/04-channels/01-usingChannels"
	unbufferedChannels "go-advanced/04-channels/02-bufferedAndUnbuffered/01-unbufferedChannels"
	bufferedChannels "go-advanced/04-channels/02-bufferedAndUnbuffered/02-bufferedChannels"
	rangeAndClose "go-advanced/04-channels/03-rangeAndClose"
	selectStatement "go-advanced/04-channels/04-selectStatement"
	nonblockingSelect "go-advanced/04-channels/05-nonblockingSelect"
	channelsAndWaitGroups "go-advanced/04-channels/06-channelsAndWaitGroups"
	pipelines "go-advanced/04-channels/07-pipelines"
)

var Demos = struct {
	GoroutinesAndChannels func()
	UnbufferedChannels    func()
	BufferedChannels      func()
	RangeAndClose         func()
	SelectStatement       func()
	NonblockingSelect     func()
	ChannelsAndWaitGroups func()
	Pipelines             func()
}{
	GoroutinesAndChannels: usingChannels.Demo,
	UnbufferedChannels:    unbufferedChannels.Demo,
	BufferedChannels:      bufferedChannels.Demo,
	RangeAndClose:         rangeAndClose.Demo,
	SelectStatement:       selectStatement.Demo,
	NonblockingSelect:     nonblockingSelect.Demo,
	ChannelsAndWaitGroups: channelsAndWaitGroups.Demo,
	Pipelines:             pipelines.Demo,
}
