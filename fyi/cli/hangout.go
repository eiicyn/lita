package cli

import (
	"github.com/kyokomi/emoji/v2"
	"github.com/spf13/cobra"

	"fyi/pkg"
)

const hangout = "去"

var hangoutCmd = &cobra.Command{
	Use:   "hangout",
	Short: emoji.Sprintf(":crystal_ball:出去玩"),
	Long:  emoji.Sprintf(`:crystal_ball:出去玩！`),
}

var parkCmd = &cobra.Command{
	Use:   "park",
	Short: emoji.Sprintf(":fountain:公园"),
	Long:  emoji.Sprintf(`:camping:公园！`),
	Run: func(cmd *cobra.Command, args []string) {
		pkg.SendSMS(hangout + "公园")
	},
}

var walkCmd = &cobra.Command{
	Use:   "walk",
	Short: emoji.Sprintf(":person_walking:散步"),
	Long:  emoji.Sprintf(`:person_walking:散步！`),
	Run: func(cmd *cobra.Command, args []string) {
		pkg.SendSMS(hangout + "散步")
	},
}

var supermarketCmd = &cobra.Command{
	Use:   "supermarket",
	Short: emoji.Sprintf(":shopping_cart:逛超市"),
	Long:  emoji.Sprintf(`:shopping_cart:逛超市！`),
	Run: func(cmd *cobra.Command, args []string) {
		pkg.SendSMS(hangout + "逛超市")
	},
}

var shoppingCmd = &cobra.Command{
	Use:   "shopping",
	Short: emoji.Sprintf(":lipstick:上街购物"),
	Long:  emoji.Sprintf(`:shopping_bags:上街购物！`),
	Run: func(cmd *cobra.Command, args []string) {
		pkg.SendSMS(hangout + "上街购物")
	},
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: emoji.Sprintf(":artist_palette:看展"),
	Long:  emoji.Sprintf(`:artist_palette:看展！`),
	Run: func(cmd *cobra.Command, args []string) {
		pkg.SendSMS(hangout + "看展")
	},
}

var movieCmd = &cobra.Command{
	Use:   "movie",
	Short: emoji.Sprintf(":movie_camera:看电影"),
	Long:  emoji.Sprintf(`:film_frames:看电影！`),
	Run: func(cmd *cobra.Command, args []string) {
		pkg.SendSMS(hangout + "看电影")
	},
}

var amusementParkCmd = &cobra.Command{
	Use:   "fairytown",
	Short: emoji.Sprintf(":ferris_wheel:游乐园"),
	Long:  emoji.Sprintf(`:ferris_wheel:游乐园！`),
	Run: func(cmd *cobra.Command, args []string) {
		pkg.SendSMS(hangout + "游乐园")
	},
}

var singCmd = &cobra.Command{
	Use:   "sing",
	Short: emoji.Sprintf(":microphone:唱歌"),
	Long:  emoji.Sprintf(`:microphone:唱歌！`),
	Run: func(cmd *cobra.Command, args []string) {
		pkg.SendSMS(hangout + "唱歌")
	},
}

var mountainCmd = &cobra.Command{
	Use:   "mountain",
	Short: emoji.Sprintf(":mountain_cableway:爬山"),
	Long:  emoji.Sprintf(`:snow_capped_mountain:爬山！`),
	Run: func(cmd *cobra.Command, args []string) {
		pkg.SendSMS(hangout + "爬山")
	},
}

func init() {
	hangoutCmd.AddCommand(singCmd)
	hangoutCmd.AddCommand(movieCmd)
	hangoutCmd.AddCommand(showCmd)
	hangoutCmd.AddCommand(shoppingCmd)
	hangoutCmd.AddCommand(supermarketCmd)
	hangoutCmd.AddCommand(walkCmd)
	hangoutCmd.AddCommand(parkCmd)
	hangoutCmd.AddCommand(mountainCmd)
	hangoutCmd.AddCommand(amusementParkCmd)
}
