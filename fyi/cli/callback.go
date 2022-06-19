package cli

import (
	"github.com/kyokomi/emoji/v2"
	"github.com/spf13/cobra"

	"fyi/pkg"
)

const callback = "回电话！！！！！！"

var callbackCmd = &cobra.Command{
	Use:   "callback",
	Short: emoji.Sprintf(":telephone_receiver:回电话！"),
	Long:  emoji.Sprintf(`:telephone_receiver:回电话！！！`),
	Run: func(cmd *cobra.Command, args []string) {
		pkg.SendSMS(callback)
	},
}
