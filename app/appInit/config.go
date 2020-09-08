package appInit

import (
	"github.com/spf13/viper"
)

var Config config

type config struct {
	Http struct {
		Domain string `mapstructure:"domain"`
		Port   string `mapstructure:"port"`
	} `mapstructure:"http"`
	Jwt struct {
		Secret     string `mapstructure:"secret"`
		ExpireTime int64  `mapstructure:"expire_time"`
		Issuer     string `mapstructure:"issuer"`
	} `mapstructure:"jwt"`
	Admin struct {
		SuperAdmin     string `mapstructure:"super_admin"`
		SuperAdminPass string `mapstructure:"super_admin_pass"`
		InitAdminPass  string `mapstructure:"init_admin_pass"`
		InitAdminAvt   string `mapstructure:"init_admin_avt"`
	} `mapstructure:"admin"`
	AdminJwt struct {
		Secret     string `mapstructure:"secret"`
		ExpireTime int64  `mapstructure:"expire_time"`
		Issuer     string `mapstructure:"issuer"`
	} `mapstructure:"adminjwt"`
	Captcha struct {
		ImgHeight int `mapstructure:"img_height"`
		ImgWidth  int `mapstructure:"img_width"`
		KeyLong   int `mapstructure:"key_long"`
	} `mapstructure:"captcha"`
	System struct {
		Env           string `mapstructure:"env"`
		UniversalPass string `mapstructure:"universal_pass"`
	} `mapstructure:"System"`
	DB struct {
		Host               string `mapstructure:"host"`
		Port               string `mapstructure:"port"`
		User               string `mapstructure:"user"`
		Password           string `mapstructure:"password"`
		Name               string `mapstructure:"name"`
		MaxIdleConnections int    `mapstructure:"max_idle_connections"`
		MaxOpenConnections int    `mapstructure:"max_idle_connections"`
	} `mapstructure:"db"`
	Push struct {
		Url                  string `mapstructure:"url"`
		AccessKeyId          string `mapstructure:"access_key_id"`
		AccessKeySecret      string `mapstructure:"access_key_secret"`
		AppDesignAndroidKey  string `mapstructure:"app_design_android_key"`
		AppFactoryAndroidKey string `mapstructure:"app_factory_android_key"`
		AppDesignIOSKey      string `mapstructure:"app_design_ios_key"`
		AppFactoryIOSKey     string `mapstructure:"app_factory_ios_key"`
	} `mapstructure:"push"`
	SendCloud struct {
		SmsKey  string `mapstructure:"sms_key"`
		SmsUser string `mapstructure:"sms_user"`
		MsgType string `mapstructure:"msg_type"`
		SmsUrl  string `mapstructure:"sms_url"`
		ApiKey  string `mapstructure:"api_key"`
		ApiUser string `mapstructure:"api_user"`
		MailUrl string `mapstructure:"mail_url"`
	} `mapstructure:"sendcloud"`
	Redis struct {
		Addr     string `mapstructure:"addr"`
		Password string `mapstructure:"password"`
		PoolSize int    `mapstructure:"pool_size"`
		DB       int    `mapstructure:"db"`
	} `mapstructure:"redis"`
	Elastic struct {
		Addr string `mapstructure:"addr"`
	} `mapstructure:"elastic"`
	Qiniu struct {
		AccessKey   string `mapstructure:"access_key"`
		SecretKey   string `mapstructure:"secret_key"`
		Bucket      string `mapstructure:"bucket"`
		Domain      string `mapstructure:"domain"`
		FileDomain  string `mapstructure:"file_domain"`
		CallbackURL string `mapstructure:"callback_url"`
	} `mapstructure:"qiniu"`
	Wx struct {
		AccessToken string `mapstructure:"access_token"`
	} `mapstructure:"wx"`
	AliPay struct {
		IsProd        bool   `mapstructure:"is_prod"`
		AppID         string `mapstructure:"app_id"`
		AppPrivateKey string `mapstructure:"app_private_key"`
		AliPublicKey  string `mapstructure:"ali_public_key"`
		NotifyUrl     string `mapstructure:"notify_url"`
		ReturnUrl     string `mapstructure:"return_url"`
	} `mapstructure:"alipay"`
	WeChatPay struct {
		IsProd    bool   `mapstructure:"is_prod"`
		AppID     string `mapstructure:"app_id"`
		MchID     string `mapstructure:"mch_id"`
		ApiKey    string `mapstructure:"api_key"`
		NotifyUrl string `mapstructure:"notify_url"`
		ReturnUrl string `mapstructure:"return_url"`
	} `mapstructure:"wechatpay"`
	Kuaidi100 struct {
		Key                         string `mapstructure:"key"`
		Customer                    string `mapstructure:"customer"`
		AutoNumberRequestUrl        string `mapstructure:"auto_number_request_url"`
		SubscribeRequestUrl         string `mapstructure:"subscribe_request_url"`
		SubscribeCallbackRequestUrl string `mapstructure:"subscribe_callback_request_url"`
	} `mapstructure:"kuaidi100"`
}

func InitConfig() {
	//local
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yml")
	viper.ReadInConfig()
	//viper.ReadConfig(bytes.NewBufferString(remoteConfig))
	if err := viper.Unmarshal(&Config); err != nil {
		panic(err)
	}
}
