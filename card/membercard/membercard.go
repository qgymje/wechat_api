package membercard

var url = "https://api.weixin.qq.com/card/membercard/activate?access_token=%s"

type MemberCardActiveResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type MemberCardActiveRequest struct {
	MembershipNumber      string `json:"membership_number" required`
	Code                  string `json:"code" required`
	CardID                string `json:"card_id" required`   // 自定义code卡必填
	BackgroundPicUrl      string `json:"background_pic_url"` // 商家自定义会员卡背景图片, 必须为已经通过上传图片接口后的地址
	ActiveBeginTime       int    `json:"activate_begin_time"`
	ActiveEndTime         int    `json:"activate_end_time"`
	InitBonus             int    `json:"init_bonus"`               // 初始积分, defulat 0
	InitBonusRecord       string `json:"init_bonus_record"`        //积分同步说明
	InitBalance           int    `json:"init_balance"`             // 初始余额, default 0
	InitCustomFieldValue1 string `json:"init_custom_field_value1"` // 4个汉字,12个字节
	InitCustomFieldValue2 string `json:"init_custom_field_value2"`
	InitCustomFieldValue3 string `json:"init_custom_field_value3"`
}
