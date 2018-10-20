package cmd

import (
	"flag"
	"log"

	"github.com/golang/glog"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var info bool

// configCmd represents the config command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "日志测试",
	Long:  `日志输出测试.`,
	Run: func(cmd *cobra.Command, args []string) {
		// logrus
		if info {
			logrus.SetLevel(logrus.InfoLevel)
		} else {
			logrus.SetLevel(logrus.DebugLevel)
		}
		logrus.Debug("调试")
		logrus.Info("信息")
		logrus.Infoln("信息", "其他", 123, true)
		logrus.Warn("警告")
		logrus.Error("错误")
		// log.Fatal("退出")
		// log.Panic("panic")

		// glog
		flag.Parse()
		defer glog.Flush()
		glog.Info("hello, glog")
		// glog.V(1)
		glog.Info("信息")
		glog.Warning("警告")
		glog.Errorln("错误11")

		// 内置log
		log.SetFlags(log.Lshortfile | log.LstdFlags)
		log.Println("信息")
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
	logCmd.Flags().BoolVarP(&info, "info", "i", false, "只显示Info以上信息")
}
