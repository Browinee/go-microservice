package forms

// NOTE: like nest dto with class-validator and class-transformer
type PassWordLoginForm struct {
	Mobile   string `form:"mobile"   json:"mobile" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=3,max=20"`
	// Password string `form:"password" json:"password" binding:"required,min=3,max=20,mobile"`
	Captcha string `form:"captcha" json:"captcha" binding:"required,min=5,max=5"`
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"`

}


type RegisterForm struct{
	Mobile   string `form:"mobile"   json:"mobile" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=3,max=20"`
	Code string `form"code" json:"code" binding:"required,min=6,max=6"`
}