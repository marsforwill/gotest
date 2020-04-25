package main

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/binary"
	"fmt"
)

type MallcooTraceBody struct {
	EventType int              `json:"EventType,omitempty"`
	EventData MallcooEventData `json:"EventData,omitempty"`
}
type MallcooEventData struct {
	Time            string `json:"time,omitempty"`
	UniqueID        string `json:"unique_id,omitempty"`
	PositionID      string `json:"position_id,omitempty"`
	FaceID          string `json:"face_id,omitempty"`
	VendorID        string `json:"vendor_id,omitempty"`
	FaceImageURL    string `json:"face_img_url,omitempty"`
	IsNew           string `json:"is_new,omitempty"`
	AccessDirection string `json:"access_direction,omitempty"`
}

func md5Tool(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5Str := fmt.Sprintf("%x", has)

	return md5Str
}

type DeleteAreaInfo struct {
	RequestID string `json:"request_id"`
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

//GetFeatureWithNoHeader 从base64编码获取feature变量,该feature不包含Header，长度为在参数中
func GetFeatureWithNoHeader(base64Str string, length int64) []float32 {
	// base64解码
	deBytes, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return nil
	}
	if int64(len(deBytes)) < length*4 {
		return nil
	}
	// 获取feature
	var feature = make([]float32, length)
	for i := int64(0); i < length; i++ {
		b := []byte{deBytes[3+4*i], deBytes[2+4*i], deBytes[1+4*i], deBytes[4*i]}
		buf := bytes.NewBuffer(b)
		binary.Read(buf, binary.BigEndian, &feature[i])
	}
	return feature
}

func FeatureToString(feature []float32) string {
	var feaBuf bytes.Buffer
	binary.Write(&feaBuf, binary.LittleEndian, feature)
	return string(feaBuf.Bytes())
}

func main() {
	fmt.Printf("hello atty")

}

func excute(i int) {

	print(i)
}
