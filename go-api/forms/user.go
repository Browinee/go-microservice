package forms

// NOTE: like nest dto with class-validator and class-transformer
type PassWordLoginForm struct {
	Mobile   string `form:"mobile"   json:"mobile" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=3,max=20"`
	// Password string `form:"password" json:"password" binding:"required,min=3,max=20,mobile"`

}