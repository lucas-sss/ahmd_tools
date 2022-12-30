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

type BindInfo struct {
	Id              int64    `json:"tokenid" gorm:"column:id;type:bigint(20) NOT NULL;primaryKey"`
	Type            int      `json:"type" gorm:"column:type;type:tinyint(1) NOT NULL"`
	CipherFlag      int      `json:"cipherFlag" gorm:"column:cipher_flag;type:tinyint(1) NOT NULL DEFAULT 0"`
	ThirdId         string   `json:"thirdId" gorm:"column:third_id;type:varchar(50) DEFAULT NULL"`
	KeyId           string   `json:"keyId" gorm:"column:key_id;type:varchar(50) NOT NULL"`
	NewKeyId        string   `json:"newKeyId" gorm:"column:new_key_id;type:varchar(50) DEFAULT NULL"`
	PrepKeyId       string   `json:"prepKeyId" gorm:"column:prep_key_id;type:varchar(50) DEFAULT NULL"`
	BelongCorp      string   `json:"belongCorp" gorm:"column:belong_corp;type:varchar(50) DEFAULT NULL"`
	BoxCode         string   `json:"boxCode" gorm:"column:box_code;type:varchar(50) NOT NULL"`
	SlaveBoxCode    string   `json:"slaveBoxCode" gorm:"column:slave_box_code;type:varchar(50) NOT NULL"`
	TitularBoxCode  string   `json:"titularBoxCode" gorm:"column:titular_box_code;type:varchar(50) DEFAULT NULL"`
	RouterMac       string   `json:"routerMac" gorm:"column:router_mac;type:varchar(255) DEFAULT NULL"`
	MacAddress      string   `json:"macAddress" gorm:"column:mac_address;type:varchar(50) NOT NULL"`
	SlaveMacAddress string   `json:"slaveMacAddress" gorm:"column:slave_mac_address;type:varchar(50) NOT NULL"`
	SoftVersion     string   `json:"softVersion" gorm:"column:soft_version;type:varchar(20) NOT NULL"`
	Ip              string   `json:"ip" gorm:"column:ip;type:varchar(20) NOT NULL"`
	State           int      `json:"state" gorm:"column:state;type:tinyint(1) NOT NULL DEFAULT 3"`
	Reason          string   `json:"reason" gorm:"column:reason;type:varchar(50) DEFAULT NULL"`
	Link            bool     `json:"link" gorm:"column:link;type:tinyint(1) NOT NULL DEFAULT 1"`
	LinkType        int      `json:"linkType" gorm:"column:link_type;type:tinyint(1) NOT NULL DEFAULT 1"`
	Bind            bool     `json:"bind" gorm:"column:bind;type:tinyint(1) NOT NULL DEFAULT 0"`
	CreateTime      JsonTime `json:"createTime" gorm:"column:createdate;type:datetime DEFAULT NULL"`
	UpdateTime      JsonTime `json:"updateTime" gorm:"column:createdate;type:datetime DEFAULT NULL"`
}

func (BindInfo) TableName() string {
	return "t_bindinfo"
}
