package forms

type LoginForm struct {
	Mobile    string `json:"mobile" binding:"required,mobile"`
	Password  string `json:"password" binding:"required"`
	CaptchaId string `json:"captchaId" binding:"required"`
	Captcha   string `json:"captcha" binding:"required,min=5,max=5"`
}
