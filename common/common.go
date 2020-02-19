package common

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"net/http"
	"net/url"
)

var (
	MogUrl string
	MogDB  string
	Param  *url.Values
	Header *http.Header
	ApiUrl = "https://api21-h2.tiktokv.com/aweme/v1/aweme/post/?"
)
var (
//mgoBidSession *mgo.Session

)

func init() {
	MogUrl = "mongodb://pro_tictok:cKG5g32*2AMWa3dFW3zS@116.202.128.164:27017/tictok?serverSelectionTimeoutMS=5000&connectTimeoutMS=10000&authSource=admin&authMechanism=SCRAM-SHA-256"
	//MogUrl = "mongodb://pro_push:Q4M5tHSvUBWEN2t5WKNF@52.91.43.99,34.227.206.79,54.210.24.215:27017/push?maxPoolSize=100&authSource=admin"
	MogDB = "tictok"

	header := `{"Cookie": "ttreq=1$8f383426a2f4a83505502c07879de1fcf5f16638; odin_tt=f8975fde5e23cfb1c012bd948de41e1be4067620e1260f53431a7ac583601b3609baad2574f69d65394560d5fb7bb039c0390cea8469f8539cba76b6b7cd9875",
    "Accept-Encoding": "gzip",
    "X-SS-REQ-TICKET": "1582013667731",
    "sdk-version": "1",
    "User-Agent": "okhttp/3.10.0.1",
    "Connection": "keep-alive"}`
	param := `
		{
	"max_cursor": 0,
	"user_id": 0,
	"count": 20,
	"retry_type": "no_retry",
	"mcc_mnc": 44000,
	"app_language": "en",
	"language": "ja",
	"region": "US",
	"sys_region": "JP",
	"carrier_region": "JP",
	"carrier_region_v2": "",
	"build_number": "5.1.8",
	"timezone_offset": 32400,
	"timezone_name": "Asia/Tokyo",
	"is_my_cn": 0,
	"fp": "a_fake_fp",
	"account_region": "US",
	"pass-region": 1,
	"pass-route": 1,
	"device_id": 0,
	"ac": "wifi",
	"channel": "googleplay",
	"aid": 1180,
	"app_name": "trill",
	"version_code": 518,
	"version_name": "5.1.8",
	"device_platform": "android",
	"ab_version": "5.1.8",
	"ssmix": "a",
	"device_type": "HUAWEI MLA-AL10",
	"device_brand": "HUAWEI",
	"os_api": 19,
	"os_version": "5.1.1",
	"openudid": "F01898567C920000",
	"manifest_version_code": 518,
	"resolution": "900*1440",
	"dpi": 320,
	"update_version_code": 5180,
	"_rticket": 0,
	"ts": 0,
	"as": "a185fae22e451ee4066277",
	"cp": "a551e55be66e2a4ae1McUg",
	"mas": "0136dc011cef2d388f24691732a42e2a28ecec4c6c6c0c2ca6a68c"
		}
	`

	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(param), &m)
	if err != nil {
		fmt.Println(err)
	} else {
		v := url.Values{}
		for key, value := range m {
			v.Add(key, fmt.Sprintf("%v", value))
		}
		Param = &v
		//fmt.Println(v.Encode())
	}

	m1 := make(map[string]interface{})
	err1 := json.Unmarshal([]byte(header), &m1)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		v := http.Header{}
		for key, value := range m {
			v.Add(key, fmt.Sprintf("%v", value))
		}
		Header = &v
		//fmt.Println(v.Encode())
	}

}

///**
// * 公共方法，获取session，如果存在则拷贝一份
// */
//func GetMongoSession() {
//	// Set client options
//	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
//
//	// Connect to MongoDB
//	client, err := mongo.Connect(context.TODO(), clientOptions)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Check the connection
//	err = client.Ping(context.TODO(), nil)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println("Connected to MongoDB!")
//}
