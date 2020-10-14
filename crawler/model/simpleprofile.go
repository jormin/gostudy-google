package model

// 简易信息
type SimpleProfile struct {
	ID          int    `json:"id" remark:"ID"`
	Avatar      string `json:"avatar" remark:"头像"`
	Nickname    string `json:"nickname" remark:"昵称"`
	City        string `json:"city" remark:"城市"`
	Description string `json:"description" remark:"内心独白"`
	Gender      int    `json:"gender" remark:"性别"`
	Age         string `json:"age" remark:"年龄"`
	Education   string `json:"education" remark:"学历"`
	Marital     string `json:"marital" remark:"婚姻"`
	Height      string `json:"height" remark:"身高"`
	Salary      string `json:"salary" remark:"月薪"`
}
