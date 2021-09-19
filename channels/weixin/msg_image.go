package weixin

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"

	"github.com/bougou/webhook-adapter/models"
)

const (
	MsgTypeImage = "image"
	maxImageSize = 2 * 1024 * 1024 // 2MB
)

func init() {
	SupportedMsgtypes[MsgTypeImage] = NewMsgImageFromPayload
}

type Image struct {
	Base64 string `json:"base64"`
	MD5    string `json:"md5"`
}

func GetMD5Hash(data []byte) []byte {
	md5sum := md5.Sum(data)
	return []byte(hex.EncodeToString(md5sum[:]))
}

func NewMsgImage(imgByte []byte) *Msg {
	imgMD5 := GetMD5Hash(imgByte)
	imgBase64 := base64.StdEncoding.EncodeToString(imgByte)

	return &Msg{
		MsgType: MsgTypeImage,
		Image: &Image{
			Base64: imgBase64,
			MD5:    string(imgMD5),
		},
	}
}

func NewMsgImageFromPayload(payload *models.Payload) *Msg {
	imgByte := []byte{}
	if len(payload.Images) > 0 {
		imgByte = payload.Images[0].Bytes
	}
	return NewMsgImage(imgByte)
}
