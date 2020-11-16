package errorUtil

import (
	"fmt"
)

const HERROR_PREFIX = "HERROR:"

//ZError的基本code，如果service相关的code，请service自行定义
const (
	ERROR_CODE_PARAMETER_FORMAT_INVALID = 400
	ERROR_CODE_ACCESS_TOKEN_TIMEOUT     = 401
	ERROR_CODE_NOT_FOUND                = 404
	ERROR_CODE_PARAMETER_AUTH_INVALID   = 406
	ERROR_INTERNAL_SERVER_ERROR         = 500
	ERROR_CODE_ACCESS_TOKEN_ERROR       = 1001
	ERROR_CODE_MOBILE_IS_REGISTERED     = 1002
	ERROR_CODE_MOBILE_IS_NOT_EXISTED    = 1003
	ERROR_CODE_CAPTCHA_TYPE_ERROR       = 1004
	ERROR_CODE_USER_NOT_FOUND           = 1005
	ERROR_CODE_MOBILE_IS_INVALID        = 1006
	ERROR_CODE_REQUEST_SMS_OVERLOAD     = 1007
	ERROR_CODE_CAPTCHA_TOKEN_ERROR      = 1008
	ERROR_CODE_REDIS_KEY_NULL           = 1015
	ERROR_CODE_REDIS_VALUE_NULL_PTR     = 1016
	ERROR_CODE_REDIS_VALUE_NULL         = 1017
	ERROR_CODE_REDIS_KEY_NOT_EXIST      = 1018

	ERROR_CODE_ALIOSS_CONFIG_IS_EMPTY          = 1200
	ERROR_CODE_UPLOAD_FILE_CONTENT_IS_EMPTY    = 1201
	ERROR_CODE_GET_BASE_URL_ERROR              = 1202
	ERROR_CODE_GET_FILE_BY_URL_ERROR           = 1203
	ERROR_CODE_GET_DEPARTMENT_ROLE_ERROR       = 1204
	ERROR_CODE_DEFAULT_ROLE_ERROR              = 1205
	ERROR_CODE_DEFAULT_RULE_ERROR              = 1206
	ERROR_CODE_ROLE_NOT_EXSIT_ERROR            = 1207
	ERROR_CODE_RULE_NOT_EXSIT_ERROR            = 1208
	ERROR_CODE_DEPARTMENT_ROLE_NOT_EXSIT_ERROR = 1209
	ERROR_CODE_RPC_ADDRESS_NOT_EXSIT_ERROR     = 1210
	ERROR_CODE_MULTI_DEPARTMENT_IDS_ERROR      = 1211

	ERROR_CODE_KAFKA_PRODUCE_NIL = 1500 //kafka相关
	ERROR_CODE_KAFKA_TOPIC_EMPTY = 1501 //topic为空
	ERROR_CODE_KAFKA_PUSH_FAIL   = 1502 //push失败

	ERROR_CODE_LOCATION_NOT_EXSIT          = 1503
	ERROR_CODE_INVENTORY_INSUFFICIENCY     = 1504
	ERROR_CODE_INVENTORY_MOUNT_IS_NEGATIVE = 1505
	ERROR_CODE_UNIT_NOT_EXSIT              = 1506
	ERROR_CODE_UNIT_NOT_MATCH              = 1507
	ERROR_CODE_PRODUCT_NOT_EXIST           = 1508
	ERROR_CODE_PRODUCT_PARAM_ERROR         = 1509
	ERROR_CODE_UNIT_NAME_IS_NULL           = 1510
	ERROR_CODE_CUSTOMER_PARAM_ERROR        = 1511
	ERROR_CODE_SUPPLIER_PARAM_ERROR        = 1512
	ERROR_CODE_DELIVERY_NOT_EXIST          = 1513
	ERROR_CODE_INVENTORY_NOT_EXIST         = 1514
	ERROR_CODE_DELIVERY_IS_DELETED         = 1515
	ERROR_CODE_TIME_FORMAT_ERROR           = 1516
	ERROR_CODE_PRODUCT_EXIST               = 1517
	ERROR_CODE_SUPPLIER_EXIST              = 1518
	ERROR_CODE_CUSTOMER_EXIST              = 1519
	ERROR_CODE_PURCHASE_NOT_EXIST          = 1520
)

