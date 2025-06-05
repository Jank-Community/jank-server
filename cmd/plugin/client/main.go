package main

import (
	"fmt"
	"jank.com/jank_blog/internal/rpc"
	pb "jank.com/jank_blog/pkg/proto"
	"log"
)

/*
客户端案例，后续封装成接口
*/

func main() {
	// 连接到服务器
	client, err := rpc.NewClient("localhost:8080")
	if err != nil {
		log.Fatalf("连接服务器失败: %v", err)
	}
	defer client.Close()

	// 测试注册插件
	plugin := &pb.Plugin{
		Name:        "示例插件1",
		Version:     "1.0.0",
		Description: "这是一个示例插件",
		Author:      "开发者",
		Category:    "工具",
		Tags:        []string{"实用工具", "示例"},
		DownloadUrl: "https://example.com/plugin.zip",
	}

	registerResp, err := client.RegisterPlugin(plugin)
	if err != nil {
		log.Fatalf("注册插件失败: %v", err)
	}

	if registerResp.Success {
		fmt.Printf("插件注册成功, ID: %s\n", registerResp.PluginId)

		// 测试获取插件
		getResp, err := client.GetPlugin(registerResp.PluginId)
		if err != nil {
			log.Printf("获取插件失败: %v", err)
		} else {
			fmt.Printf("获取插件成功: %s\n", getResp.Plugin.Name)
		}

		// 测试列出插件
		listResp, err := client.ListPlugins("", nil, "", 1, 10, "name", true)
		if err != nil {
			log.Printf("列出插件失败: %v", err)
		} else {
			fmt.Printf("找到 %d 个插件:\n", listResp.Total)
			for _, p := range listResp.Plugins {
				fmt.Printf("- %s v%s by %s\n", p.Name, p.Version, p.Author)
			}
		}

		// 测试搜索插件
		searchResp, err := client.SearchPlugins("示例", 10)
		if err != nil {
			log.Printf("搜索插件失败: %v", err)
		} else {
			fmt.Printf("搜索结果: %d 个插件\n", len(searchResp.Plugins))
		}

		// 测试下载插件
		downloadResp, err := client.DownloadPlugin(registerResp.PluginId)
		if err != nil {
			log.Printf("下载插件失败: %v", err)
		} else {
			fmt.Printf("下载链接: %s\n", downloadResp.DownloadUrl)
		}
	} else {
		fmt.Printf("插件注册失败: %s\n", registerResp.Message)
	}
}
