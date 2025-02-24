package api // router api

import (
	"blog/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	
	"net/http"
)

var Captcha = captchaHandle{}

type captchaHandle struct {
}

// handle index routes

func (this captchaHandle) RouterGroup(r *gin.RouterGroup) {
	r.GET("/captcha/get", this.generateCaptcha)   // 添加生成验证码的路由
	r.POST("/captcha/verify", this.verifyCaptcha) // 添加验证验证码的路由
}

// generateCaptcha 生成验证码
func (this captchaHandle) generateCaptcha(c *gin.Context) {
	result := service.Captcha.GenerateCaptcha()
	if result.Error != nil {
		c.JSON(500, gin.H{"code": http.StatusInternalServerError, "message": "无法生成验证码", "error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "success", "data": result.Data})
}

// verifyCaptcha 验证验证码
func (this captchaHandle) verifyCaptcha(c *gin.Context) {
	var req struct {
		CaptchaID string `json:"id"`
		Answer    string `json:"answer"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "请求格式错误"})
		return
	}
	result := service.Captcha.ValidateCaptcha(req.CaptchaID, req.Answer)
	if result.Error != nil {
		logrus.Errorf("ValidateCaptcha Error: %v", result.Error)
		c.JSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "message": "Invalid captcha", "error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "success", "data": result.Data})
}