var ErrorCodeMessageMap = map[int]string{
	ERROR_CODE_PARAMETER_FORMAT_INVALID: "参数格式不对",
	ERROR_CODE_ACCESS_TOKEN_TIMEOUT:     "access token失效",
	ERROR_CODE_NOT_FOUND:                "资源没有找到",
	ERROR_INTERNAL_SERVER_ERROR:         "参数业务验证失败",
	ERROR_CODE_ACCESS_TOKEN_ERROR:       "access token错误",
	ERROR_CODE_MOBILE_IS_REGISTERED:     "手机已经登录",
	ERROR_CODE_MOBILE_IS_NOT_EXISTED:    "手机不存在",
	ERROR_CODE_CAPTCHA_TYPE_ERROR:       "验证码类型错误",
	ERROR_CODE_USER_NOT_FOUND:           "表中不存在用户",
	ERROR_CODE_MOBILE_IS_INVALID:        "手机号码无效",
	ERROR_CODE_REQUEST_SMS_OVERLOAD:     "申请验证码超出次数",
	ERROR_CODE_CAPTCHA_TOKEN_ERROR:      "验证码token错误",
	ERROR_CODE_REDIS_KEY_NULL:           "redis: key值为空",
	ERROR_CODE_REDIS_VALUE_NULL_PTR:     "redis: value值为空指针",
	ERROR_CODE_REDIS_VALUE_NULL:         "redis: value值为空",
	ERROR_CODE_REDIS_KEY_NOT_EXIST:      "redis: key不存在或过期",

	ERROR_CODE_ALIOSS_CONFIG_IS_EMPTY:          "阿里oss配置为空",
	ERROR_CODE_UPLOAD_FILE_CONTENT_IS_EMPTY:    "上传文件内容为空",
	ERROR_CODE_GET_BASE_URL_ERROR:              "获取oss基本url和endpoint失败",
	ERROR_CODE_GET_FILE_BY_URL_ERROR:           "通过url获取网络文件失败",
	ERROR_CODE_GET_DEPARTMENT_ROLE_ERROR:       "获取门店角色详情失败",
	ERROR_CODE_DEFAULT_ROLE_ERROR:              "该角色为默认角色",
	ERROR_CODE_DEFAULT_RULE_ERROR:              "该权限为默认权限",
	ERROR_CODE_ROLE_NOT_EXSIT_ERROR:            "角色不存在",
	ERROR_CODE_RULE_NOT_EXSIT_ERROR:            "权限不存在",
	ERROR_CODE_DEPARTMENT_ROLE_NOT_EXSIT_ERROR: "部门角色不存在",
	ERROR_CODE_RPC_ADDRESS_NOT_EXSIT_ERROR:     "Rpc地址不存在",
	ERROR_CODE_MULTI_DEPARTMENT_IDS_ERROR:      "多门店id列表字符串有误",

	ERROR_CODE_KAFKA_PRODUCE_NIL: "producer对象为空",
	ERROR_CODE_KAFKA_TOPIC_EMPTY: "topic为空",
	ERROR_CODE_KAFKA_PUSH_FAIL:   "推送失败",

	ERROR_CODE_LOCATION_NOT_EXSIT:          "仓库不存在",
	ERROR_CODE_INVENTORY_INSUFFICIENCY:     "库存不足",
	ERROR_CODE_INVENTORY_MOUNT_IS_NEGATIVE: "库存总成本为负",
	ERROR_CODE_UNIT_NOT_EXSIT:              "单位不存在",
	ERROR_CODE_UNIT_NOT_MATCH:              "单位不匹配",
	ERROR_CODE_PRODUCT_NOT_EXIST:           "商品不存在",
	ERROR_CODE_PRODUCT_PARAM_ERROR:         "商品参数错误",
	ERROR_CODE_UNIT_NAME_IS_NULL:           "单位名称为空",
	ERROR_CODE_CUSTOMER_PARAM_ERROR:        "客户参数错误",
	ERROR_CODE_SUPPLIER_PARAM_ERROR:        "供应商参数错误",
	ERROR_CODE_DELIVERY_NOT_EXIST:          "订单出库单不存在",
	ERROR_CODE_INVENTORY_NOT_EXIST:         "库存不存在",
	ERROR_CODE_DELIVERY_IS_DELETED:         "订单出库单已经撤销",
	ERROR_CODE_TIME_FORMAT_ERROR:           "时间格式错误",
	ERROR_CODE_PRODUCT_EXIST:               "商品已经存在",
	ERROR_CODE_SUPPLIER_EXIST:              "供应商已经存在",
	ERROR_CODE_CUSTOMER_EXIST:              "客户已经存在",
	ERROR_CODE_PURCHASE_NOT_EXIST:          "采购入库单不存在",
}

type HError struct {
	ResCode int    `json:"code" description:"0成功 >1失败"`
	Message string `json:"message" description:"提示消息"`
}

func (ze HError) Error() string {
	return fmt.Sprintf("%s%s", HERROR_PREFIX, ze.Message)
}

func (ze HError) Code() int {
	return ze.ResCode
}

//系统错误返回数据结构
func NewHError(code int, message string) error {
	return &HError{code, message}
}

//自定义错误返回数据结构
func NewHErrorCustom(code int) error {
	return &HError{code, ErrorCodeMessageMap[code]}
}
