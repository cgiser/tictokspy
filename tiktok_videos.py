# coding=utf-8
"""
* create by hujia on 2019/5/23
* email : Healy.hu@yeahmobi.com

* tiktok个人主页下的视频抓取
"""
import json
import pymongo
import random
import time
from urllib import urlencode

import pymysql
import requests
import sys

reload(sys)
sys.setdefaultencoding('utf8')

HEADER = {
    "Cookie": "ttreq=1$8f383426a2f4a83505502c07879de1fcf5f16638; odin_tt=f8975fde5e23cfb1c012bd948de41e1be4067620e1260f53431a7ac583601b3609baad2574f69d65394560d5fb7bb039c0390cea8469f8539cba76b6b7cd9875",
    "Accept-Encoding": "gzip",
    "X-SS-REQ-TICKET": "1582013667731",
    "sdk-version": "1",
    "User-Agent": "okhttp/3.10.0.1",
    "Connection": "keep-alive"}

# 标识是哪个设备请求，目前的接口支持随意更换设备id
input_id = sys.argv[1]  # 6502641543550222337 user_id
if len(sys.argv) > 2:
    TOTAL_COUNT = sys.argv[2]
else:
    TOTAL_COUNT = 100000
if len(sys.argv) > 3:
    proxies = sys.argv[3]
else:
    proxies = ""

# 初始化mongodb 并连接表
db_config = {
    "host": "10.12.35.93",
    "port": 27017,
    # 需要时释放以下参数
    "username": "pro_tictok",
    "password": "cKG5g32*2AMWa3dFW3zS",
    "authSource": "admin",
    "authMechanism": "SCRAM-SHA-1"
}
db_name = "tictok_new"
client = pymongo.MongoClient(**db_config)
tiktok_db = client[db_name]
tiktok_comment_table = tiktok_db["crawler-tiktok-comment"]
tiktok_videos_table = tiktok_db["crawler-tiktok-videos"]
tiktok_user_table = tiktok_db["crawler-tiktok-user"]

if proxies == "":
    PROXIES = ""
else:
    PROXIES = {"http": "http://" + proxies, "https": "https://" + proxies}  # 请求的代理

#
# user_id = 6502641543550222337
# TOTAL_COMMENT_COUNT = 60
# TOTAL_PAGE = 0
# PROXIES = ""


