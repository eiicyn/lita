package cli

import (
	"github.com/kyokomi/emoji/v2"
	"github.com/spf13/cobra"

	"fyi/pkg"
)

const drink = "喝"

var drinkCmd = &cobra.Command{
	Use:   "drink",
	Short: emoji.Sprintf(":sake:喝"),
	Long:  emoji.Sprintf(`:sake:喝！`),
}

var coffeeCmd = &cobra.Command{
	Use:   "coffee",
	Short: emoji.Sprintf(":coffee:咖啡"),
	Long:  emoji.Sprintf(`:coffee:咖啡！`),
	Run: func(cmd *cobra.Command, args []string) {
		pkg.SendSMS(drink + "咖啡")
	},
}

var teaCmd = &cobra.Command{
	Use:   "tea",
	Short: emoji.Sprintf(":teacup_without_handle:茶"),
	Long:  emoji.Sprintf(`:teacup_without_handle:茶！`),
	Run: func(cmd *cobra.Command, args []string) {
		pkg.SendSMS(drink + "茶")
	},
}

var milkTeaCmd = &cobra.Command{
	Use:   "milktea",
	Short: emoji.Sprintf(":bubble_tea:奶茶"),
	Long:  emoji.Sprintf(`:bubble_tea:奶茶！`),
	Run: func(cmd *cobra.Command, args []string) {
		pkg.SendSMS(drink + "奶茶")
	},
}

var porridgeCmd = &cobra.Command{
	Use:   "porridge",
	Short: emoji.Sprintf(":fondue:粥"),
	Long:  emoji.Sprintf(`:fondue:粥！`),
	Run: func(cmd *cobra.Command, args []string) {
		pkg.SendSMS(drink + "粥")
	},
}

var liquorCmd = &cobra.Command{
	Use:   "liquor",
	Short: emoji.Sprintf(":tropical_drink:酒"),
	Long:  emoji.Sprintf(`:tropical_drink:酒！`),
	Run: func(cmd *cobra.Command, args []string) {
		pkg.SendSMS(drink + "酒")
	},
}

func init() {
	drinkCmd.AddCommand(teaCmd)
	drinkCmd.AddCommand(milkTeaCmd)
	drinkCmd.AddCommand(coffeeCmd)
	drinkCmd.AddCommand(liquorCmd)
	drinkCmd.AddCommand(porridgeCmd)
}
