package mdparse

import (
	"encoding/base64"
	"testing"
)

func TestParse(t *testing.T) {
	text := "MeLine AMA 抽奖活动来啦🎉完成关注、转发推特即可参与抽奖，抽5个幸运儿，每人10Matic🔥🔥\n\n《首次使用Tristan bot的用户请注意：点击“Forward and Start”之后，会自动跳转到机器人私聊界面，点击“Start”会自动生成一个专属的Tristan内置钱包地址以及一个Pincode。这个Pincode类似于私钥，可以使用“/change_pin_code”指令更改初始Pincode。活动结束后，奖励会自动发到内置钱包地址🎉》\n\n📕参与抽奖条件(To participate):\n①关注推特(Follow twitter)\n②转发推特(RT)\n\n😎 Number of people awarded: 5 \n\n🏆 Reward:  50.0000 Matic\n\n📃 Minimum Eligible Score: 3\n\n⏰ End Time: 23 Nov 22 11:00 UTC"
	t.Log(base64.StdEncoding.EncodeToString([]byte(ParseV2(text))))
}
