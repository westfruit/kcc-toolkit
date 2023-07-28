package conf

import (
	"fmt"
	"gitee.com/westfruit/kcc-toolkit/fileop"
	"io"
	"os"
	"os/exec"
	"path"

	"github.com/fsnotify/fsnotify"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {

	// 日志配置
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
	//log.SetReportCaller(true)

	// 创建日志文件夹
	if err := os.MkdirAll("log", os.ModePerm); err != nil {
		fmt.Errorf("创建日志文件夹错误: %s", err.Error())
	}

	//rl, _ := rotatelogs.New("adminlog/app.%Y%m%d%H%M.adminlog")
	rl, _ := rotatelogs.New("log/app.%Y%m%d.log") // 每天一个日志文件
	mw := io.MultiWriter(os.Stdout, rl)
	logrus.SetOutput(mw)

	dir, _ := os.Getwd()
	name, _ := os.Hostname()

	logrus.Println("当前路径:", dir)
	logrus.Println("主机名:", name)

	file, _ := exec.LookPath(os.Args[0])
	logrus.Println("exec path:", file)

	/*
		Viper配置键不区分大小写, Viper读取配置顺序（优先级从高到低）:
		1. explicit call to Set   //显式设置
		2. flag					  //命令行
		3. env                    //环境变量
		4. config                 //配置文件
		5. key/value store        //远程K/V存储
		6. default                //默认值
	*/

	//1、设置命令行
	if !pflag.CommandLine.HasFlags() {
		pflag.String("confpath", "conf", "path to look for the config file in")
		pflag.String("confname", "app", "name of config file (without extension)")
		pflag.String("conftype", "yml", "type of config file")
		pflag.String("mode", "", "run mode")
		pflag.String("prefix", "APP", "env var prefix")
		pflag.Parse() //解析
		viper.BindPFlags(pflag.CommandLine)
	}

	confPath := viper.GetString("confpath")
	confName := viper.GetString("confname")
	confType := viper.GetString("conftype")
	mode := viper.GetString("mode")
	prefix := viper.GetString("prefix")

	if ok := fileop.Exists(path.Join(dir, confPath)); !ok {
		confPath = path.Join("..", confPath)
	}

	logrus.Println("命令行--配置文件查找路径:", confPath)
	logrus.Println("命令行--设置配置文件名称(不包含后缀):", confName)
	logrus.Println("命令行--配置文件类型:", confType)
	logrus.Println("命令行--运行模式:", mode)
	logrus.Println("命令行--环境变量前缀:", prefix)

	//2、设置环境变量
	viper.AutomaticEnv()       //检查所有环境变量
	viper.SetEnvPrefix(prefix) //设置环境变量前缀

	//3、读取app.yml配置文件
	viper.AddConfigPath(confPath) //添加配置文件查找路径
	viper.AddConfigPath(".")      // 添加当前路径
	viper.SetConfigName(confName) //设置配置文件名称(不包含后缀)
	viper.SetConfigType(confType) //设置配置文件类型
	err := viper.ReadInConfig()   //读取配置文件
	if err != nil {
		logrus.Fatal("Fatal error config file", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {

		logrus.Info("Config file changed:", e.Name)
		logrus.Info("app name ", viper.GetString("name"))
		logrus.Info("snowflake.machineId = ", viper.GetInt("snowflake.machineId"))

	})

	//4、设置默认值
	//viper.SetDefault("mode", "prod")
	viper.SetDefault("db.maxIdleConn", 10)
	viper.SetDefault("db.maxOpenConn", 100)

	//5、读取其它配置文件
	mode = viper.GetString("mode")

	logrus.Info("最终运行模式 = ", mode)

	if len(mode) > 0 {
		viper.SetConfigName(confName + "." + mode)
		err = viper.MergeInConfig()
		if err != nil {
			logrus.Fatal("Fatal error config file", err)
		}
	}

	// 配置监听，在修改配置时候，会自动检查并做对应的回调事件处理
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {

		logrus.Info("Config file changed:", e.Name)
		logrus.Info("app name ", viper.GetString("name"))
		logrus.Info("snowflake.machineId = ", viper.GetInt("snowflake.machineId"))

	})
}
