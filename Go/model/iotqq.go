package iotqq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var url1, qq string

type QQinfo struct {
	Code    int    `json:"code"`
	Data    Data1  `json:"data"`
	Default int    `json:"default"`
	Message string `json:"message"`
	Subcode int    `json:"subcode"`
}
type Data1 struct {
	AvatarURL     string `json:"avatarUrl"`
	Bitmap        string `json:"bitmap"`
	Commfrd       int    `json:"commfrd"`
	Friendship    int    `json:"friendship"`
	Greenvip      int    `json:"greenvip"`
	IntimacyScore int    `json:"intimacyScore"`
	IsFriend      int    `json:"isFriend"`
	Logolabel     string `json:"logolabel"`
	Nickname      string `json:"nickname"`
	Qqvip         int    `json:"qqvip"`
	Qzone         int    `json:"qzone"`
	Realname      string `json:"realname"`
	Redvip        int    `json:"redvip"`
	Smartname     string `json:"smartname"`
	Uin           int    `json:"uin"`
}
type QQ struct {
	Cont int
}
type PSkey struct {
	Connect     string `json:"connect"`
	Docs        string `json:"docs"`
	Docx        string `json:"docx"`
	Game        string `json:"game"`
	Gamecenter  string `json:"gamecenter"`
	Imgcache    string `json:"imgcache"`
	MTencentCom string `json:"m.tencent.com"`
	Mail        string `json:"mail"`
	Mma         string `json:"mma"`
	Now         string `json:"now"`
	Office      string `json:"office"`
	Openmobile  string `json:"openmobile"`
	Qqweb       string `json:"qqweb"`
	Qun         string `json:"qun"`
	Qzone       string `json:"qzone"`
	QzoneCom    string `json:"qzone.com"`
	TenpayCom   string `json:"tenpay.com"`
	Ti          string `json:"ti"`
	Vip         string `json:"vip"`
	Weishi      string `json:"weishi"`
}
type Cook struct {
	ClientKey string `json:"ClientKey"`
	Cookies   string `json:"Cookies"`
	Gtk       string `json:"Gtk"`
	Gtk32     string `json:"Gtk32"`
	PSkey     PSkey  `json:"PSkey"`
	Skey      string `json:"Skey"`
}
type Conf struct {
	Enable bool
	GData  map[string]int
}
type Data2 struct {
	Date   string `json:"date"`
	City   string `json:"city"`
	Adcode string `json:"adcode"`
	Min    string `json:"min"`
	Max    string `json:"max"`
	Type   string `json:"type"`
	Air    string `json:"air"`
	Wind   string `json:"wind"`
}
type Weather struct {
	Code int   `json:"code"`
	Data Data2 `json:"data"`
}
type CurrentPacket struct {
	Data      Data   `json:"Data"`
	WebConnID string `json:"WebConnId"`
}
type Data struct {
	Content       string      `json:"Content"`
	FromGroupID   int         `json:"FromGroupId"`
	FromGroupName string      `json:"FromGroupName"`
	FromNickName  string      `json:"FromNickName"`
	FromUserID    int64       `json:"FromUserId"`
	MsgRandom     int         `json:"MsgRandom"`
	MsgSeq        int         `json:"MsgSeq"`
	MsgTime       int         `json:"MsgTime"`
	MsgType       string      `json:"MsgType"`
	RedBaginfo    interface{} `json:"RedBaginfo"`
}
type Message struct {
	CurrentPacket CurrentPacket `json:"CurrentPacket"`
	CurrentQQ     int64         `json:"CurrentQQ"`
}
type Channel struct {
	Channel string `json:"channel"`
}

