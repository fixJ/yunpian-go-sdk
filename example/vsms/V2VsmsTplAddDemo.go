package main

import (
	"encoding/json"
	"fmt"
	"github.com/fixJ/yunpian-go-sdk/client"
	"github.com/fixJ/yunpian-go-sdk/param"
)

func main() {
	// 视频短信模版添加参数需要layout，material
	// V2VsmsTplAddParam中的Layout是一个json字符串，对应的结构体为V2VsmsLayout
	// material为压缩包的字节码序列，使用sdk只需要传入压缩包路径，sdk会读取并转成字节码序列

	//创建layout
	layout := param.V2VsmsLayout{
		VlVersion: "0.0.1",
		Subject: "标题",
		Frames: []param.V2VsmsFrames {
			{
				Index: 1,
				PlayTimes: 1,
				Attachments: []param.V2VsmsAttachment {
					{
						Index: 1,
						FileName: "content.txt",
					},
					{
						Index: 2,
						FileName: "show.mp4",
					},
					{
						Index: 3,
						FileName: "show_img.png",
					},
				},
			},
		},
	}

	// 将layout转成json字符串
	layoutBytes, err := json.Marshal(&layout)
	if err != nil {
		fmt.Println(err)
	}
	p := param.V2VsmsTplAddParam{
		Sign: "【视频】",
		Layout: string(layoutBytes),
	}
	c := client.InitDefaultClient("your apikey")
	r, _ := c.Vsms.V2VsmsTplAdd(p, "your path")
	fmt.Println(r.V2VsmsTplAddResult)
}
