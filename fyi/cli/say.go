package cli

import (
	"fmt"
	"fyi/pkg"
	"github.com/kyokomi/emoji/v2"
	"github.com/spf13/cobra"
)

const (
	say = "说"
)

var sayCmd = &cobra.Command{
	Use:   "say",
	Short: emoji.Sprintf(":scroll:说"),
	Long:  emoji.Sprintf(`:scroll:说！`),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("You must say something...")
			return
		}

		content := ""
		for _, arg := range args {
			content += arg
		}
		if len(content) > 24 {
			fmt.Println("You say too much, less than 8 kanji please...")
			return
		}

		pkg.SendSMS(say + content)
	},
}
