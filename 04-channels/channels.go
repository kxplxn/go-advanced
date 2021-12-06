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

func Demos() {
	usingChannels.Demo()
	unbufferedChannels.Demo()
	bufferedChannels.Demo()
	rangeAndClose.Demo()
	selectStatement.Demo()
	nonblockingSelect.Demo()
	channelsAndWaitGroups.Demo()
	pipelines.Demo()
}
