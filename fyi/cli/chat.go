package cli

import (
	"github.com/kyokomi/emoji/v2"
	"github.com/spf13/cobra"

	"fyi/pkg"
)

const chat = "聊天"

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: emoji.Sprintf(":mobile_phone:聊天"),
	Long:  emoji.Sprintf(`:mobile_phone:聊天！`),
	Run: func(cmd *cobra.Command, args []string) {
		pkg.SendSMS(chat)
	},
}
