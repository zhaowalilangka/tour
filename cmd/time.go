package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/zhaowalilangka/tour/internal/timer"
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("输出结果是： %s, %d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
	},
}
