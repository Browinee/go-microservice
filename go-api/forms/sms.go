package forms

type SendSmsForm struct{
	Mobile   string `form:"mobile" json:"mobile" binding:"required"`
	// NOTE: 註冊/密碼都需要驗證碼，所以有不同的type
	// 1:register 2.動態驗證碼登入
	Type uint  `form:"type" json:"type" binding:"required,oneof=1 2"`
}