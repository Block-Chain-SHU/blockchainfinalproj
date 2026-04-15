package chaincode

/*
定义用户结构体
用户ID
用户类型
实名认证信息哈希,包括用户注册的姓名、身份证号、手机号、注册平台同意协议签名的哈希
追踪记录列表
*/
type User struct {
	UserID       string   `json:"userID"`
	UserType     string   `json:"userType"`
	RealInfoHash string   `json:"realInfoHash"`
	TraceList    []*TraceRecord `json:"traceList"`
}

/*
定义追踪记录结构体
学籍追踪码
学生输入
课程/实验室输入
项目实践/竞赛输入
就业/升学输入
*/
type TraceRecord struct {
	Student_trace_id string        `json:"student_trace_id"`
	Farmer_input      Farmer_input  `json:"farmer_input"`
	Factory_input     Factory_input `json:"factory_input"`
	Driver_input      Driver_input  `json:"driver_input"`
	Shop_input        Shop_input    `json:"shop_input"`
}

// HistoryQueryResult structure used for handling result of history query
type HistoryQueryResult struct {
	Record    *TraceRecord `json:"record"`
	TxId      string `json:"txId"`
	Timestamp string `json:"timestamp"`
	IsDelete  bool   `json:"isDelete"`
}

/*
学生
学籍追踪码起始角色（自动生成）
学生姓名
专业方向
入学时间
毕业时间
培养导师
*/
type Farmer_input struct {
	Fa_studentName   string `json:"fa_studentName"`
	Fa_major         string `json:"fa_major"`
	Fa_enrollTime    string `json:"fa_enrollTime"`
	Fa_graduationTime string `json:"fa_graduationTime"`
	Fa_mentorName    string `json:"fa_mentorName"`
	Fa_img         string `json:"fa_img"` // 图片链接
	Fa_Txid        string `json:"fa_txid"`
	Fa_Timestamp   string `json:"fa_timestamp"`
}

/*
课程/实验室
课程/实验室名称
课程批次
开课/实验时间
实验室/课程平台
联系方式
*/
type Factory_input struct {
	Fac_courseOrLabName string `json:"fac_courseOrLabName"`
	Fac_courseBatch     string `json:"fac_courseBatch"`
	Fac_courseTime      string `json:"fac_courseTime"`
	Fac_platformName    string `json:"fac_platformName"`
	Fac_contactNumber   string `json:"fac_contactNumber"`
	Fac_img             string `json:"fac_img"` // 图片链接
	Fac_Txid            string `json:"fac_txid"`
	Fac_Timestamp       string `json:"fac_timestamp"`
}

/*
项目实践/竞赛导师
导师姓名
指导年限
电话
项目/竞赛编号
项目贡献说明
*/
type Driver_input struct {
	Dr_mentorName      string `json:"dr_mentorName"`
	Dr_experienceYears string `json:"dr_experienceYears"`
	Dr_phone           string `json:"dr_phone"`
	Dr_projectCode     string `json:"dr_projectCode"`
	Dr_contribution    string `json:"dr_contribution"`
	Dr_img             string `json:"dr_img"` // 图片链接
	Dr_Txid            string `json:"dr_txid"`
	Dr_Timestamp       string `json:"dr_timestamp"`
}

/*
就业单位/研究生去向
签约/录取时间
入职/入学时间
单位/目标院校
单位/院校地址
联系人电话
*/
type Shop_input struct {
	Sh_offerTime    string `json:"sh_offerTime"`
	Sh_onboardTime  string `json:"sh_onboardTime"`
	Sh_orgName      string `json:"sh_orgName"`
	Sh_orgAddress   string `json:"sh_orgAddress"`
	Sh_contactPhone string `json:"sh_contactPhone"`
	Sh_img          string `json:"sh_img"` // 图片链接
	Sh_Txid         string `json:"sh_txid"`
	Sh_Timestamp    string `json:"sh_timestamp"`
}