func Set(url string, qq1 string) {
	qq = qq1
	url1 = url
}
func GetCook() Cook {
	resp, err := http.Get("http://" + url1 + "/v1/LuaApiCaller?funcname=GetUserCook&timeout=10&qq=" + qq)
	if err != nil {
		log.Fatal(err)

	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	tmp := string(body)
	var thecook Cook
	err = json.Unmarshal([]byte(tmp), &thecook)
	if err != nil {
		fmt.Println("反序列化出错,info:", err)
	}
	return thecook
}
func SendPic(ToUser int, SendToType int, Content string, PicUrl string) {
	//发送图文信息
	tmp := make(map[string]interface{})
	tmp["toUser"] = ToUser
	tmp["sendToType"] = SendToType
	tmp["sendMsgType"] = "PicMsg"
	tmp["picBase64Buf"] = ""
	tmp["fileMd5"] = ""
	tmp["picUrl"] = PicUrl
	tmp["content"] = Content
	tmp["groupid"] = 0
	tmp["atUser"] = 0
	tmp["pwd"] = "mcoo"
	tmp1, _ := json.Marshal(tmp)
	resp, err := (http.Post("http://"+url1+"/v1/LuaApiCaller?funcname=SendMsg&timeout=10&qq="+qq, "application/json", bytes.NewBuffer(tmp1)))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
func Send(ToUser int, SendToType int, Content string) {
	//发送文本信息
	tmp := make(map[string]interface{})
	tmp["toUser"] = ToUser
	tmp["sendToType"] = SendToType
	tmp["sendMsgType"] = "TextMsg"
	tmp["content"] = Content
	tmp["groupid"] = 0
	tmp["atUser"] = 0
	tmp["pwd"] = "mcoo"
	tmp1, _ := json.Marshal(tmp)
	resp, err := (http.Post("http://"+url1+"/v1/LuaApiCaller?funcname=SendMsg&timeout=10&qq="+qq, "application/json", bytes.NewBuffer(tmp1)))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
func SendA(ToUser int, SendToType int, Content string, SendMsgType string) {
	//发送其他信息
	tmp := make(map[string]interface{})
	tmp["toUser"] = ToUser
	tmp["sendToType"] = SendToType
	tmp["sendMsgType"] = SendMsgType
	tmp["content"] = Content
	tmp["groupid"] = 0
	tmp["atUser"] = 0
	tmp["pwd"] = "mcoo"
	tmp1, _ := json.Marshal(tmp)
	resp, err := (http.Post("http://"+url1+"/v1/LuaApiCaller?funcname=SendMsg&timeout=10&qq="+qq, "application/json", bytes.NewBuffer(tmp1)))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
func SendVoice(ToUser int, SendToType int, Content string) {
	//发送语音信息
	tmp := make(map[string]interface{})
	tmp["toUser"] = ToUser
	tmp["sendToType"] = SendToType
	tmp["sendMsgType"] = "VoiceMsg"
	tmp["content"] = ""
	tmp["voiceUrl"] = "https://dds.dui.ai/runtime/v1/synthesize?voiceId=qianranfa&speed=0.7&volume=100&audioType=wav&text=" + url.PathEscape(Content)
	tmp["groupid"] = 0
	tmp["atUser"] = 0
	tmp["voiceBase64Buf"] = ""
	tmp["pwd"] = "mcoo"
	tmp1, _ := json.Marshal(tmp)
	resp, err := (http.Post("http://"+url1+"/v1/LuaApiCaller?funcname=SendMsg&timeout=10&qq="+qq, "application/json", bytes.NewBuffer(tmp1)))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
func Zan(qq1 int, err error) {
	//名片点赞
	tmp := make(map[string]interface{})
	tmp["UserID"] = qq1
	tmp["pwd"] = "mcoo"
	tmp1, _ := json.Marshal(tmp)
	fmt.Println(string(tmp1))
	resp, err := (http.Post("http://"+url1+"/v1/LuaApiCaller?funcname=QQZan&timeout=10&qq="+qq, "application/json", bytes.NewBuffer(tmp1)))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
func Getinfo(qq1 int) string {
	tmp := make(map[string]interface{})
	tmp["UserID"] = qq1
	tmp["pwd"] = "mcoo"
	tmp1, _ := json.Marshal(tmp)
	resp, err := (http.Post("http://"+url1+"/v1/LuaApiCaller?funcname=GetUserInfo&timeout=10&qq="+qq, "application/json", bytes.NewBuffer(tmp1)))
	if err != nil {
		log.Fatal(err)
		return "err"
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
