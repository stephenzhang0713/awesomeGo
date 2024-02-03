package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type GlobalConfig struct {
	AgentId             string      `json:"agentUuid"`
	AgentQueueId        string      `json:"agentQueueId"`
	CmdbUuid            string      `json:"cmdbUuid"`
	CfgUuid             string      `json:"cfgUuid"`
	LogPath             string      `json:"logPath"`
	LogSize             int         `json:"logSize"`
	AgentIp             string      `json:"localIp"`
	AgentName           string      `json:"agentName"`
	CompanyName         string      `json:"companyName"`
	DataServiceIp       string      `json:"dataServiceIp"`
	DataServicePort     int         `json:"dataServicePort"`
	DataPushMode        string      `json:"dataPushMode"`
	PortalPrefix        string      `json:"portalUrlPrefix"`
	MQAddr              string      `json:"mqAddr"`
	IsProxy             bool        `json:"isProxy"`
	ProxyPort           *ProxyPort  `json:"proxyPort"`
	PluginDownPort      string      `json:"pluginsPort"`
	DebugAgent          bool        `json:"debugAgent"`
	DebugScanner        bool        `json:"debugScanner"`
	DebugCollector      bool        `json:"debugCollector"`
	DebugCollectorList  []int64     `json:"debugCollectorList"`
	Is7xCli             bool        `json:"is7xCli"`
	SevenXCliParams     []string    `json:"sevenXCliParams"`
	AutoCollect         bool        `json:"autoCollect"`
	JudgeIpAgain        bool        `json:"judgeIpAgain"`
	IsOpenMQSSL         bool        `json:"isOpenMQSSL"`
	Http                *HttpConfig `json:"http"`
	DeviceSn            string      `json:"deviceSn"`
	Operator            string      `json:"operator"`
	Area                string      `json:"area"`
	Region              string      `json:"region"`
	StatusCheckAssignIP string      `json:"statusCheckAssignIP"` // 状态检测指定IP
	HttpPostTransferIP  string      `json:"httpPostTransferIP"`  // Transfer指定IP
	AgentVersion        string      `json:"agentVersion"`        //agent版本信息

	UFileConfig  *UFileConfig `json:"ufileConfig"`  //oss配置对象
	RegisterCode string       `json:"registerCode"` //注册码
}

type ProxyPort struct {
	Servicepory    string `json:"servicePort"`
	MQprot         string `json:"mqPort"`
	Opsprot        string `json:"opsPort"`
	PluginDownPort string `json:"pluginsDown"`
}

type HttpConfig struct {
	Enabled bool   `json:"enabled"`
	Listen  string `json:"listen"`
}

type UFileConfig struct {
	PublicKey       string `json:"public_key"`
	PrivateKey      string `json:"private_key"`
	BucketHost      string `json:"bucket_host"`
	BucketName      string `json:"bucket_name"`
	FileHost        string `json:"file_host"`
	VerifyUploadMD5 bool   `json:"verfiy_upload_md5"`
	Endpoint        string `json:"endpoint"`
}

var GlobalCfg *GlobalConfig
var IsUcpe string

func main() {
	// 获取当前工作目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	cfg := dir + "/02-code-project-organization/cfg.json"
	viper.SetConfigFile(cfg)
	// 尝试读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			// Config file not found; ignore error if desired
			log.Println("no such config file")
		} else {
			// Config file was found but another error was produced
			log.Println("read config error")
		}
	}

	// 如果配置文件存在，尝试将配置映射到GlobalCfg
	if err := viper.Unmarshal(&GlobalCfg); err != nil {
		log.Fatalf("ParseCfg Unmarshal Config Error: %v", err)
	}

	viper.SetDefault("judgeIpAgain", true)

	if IsUcpe != "" {
		viper.SetDefault("autoCollect", false)
		viper.SetDefault("deviceSn", IsUcpe)
	} else {
		viper.SetDefault("autoCollect", true)
	}

	GlobalCfgJSON, err := json.MarshalIndent(GlobalCfg, "", "  ")
	if err != nil {
		return
	}
	log.Println(string(GlobalCfgJSON))
}

func WriteBackCfg(filePath string) {
	log.Println("WriteBackCfg Start")
	// 使用 Viper 直接写入配置
	viper.SetConfigFile(filePath)

	GlobalCfg.DebugAgent = true
	GlobalCfg.AgentName = "test112"

	// 根据 global.GlobalCfg 更新配置值
	viper.Set("debugAgent", GlobalCfg.DebugAgent)
	viper.Set("agentName", GlobalCfg.AgentName)
	// ...设置 global.GlobalCfg 中其他需要保存的字段

	// 保存配置到文件
	if err := viper.WriteConfig(); err != nil {
		log.Fatalf("WriteBackCfg Write File Error: %v", err)
	}

	log.Println(" WriteBackCfg End And Success")
}
