package model

type ThridObject struct {
	ObjectId      string   `json:"objectid" gorm:"column:objectid;type:varchar(100) NOT NULL;primaryKey"`
	AppKey        string   `json:"appkey" gorm:"column:appkey;type:varchar(100) DEFAULT NULL"`
	ThridObjectId string   `json:"thridobjectid" gorm:"column:thridobjectid;type:varchar(100) DEFAULT NULL"`
	ObjectName    string   `json:"objectname" gorm:"column:objectname;type:varchar(200) DEFAULT NULL"`
	CreateDate    JsonTime `json:"createdate" gorm:"column:createdate;type:datetime DEFAULT NULL"`
	ObjectStatus  int      `json:"objectstatus" gorm:"column:objectstatus;type:int(11) DEFAULT NULL"`
	ObjectPrivate int      `json:"objectprivate" gorm:"column:objectprivate;type:int(11) DEFAULT NULL"`
	ObjectServer  string   `json:"objectserver" gorm:"column:objectserver;type:varchar(500) DEFAULT NULL"`
	ManageInfo    string   `json:"manageinfo" gorm:"column:manageinfo;type:varchar(2000) DEFAULT ''"`
	PermanentCode string   `json:"permanentcode" gorm:"column:permanentcode;type:varchar(200) DEFAULT NULL"`
	OrderId       string   `json:"orderid" gorm:"column:orderid;type:varchar(100) DEFAULT NULL"`
	ExpiredDate   string   `json:"expiredate" gorm:"column:expiredate;type:varchar(255) DEFAULT NULL"`
	ExtattrShort  string   `json:"extattrshort" gorm:"column:extattrshort;type:varchar(200) DEFAULT NULL"`
	ExtattrLarge  string   `json:"extattrlarge" gorm:"column:extattrlarge;type:text"`
}

func (ThridObject) TableName() string {
	return "t_objectlist"
}

type DingDingToken struct {
	TokenId    string   `json:"tokenid" gorm:"column:tokenid;type:varchar(100) NOT NULL;primaryKey"`
	TokenValue string   `json:"tokenvalue" gorm:"column:tokenvalue;type:varchar(200) DEFAULT NULL"`
	CreateDate JsonTime `json:"createdate" gorm:"column:createdate;type:datetime DEFAULT NULL"`
}

func (DingDingToken) TableName() string {
	return "t_dingding_token"
}
