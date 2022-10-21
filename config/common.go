package config

import "os"

func EnvIsDev() bool {
	return os.Getenv("ENV") == "dev"
}

func IsTestNet() bool {
	return os.Getenv("IS_TEST_NET") == "1"
}

func UseWebsocket() bool {
	return os.Getenv("TG_WEBSOCKET") == "1"
}
func UseConsoleWrite() bool {
	return os.Getenv("LOG_CONSOLE_WRITE") == "1"
}

func GetAppId() string {
	if appId := os.Getenv("APP_ID"); appId != "" {
		return appId
	}
	return "TestAppId"
}

func ClearDiscordCommands() bool {
	return os.Getenv("CLEAR_DISCORD_COMMANDS") == "1"
}

func GetDingDingToken() string {
	return os.Getenv("DINGDING_TOKEN")
}

func GetDingDingAppKeyAndSecret() (string, string) {
	return os.Getenv("DINGDING_APP_KEY"), os.Getenv("DINGDING_SECRET")
}

func RedisScheduleDataKey(id string) string {
	return "job_data_" + id
}

func RedisScheduleLockKey(id string) string {
	return "job_lock_" + id
}

func RedisScheduleJobKey(ServiceName string) string {
	return "job_schedule_" + ServiceName
}

func GetDingConfig() (string, string) {
	return os.Getenv("DING_TOKEN"), os.Getenv("DING_SECRET")
}
