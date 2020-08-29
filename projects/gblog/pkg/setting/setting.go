package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret       string
	PageSize        int
	RuntimeRootPath string

	PrefixUrl      string
	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string

	ExportSavePath string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

var DbSetting = &Database{}

func Setup() {
	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalln("Fail to parse app.ini: ", err)
	}

	err = cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatalln("Fail to parse app section: ", err)
	}
	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024

	err = cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalln("Fail to parse server section: ", err)
	}
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

	err = cfg.Section("database").MapTo(DbSetting)
	if err != nil {
		log.Fatalln("Fail to parse app section: ", err)
	}

	err = cfg.Section("redis").MapTo(RedisSetting)
	if err != nil {
		log.Fatalln("Fail to parse app section: ", err)
	}
}
