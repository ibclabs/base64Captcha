package base64Captcha

import "errors"

func CreateBothDigits(id string, configDigit ConfigDigit, configAudio ConfigAudio) (image *CaptchaImageDigit, audio *Audio, value string, err error) {
	if configDigit.CaptchaLen != configAudio.CaptchaLen {
		err = errors.New("len mismatch")
	}
	digits := randomDigits(configDigit.CaptchaLen)
	image = newDigits(id, digits, configDigit)
	audio = newAudio(id, digits, configAudio)
	if image.VerifyValue != audio.VerifyValue {
		err = errors.New("value mismatch")
	}
	value = image.VerifyValue
	return
}
