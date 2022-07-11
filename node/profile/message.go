package profile

import "cloudCli/channel"

/**
 *
 * @author jensen.chen
 * @date 2022/7/11
 */
const MESSAGE_PROFILE_RESET = "profile_reset" //重置

func BuildRestCommand() *channel.CommandMessage {
	return &channel.CommandMessage{Name: MESSAGE_PROFILE_RESET}
}
