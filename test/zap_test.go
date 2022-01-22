/**
 * @Author Mr.LiuQH
 * @Description zap 使用
 * @Date 2021/6/18 10:52 上午
 **/
package test

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"testing"
)

// 创建日志记录器
func TestCreateLogger(t *testing.T) {
	// 初始化logger
	logger := zap.NewExample()
	// 使用defer logger.Sync()将缓存同步到文件中。
	defer logger.Sync()
	// 记录日志
	logger.Info("NewExample",
		zap.String("name", "张三"),
		zap.Int("age", 18),
	)
	productionLogger, _ := zap.NewProduction()
	defer productionLogger.Sync()
	productionLogger.Info("NewProduction",
		zap.String("name", "张三"),
		zap.Int("age", 18),
	)
	devLogger, _ := zap.NewDevelopment()
	defer devLogger.Sync()
	devLogger.Info("NewDevelopment",
		zap.String("name", "张三"),
		zap.Int("age", 18),
	)
}

// 使用默认记录日志
func TestRecordLogWithDefault(t *testing.T) {
	// 初始化记录器（使用默认记录器）
	logger := zap.NewExample()
	defer logger.Sync()
	// 记录日志
	logger.Debug("这是debug日志")
	logger.Debug("这是debug日志", zap.String("name", "张三"))
	logger.Info("这是info日志", zap.Int("age", 18))
	logger.Error("这是error日志", zap.Int("line", 130), zap.Error(fmt.Errorf("错误示例")))
	logger.Warn("这是Warn日志")
	// 下面两个都会中断程序
	//logger.Fatal("这是Fatal日志")
	//logger.Panic("这是Panic日志")
}

// 使用Sugar记录器
func TestRecordLogWithSuage(t *testing.T) {
	// 初始化记录器
	logger := zap.NewExample()
	// 把日志记录器转成Sugar
	sugarLogger := logger.Sugar()
	defer sugarLogger.Sync()
	// 记录日志
	sugarLogger.Debug("这是debug日志 ", zap.String("name", "张三"))
	sugarLogger.Debugf("这是Debugf日志 name:%s ", "张三")
	sugarLogger.Info("这是info日志", zap.Int("age", 18))
	sugarLogger.Infof("这是Infof日志  内容:%v", map[string]string{"爱好": "动漫"})
	sugarLogger.Error("这是error日志", zap.Int("line", 130), zap.Error(fmt.Errorf("错误示例")))
	sugarLogger.Errorf("这是Errorf日志，错误信息：%s", "错误报告！")
	sugarLogger.Warn("这是Warn日志")
	sugarLogger.Warnf("这是Warnf日志 %v", []int{1, 2, 4, 5})
	// 下面两个都会中断程序
	//sugarLogger.Fatal("这是Fatal日志")
	//sugarLogger.Panic("这是Panic日志")
}

func Test2File(t *testing.T) {
	// 指定写入文件
	fileHandle, _ := os.Create("./test.log")
	writeFile := zapcore.AddSync(fileHandle)
	// 设置日志输出格式为JSON (参数复用NewDevelopmentEncoderConfig)
	encoder := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())
	// 返回zapcore.Core
	zcore := zapcore.NewCore(encoder, zapcore.Lock(writeFile), zap.DebugLevel)
	// 创建日志记录器
	logger := zap.New(zcore)
	defer logger.Sync()
	// 记录日志
	logger.Info("输出日志到文件", zap.String("name", "张三"))
}

// 同时输入到文件和控制台
func TestPrintFileAndStd(t *testing.T) {
	// 指定写入文件
	fileHandle, _ := os.OpenFile("./test.log",os.O_CREATE | os.O_RDWR | os.O_APPEND,os.ModePerm)
	// 同时写入文件和控制台
	writeFile := zapcore.NewMultiWriteSyncer(fileHandle,os.Stdout)
	// 设置日志输出格式为JSON (参数复用NewDevelopmentEncoderConfig)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	// 返回zapcore.Core
	zcore := zapcore.NewCore(encoder, zapcore.Lock(writeFile), zap.DebugLevel)
	// 创建日志记录器
	logger := zap.New(zcore)
	sugarLogger := logger.Sugar()
	defer sugarLogger.Sync()

	// 记录日志
	sugarLogger.Infof("请求 IP:%s Method:%s Uri:%s","127.0.0.1","GET","/index")
}

// 获取文件切割和归档配置信息
func getLumberjackConfig() zapcore.WriteSyncer {
	lumberjackLogger := &lumberjack.Logger{
		Filename: "./zap.log",//日志文件
		MaxSize: 1,//单文件最大容量(单位MB)
		MaxBackups: 3,//保留旧文件的最大数量
		MaxAge: 1,// 旧文件最多保存几天
		Compress: false, // 是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberjackLogger)
}

// 测试日志切割和归档
func TestCutAndArchive(t *testing.T) {
	// 设置日志输出格式为JSON (参数复用NewDevelopmentEncoderConfig)
	encoder := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())
	core := zapcore.NewCore(encoder, getLumberjackConfig(), zap.DebugLevel)
	sugarLogger := zap.New(core).Sugar()
	defer sugarLogger.Sync()
	// 记录日志
	sugarLogger.Infof("日志内容:%s",strings.Repeat("日志",60000))
}