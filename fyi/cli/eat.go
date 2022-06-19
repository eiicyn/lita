package cli

import (
	"github.com/kyokomi/emoji/v2"
	"github.com/spf13/cobra"

	"fyi/pkg"
)

const eat = "吃"

var eatCmd = &cobra.Command{
	Use:   "eat",
	Short: emoji.Sprintf(":face_savoring_food:吃"),
	Long:  emoji.Sprintf(`:face_savoring_food:吃！`),
}

var westFoodCmd = &cobra.Command{
	Use:   "west",
	Short: emoji.Sprintf(":cut_of_meat:西餐"),
	Long:  emoji.Sprintf(`:cut_of_meat:西餐！`),
	Run: func(cmd *cobra.Command, args []string) {
		pkg.SendSMS(eat + "西餐")
	},
}

var cakeCmd = &cobra.Command{
	Use:   "cake",
	Short: emoji.Sprintf(":shortcake:蛋糕"),
	Long:  emoji.Sprintf(`:shortcake:蛋糕！`),
	Run: func(cmd *cobra.Command, args []string) {
		pkg.SendSMS(eat + "蛋糕")
	},
}

var burgerCmd = &cobra.Command{
	Use:   "burger",
	Short: emoji.Sprintf(":hamburger:汉堡"),
	Long:  emoji.Sprintf(`:hamburger:汉堡！`),
	Run: func(cmd *cobra.Command, args []string) {
		pkg.SendSMS(eat + "汉堡")
	},
}

var pizzaCmd = &cobra.Command{
	Use:   "pizza",
	Short: emoji.Sprintf(":pizza:披萨"),
	Long:  emoji.Sprintf(`:pizza:披萨！`),
	Run: func(cmd *cobra.Command, args []string) {
		pkg.SendSMS(eat + "披萨")
	},
}

var lobsterCmd = &cobra.Command{
	Use:   "lobster",
	Short: emoji.Sprintf(":lobster:龙虾"),
	Long:  emoji.Sprintf(`:lobster:龙虾！`),
	Run: func(cmd *cobra.Command, args []string) {
		pkg.SendSMS(eat + "龙虾")
	},
}

var japanFoodCmd = &cobra.Command{
	Use:   "japan",
	Short: emoji.Sprintf(":sushi:日料"),
	Long:  emoji.Sprintf(`:sushi:日料！`),
	Run: func(cmd *cobra.Command, args []string) {
		pkg.SendSMS(eat + "日料")
	},
}

var hotpotCmd = &cobra.Command{
	Use:   "hotpot",
	Short: emoji.Sprintf(":pot_of_food:火锅"),
	Long:  emoji.Sprintf(`:pot_of_food:火锅！`),
	Run: func(cmd *cobra.Command, args []string) {
		pkg.SendSMS(eat + "火锅")
	},
}

var bbqCmd = &cobra.Command{
	Use:   "bbq",
	Short: emoji.Sprintf(":meat_on_bone:烧烤&烤肉"),
	Long:  emoji.Sprintf(`:meat_on_bone:烧烤&烤肉！`),
	Run: func(cmd *cobra.Command, args []string) {
		pkg.SendSMS(eat + "烧烤&烤肉")
	},
}

var sichuanFoodCmd = &cobra.Command{
	Use:   "sichuan",
	Short: emoji.Sprintf(":hot_face:川菜"),
	Long:  emoji.Sprintf(`:hot_pepper:川菜！`),
	Run: func(cmd *cobra.Command, args []string) {
		pkg.SendSMS(eat + "川菜")
	},
}

var northEastFoodCmd = &cobra.Command{
	Use:   "northeast",
	Short: emoji.Sprintf(":cold_face:东北菜"),
	Long:  emoji.Sprintf(`:snowflake:东北菜！`),
	Run: func(cmd *cobra.Command, args []string) {
		pkg.SendSMS(eat + "东北菜")
	},
}

func init() {
	eatCmd.AddCommand(westFoodCmd)
	eatCmd.AddCommand(burgerCmd)
	eatCmd.AddCommand(pizzaCmd)
	eatCmd.AddCommand(sichuanFoodCmd)
	eatCmd.AddCommand(northEastFoodCmd)
	eatCmd.AddCommand(bbqCmd)
	eatCmd.AddCommand(hotpotCmd)
	eatCmd.AddCommand(lobsterCmd)
	eatCmd.AddCommand(japanFoodCmd)
	eatCmd.AddCommand(cakeCmd)
}
