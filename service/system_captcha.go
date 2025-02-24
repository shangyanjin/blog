package service

import (
	"blog/model"
	"bytes"
	"encoding/base64"
	"errors"
	"net/http"

	"github.com/dchest/captcha"
	"gorm.io/gorm"
)

type CaptchaService struct {
	Length int // Length of the captcha code
	DB     *gorm.DB
}

var Captcha = NewCaptchaService()

const ()

// NewCaptchaService initializes the img service
func NewCaptchaService() *CaptchaService {
	return &CaptchaService{
		Length: 4,
		DB:     model.GetDb(),
	}
}

// GenerateCaptcha generates a new captcha ID and returns the captcha img as a base64 encoded string
func (this *CaptchaService) GenerateCaptcha() *model.Data {
	// Generate a new captcha ID
	id := captcha.NewLen(this.Length)
	// Create a buffer to store the captcha img
	var buf bytes.Buffer
	// Write the captcha img to the buffer
	if err := captcha.WriteImage(&buf, id, 240, 80); err != nil {
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Message: "Failed to generate captcha img",
			Error:   err,
		}
	}
	// Encode the img to base64
	base64Image := base64.StdEncoding.EncodeToString(buf.Bytes())
	data := map[string]string{
		"id":    id,
		"image": base64Image,
	}
	return &model.Data{
		Code:    http.StatusOK,
		Data:    data,
		Message: "Captcha generated successfully",
	}
}

// ValidateCaptcha checks if the given answer is correct for the specified captcha ID
func (this *CaptchaService) ValidateCaptcha(id string, answer string) *model.Data {

	if !captcha.VerifyString(id, answer) {
		return &model.Data{Code: http.StatusInternalServerError, Message: "Invalid captcha",
			Data: map[string]string{
				"id":     id,
				"answer": answer,
				"valid":  "false",
			},
			Error: errors.New("invalid captcha"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "Captcha is valid",
		Data: map[string]string{
			"id":     id,
			"answer": answer,
			"valid":  "true",
		},
	}
}

// ServeCaptchaImageAsBase64 serves the captcha img as a base64 encoded string
func (this *CaptchaService) ServeCaptchaImageAsBase64(id string) (base64Image string, e error) {
	var buf bytes.Buffer
	if err := captcha.WriteImage(&buf, id, 240, 80); err != nil {
		return "", errors.New("failed to serve captcha img")
	}
	base64Image = base64.StdEncoding.EncodeToString(buf.Bytes())
	return base64Image, nil
}

/*
package service

import (
	"bytes"
	"encoding/base64"
	"errors"
	"blog/model"
	"blog/response"
	"net/http"

	"github.com/dchest/captcha"
	"gorm.io/gorm"
)

type CaptchaService struct {
	Length int // Length of the captcha code
	DB     *gorm.DB
}

var Captcha = NewCaptchaService()

const ()

// NewCaptchaService initializes the img service
func NewCaptchaService() *CaptchaService {
	return &CaptchaService{
		Length: 4,
		DB:     model.GetDb(),
	}
}

// GenerateCaptcha generates a new captcha ID and returns the captcha img as a base64 encoded string
func (this *CaptchaService) GenerateCaptcha() *model.Data {
	// Generate a new captcha ID
	id := captcha.NewLen(this.Length)
	// Create a buffer to store the captcha img
	var buf bytes.Buffer
	// Write the captcha img to the buffer
	if err := captcha.WriteImage(&buf, id, 240, 80); err != nil {
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Message: "Failed to generate captcha img",
		}
	}
	// Encode the img to base64
	base64Image := base64.StdEncoding.EncodeToString(buf.Bytes())
	data := map[string]string{
		"id":    id,
		"img": base64Image,
	}
	return &model.Data{
		Code:    http.StatusOK,
		Data:    data,
		Message: "Captcha generated successfully",
	}
}

// ValidateCaptcha checks if the given answer is correct for the specified captcha ID
func (this *CaptchaService) ValidateCaptcha(id string, answer string) *model.Data {
	if !captcha.VerifyString(id, answer) {
		return &model.Data{
			Code:    http.StatusInternalServerError,
			Message: "Invalid captcha",
			Error:   errors.New("invalid captcha"),
		}
	}

	return &model.Data{Code: http.StatusOK, Message: "Captcha is valid"}
}

// ServeCaptchaImageAsBase64 serves the captcha img as a base64 encoded string
func (this *CaptchaService) ServeCaptchaImageAsBase64(id string) (base64Image string, e error) {
	var buf bytes.Buffer
	if err := captcha.WriteImage(&buf, id, 240, 80); err != nil {
		return "", errors.New("failed to serve captcha img")
	}
	base64Image = base64.StdEncoding.EncodeToString(buf.Bytes())
	return base64Image, nil
}

*/
