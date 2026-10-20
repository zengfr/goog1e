package goutil

import (
	"fmt"

	"github.com/sirupsen/logrus"
	//"github.com/sirupsen/logrus/hooks"
)

func Logrusmain() {
	//hook := new(logrus.Hook)
	//logrus.AddHook(*hook)
	logrus.SetLevel(logrus.InfoLevel) // 在测试环境中设置低等级级别，
	//logrus.SetLevel(logrus.InfoLevel) // 在生产环境中需要考虑性能，关注关键信息，level 设置高一点
	//logrus.SetReportCaller(true)          // 调用者文件名与位置
	logrus.SetFormatter(new(logrus.TextFormatter)) // 日志格式设置成json
	logrus.RegisterExitHandler(func() {
		fmt.Println(".")
	})
	//logrus.SetOutput(io.Discard)
	logrus.Traceln("trace")
	logrus.Debugln("debug")
	logrus.Infoln("Info")
	logrus.Warnln("warn")
	logrus.Errorln("error")
	//logrus.Fatalln("fatal")
	//logrus.Panicln("panic")

	/*logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})*/
}