class TiktokPhotoCrawler(object):
    deviced_id = random.randint(6000000000000000000, 6999999999999999999)
    ricket = int(time.time() * 1000)
    ts = int(ricket / 1000)
    PARAMS = {"max_cursor": 0, "user_id": str(input_id), "count": 20, "retry_type": "no_retry",
              "mcc_mnc": 44000, "app_language": "en", "language": "ja", "region": "US", "sys_region": "JP",
              "carrier_region": "JP", "carrier_region_v2": "", "build_number": "5.1.8", "timezone_offset": 32400,
              "timezone_name": "Asia/Tokyo", "is_my_cn": 0, "fp": "a_fake_fp", "account_region": "US",
              "pass-region": 1, "pass-route": 1, "device_id": deviced_id, "ac": "wifi", "channel": "googleplay",
              "aid": 1180, "app_name": "trill", "version_code": 518, "version_name": "5.1.8",
              "device_platform": "android", "ab_version": "5.1.8", "ssmix": "a", "device_type": "HUAWEI MLA-AL10",
              "device_brand": "HUAWEI", "os_api": 19, "os_version": "5.1.1", "openudid": "F01898567C920000",
              "manifest_version_code": 518, "resolution": "900*1440", "dpi": 320, "update_version_code": 5180,
              "_rticket": ricket, "ts": ts, "as": "a185fae22e451ee4066277", "cp": "a551e55be66e2a4ae1McUg",
              "mas": "0136dc011cef2d388f24691732a42e2a28ecec4c6c6c0c2ca6a68c"}
    URL_PREFIX = "https://api21-h2.tiktokv.com/aweme/v1/aweme/post/?"

    def __init__(self):
        self.header = HEADER
        self.url_prefix = TiktokPhotoCrawler.URL_PREFIX
        self.params = TiktokPhotoCrawler.PARAMS
        self.proxies = PROXIES
        self.user_count = 0

    def crawl(self):
        first_url = self._make_first_request()
        res = requests.get(url=first_url, headers=self.header)
        content = res.content
        content_dict = json.loads(content)
        max_cursor = content_dict['max_cursor'] if "max_cursor" in content_dict else 0
        has_more = content_dict['has_more'] if "has_more" in content_dict else 0
        photo_list = content_dict['aweme_list'] if 'aweme_list' in content_dict else ''
        if not photo_list:
            return
        total_count = self.parse_photo(photo_list)
        while has_more and int(total_count) < int(TOTAL_COUNT):
            print ("crawler page is : " + str(total_count))
            next_url = self._make_next_request(max_cursor)
            res = requests.get(url=next_url, headers=self.header, proxies=self.proxies)
            content = res.content
            # print content
            json_data = json.loads(content)
            has_more = json_data["has_more"] if "has_more" in json_data else 0
            max_cursor = json_data['max_cursor']
            photo_list = json_data["aweme_list"] if "aweme_list" in json_data else ""
            if not photo_list:
                return
            photo_count = self.parse_photo(photo_list)
            total_count += photo_count

    def parse_photo(self, photo_list):
        photo_count = 0
        for ele in photo_list:
            item = dict()
            photo_id = int(ele["aweme_id"])
            # item["_id"] = photo_id
            content = dict()
            content["deccription"] = ele["desc"]
            content["publish_time"] = int(ele["create_time"]) * 1000

            video_info = ele["video"] if "video" in ele else ""
            video_addr = video_info["play_addr"] if "play_addr" in video_info else ""
            video_url_list = video_addr["url_list"] if "url_list" in video_addr else ""
            content["photo_url"] = str(video_url_list[0])

            cover_addr = video_info["cover"] if "cover" in video_info else ""
            cover_url_list = cover_addr["url_list"] if "url_list" in cover_addr else ""
            content["cover_url"] = str(cover_url_list[0])

            count_info = ele["statistics"] if "statistics" in ele else ""
            content["comment_count"] = count_info["comment_count"]
            content["like_count"] = count_info["comment_count"]
            content["download_count"] = count_info["download_count"]
            content["view_count"] = count_info["play_count"]
            content["share_count"] = count_info["share_count"]

            author_info = ele["author"] if "author" in ele else ""
            user_id = author_info["uid"] if "uid" in author_info else 0
            if self.user_count == 0:
                user_item = dict()
                # user_item["_id"] = user_id
                user_item["content"] = author_info
                tiktok_user_table.update({"_id": user_id}, {"$set": user_item}, True)
                self.user_count = 1
                print "insert user sucess,user_id is : %s" % str(user_id)

            content["user_id"] = user_id
            content["unique_id"] = author_info["short_id"] if "short_id" in author_info else ""
            content["nick_name"] = author_info["nickname"] if "nickname" in author_info else ""
            avatar_info = author_info["avatar_larger"] if "avatar_larger" in author_info else ""
            avatar_list = avatar_info["url_list"] if "url_list" in avatar_info else ""
            content["avatar"] = str(avatar_list[0])
            content["region"] = ele["region"] if "region" in ele else ""
            hash_tag = ele["text_extra"] if "text_extra" in ele else ""
            tag_list = []
            if hash_tag:
                for ele in hash_tag:
                    tag = ele["hashtag_name"] if "hashtag_name" in ele else ""
                    if tag:
                        tag_list.append(tag)

            content["tags"] = tag_list
            content["update_time"] = int(time.time())
            item["content"] = content
            tiktok_videos_table.update({"_id": photo_id}, {"$set": item}, True)
            print "insert videos sucess,photo_id is : %s" % str(photo_id)
            photo_count += 1
        return photo_count

    def _make_first_request(self):
        url = self.url_prefix + urlencode(self.params)
        return url

    def _make_next_request(self, max_cursor):
        self.params['max_cursor'] = max_cursor
        url = self.url_prefix + urlencode(self.params)
        return url


if __name__ == "__main__":
    tiktok_videos_crawler = TiktokPhotoCrawler()
    tiktok_videos_crawler.crawl()
    client.close()
