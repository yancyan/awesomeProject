package log

import (
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

const logPath = "./logs"

var Log = logrus.New()

//
//func initLogger() {
//	log.SetFormatter(&log.JSONFormatter{})
//	log.SetLevel(log.InfoLevel)
//	log.SetReportCaller(true)
//
//	writer2 := os.Stdout
//	writer3, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0755)
//	if err != nil {
//		log.Fatalf("create file log.txt failed: %v", err)
//	}
//	log.SetOutput(io.MultiWriter(writer2, writer3))
//
//}

func init() {
	//InitLog("")
}

func TestLogger() {

	//InitLog("auto-change-")

	for i := 0; i < 10; i++ {
		Log.Infof("abc %d", i)
		//time.Sleep(time.Second * 1)
	}
}

func InitLog(filePrefix string) {
	Log.Out = os.Stdout

	level, err := logrus.ParseLevel("info")
	if err != nil {
		Log.Panicf("设置log级别失败：%v", err)
	}
	Log.SetLevel(level)

	configLogger(Log, logPath, filePrefix, 30)
}

/**
  文件日志
*/
func configLogger(log *logrus.Logger, logPath string, filePrefix string, save uint) {

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer(logPath, filePrefix, "debug", save), // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer(logPath, filePrefix, "info", save),
		logrus.WarnLevel:  writer(logPath, filePrefix, "warn", save),
		logrus.ErrorLevel: writer(logPath, filePrefix, "error", save),
		logrus.FatalLevel: writer(logPath, filePrefix, "fatal", save),
		logrus.PanicLevel: writer(logPath, filePrefix, "panic", save),
	}, &logrus.JSONFormatter{})

	log.AddHook(lfHook)
}

func writer(logPath string, filePrefix string, level string, save uint) *rotatelogs.RotateLogs {
	logFullPath := path.Join(logPath, filePrefix+level)
	//var fileSuffix = time.Now().In(time.UTC).Format("20060402") + FileSuffix
	logier, err := rotatelogs.New(
		logFullPath+"-%Y%m%d%H"+".log",
		rotatelogs.WithLinkName(logFullPath), // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(-1),
		rotatelogs.WithRotationCount(int(save)),   // 文件最大保存份数
		rotatelogs.WithRotationTime(time.Hour*24), // 日志切割时间间隔
	)
	if err != nil {
		panic(err)
	}
	return logier
}
