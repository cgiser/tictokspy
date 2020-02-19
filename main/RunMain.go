package main

func main()  {
	//deviced_id := rand.Int63n(999999999999999999)+6000000000000000000
	//_rticket := time.Now().Unix()*1000
	//ts :=time.Now().Unix()
	////common.GetMongoSession()
	//fmt.Println("hellio")
	//v:=*common.Param
	//v.Set("deviced_id",fmt.Sprintf("%d",deviced_id))
	//v.Set("_rticket",fmt.Sprintf("%d",_rticket))
	//v.Set("ts",fmt.Sprintf("%d",ts))
	//v.Set("user_id","6502641543550222337")
	//client:=http.Client{}
	//
	//url:=fmt.Sprintf("%s%s",common.ApiUrl,v.Encode())
	//fmt.Println(url)
	//req, _ := http.NewRequest("GET", url, nil)
	//req.Header = *common.Header
	//res, err := client.Do(req)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	s, _ := ioutil.ReadAll(res.Body) //把  body 内容读入字符串 s
	//	fmt.Println(string(s))        //在返回页面中显示内容。
	//}
	//session := common.GetMongoSession()
	//defer session.Close()
	//bundle := session.DB(common.MogDB).C("crawler-tiktok-user")
	//query := bson.M{}
	//
	//iter := bundle.Find(query).Iter()
	//
	//user := model.TicUser{}
	//for iter.Next(&user) {
	//
	//	update := exec.Command("python2", "./tiktok_videos.py", fmt.Sprintf("%d",*user.UserId),"20")
	//	update.Dir = "/dianyi/app/tiktok_videos"
	//	update.Stdout = os.Stdout
	//	update.Stderr = os.Stderr
	//	err := update.Run()
	//	if err != nil {
	//		fmt.Println("update source, error: %v", err)
	//	}
	//}


}
