package model

// 信息
type Profile struct {
	BasicInfo  BasicInfo  `json:"basic_info" remark:"基础信息"`
	NormalInfo NormalInfo `json:"normal_info" remark:"常规信息"`
	Condition  NormalInfo `json:"condition" remark:"择偶条件"`
}

// 基础信息
type BasicInfo struct {
	ID          int    `json:"id" remark:"ID"`
	Avatar      string `json:"avatar" remark:"头像"`
	Nickname    string `json:"nickname" remark:"昵称"`
	City        string `json:"city" remark:"城市"`
	Description string `json:"description" remark:"内心独白"`
}

// 常规信息
type NormalInfo struct {
	Gender        int    `json:"gender" remark:"性别"`
	Age           string `json:"age" remark:"年龄"`
	Education     string `json:"education" remark:"学历"`
	Job           string `json:"job" remark:"工作"`
	Marital       string `json:"marital" remark:"婚姻"`
	Height        string `json:"height" remark:"身高"`
	Weight        string `json:"weight" remark:"体重"`
	Salary        string `json:"salary" remark:"月薪"`
	WorkPlace     string `json:"work_place" remark:"工作地点"`
	NativePlace   string `json:"native_place" remark:"籍贯"`
	Stature       string `json:"stature" remark:"体型"`
	WhenToMarry   string `json:"when_to_marry" remark:"何时结婚"`
	Constellation string `json:"constellation" remark:"星座"`
	Nation        string `json:"nation" remark:"民族"`
	HasBaby       int    `json:"has_baby" remark:"是否有小孩"`
	IsWantBaby    int    `json:"is_want_baby" remark:"是否想要孩子"`
	IsDrink       int    `json:"is_drink" remark:"是否喝酒"`
	IsSmoke       int    `json:"is_smoke" remark:"是否抽烟"`
	HasCar        int    `json:"has_car" remark:"是否有车"`
	HasHouse      int    `json:"has_house" remark:"是否有房"`
}

// 生成新的信息
func NewProfile() Profile {
	return Profile{
		BasicInfo:  BasicInfo{},
		NormalInfo: NewNormalInfo(),
		Condition:  NewNormalInfo(),
	}
}

// 生成新的基础信息
func NewNormalInfo() NormalInfo {
	normalInfo := NormalInfo{}
	normalInfo.HasBaby = 2
	normalInfo.IsWantBaby = 2
	normalInfo.IsDrink = 2
	normalInfo.IsSmoke = 2
	normalInfo.HasCar = 2
	normalInfo.HasHouse = 2
	return normalInfo
}
