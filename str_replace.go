package util

import (
	"fmt"
	"strings"
)

func testStrReplace() {
	//str := ```	```
	/*str1 := `{
	    "actionCard": {
	        "title": "乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身",
	        "text": "![screenshot](@lADOpwk3K80C0M0FoA)
	 ### 乔布斯 20 年前想打造的苹果咖啡厅
	 Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划,20 年前想打造的苹果咖啡厅
	 Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划,20 年前想打造的苹果咖啡厅
	 Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划,20 年前想打造的苹果咖啡厅
	 Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划,20 年前想打造的苹果咖啡厅
	 Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划,20 年前想打造的苹果咖啡厅
	 Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划",
	        "hideAvatar": "0",
	        "btnOrientation": "0",
	        "singleTitle" : "阅读全文",
	        "singleURL" : "https://www.dingtalk.com/"
	    },
	    "msgtype": "actionCard"
	}`*/
	str := `{
    "actionCard": {
        "title": "乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身", 
        "text": "![screenshot](@lADOpwk3K80C0M0FoA) 
 ### 乔布斯 20 年前想打造的苹果咖啡厅 
 Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划", 
        "hideAvatar": "0", 
        "btnOrientation": "0", 
        "singleTitle" : "阅读全文",
        "singleURL" : "http://47.106.192.182:9066/sys.txt"
    }, 
    "msgtype": "actionCard"
}`
	var strNew string
	temp := []byte(str)
	strNew = string(temp)
	strNew = strings.Replace(strNew, "\n", "", -1)
	strNew = strings.Replace(strNew, "\t", "", -1)
	strNew = strings.Replace(strNew, "", "", -1)
	fmt.Println(str)
	fmt.Println(strNew)
}
