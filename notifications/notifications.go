package notifications

import (
	log "github.com/Sirupsen/logrus"
	"github.com/fzzy/radix/redis"
	"github.com/jonas747/yagpdb/bot"
	"github.com/jonas747/yagpdb/common"
	"github.com/jonas747/yagpdb/web"
)

type Plugin struct{}

func RegisterPlugin() {
	plugin := &Plugin{}
	bot.RegisterPlugin(plugin)
	web.RegisterPlugin(plugin)
}

func (p *Plugin) Name() string {
	return "Notifications"
}

func (p *Plugin) InitBot() {
	common.BotSession.AddHandler(bot.CustomGuildCreate(HandleGuildCreate))
	common.BotSession.AddHandler(bot.CustomGuildMemberAdd(HandleGuildMemberAdd))
	common.BotSession.AddHandler(bot.CustomGuildMemberRemove(HandleGuildMemberRemove))
	common.BotSession.AddHandler(bot.CustomChannelUpdate(HandleChannelUpdate))
}

type Config struct {
	JoinServerEnabled bool   `json:"join_server_enabled`
	JoinServerChannel string `json:"join_server_channel"`
	JoinServerMsg     string `json:"join_server_msg"`

	JoinDMEnabled bool   `json:"join_dm_enabled"`
	JoinDMMsg     string `json:"join_dm_msg"`

	LeaveEnabled bool   `json:"leave_enabled"`
	LeaveChannel string `json:"leave_channel"`
	LeaveMsg     string `json:"leave_msg"`

	TopicEnabled bool   `json:"topic_enabled"`
	TopicChannel string `json:"topic_channel"`

	// Deprecated
	// Need to safely remove these fields
	PinEnabled bool   `json:"pin_enabled,omitempty"`
	PinChannel string `json:"pin_channel,omitempty"`
}

var DefaultConfig = &Config{}

func GetConfig(client *redis.Client, server string) *Config {
	var config *Config
	if err := common.GetRedisJson(client, "notifications/general:"+server, &config); err != nil {
		if _, ok := err.(*redis.CmdError); ok {
			log.WithError(err).WithField("guild", server).Error("Failed retrieving noifications config")
		}
		return DefaultConfig
	}
	return config
}
