package apis

import (
	"github.com/GrayOxygen/json-to-grpc/errors"
	"github.com/GrayOxygen/json-to-grpc/parser"
	"github.com/GrayOxygen/json-to-grpc/tree"
	"github.com/GrayOxygen/json-to-grpc/util"
	"github.com/golang-collections/collections/stack"
	"strings"

	"fmt"
)

var (
	//nestStrutPath = util.GetCurPath() + "/嵌套Struct"
	protName = "protName" //输出的struct名称(顶级)  TODO 支持自定义
	//左大括，成对出栈找到子struct
	protLeftStack = stack.New()
)

//nestStruct struct
func JSON2Proto(jsonStr string) (string, error) {
	util.Log.Printf("\n 传入的json为 %s   \n", jsonStr)
	//1，解析json为嵌套struct字符串
	nestStructStr, err := parser.JsonToNestStruct(jsonStr)

	if err != nil {
		fmt.Println("解析json失败：：：", err)
		return "", err
	}
	if nestStructStr == "" {
		fmt.Println("json解析失败，请检查格式是否正确，常见错误有：json中带有//注释，中英文的逗号等")
		return "", errors.New("json解析失败，请检查格式是否正确，常见错误有：json中带有//注释，中英文的逗号等")
	}
	if !strings.Contains(nestStructStr, "{") || ! strings.Contains(nestStructStr, "}") {
		return "", errors.New("json解析失败，请检查格式是否正确，常见错误有：json中带有//注释，中英文的逗号等")
	}

	root := &tree.TreeNode{}
	root.Level = 0 //层级为0，返回的数据才是真的tree，层级从1开始
	lineCount := 1
	nss := strings.Split(nestStructStr, "\n") //下标0是第一行，1是第二行...
	util.Log.Printf("\n 直接按行划分，数组为 %s   \n", nss)

	parser.ClearCache()
	root, _ = parser.DFS(nss, protLeftStack, lineCount, root)
	//打印树
	util.PrintTree(root)

	//3，遍历树，输出非嵌套struct到文件中
	res, err := parser.GenerateProto(root)
	util.Log.Printf("\n 返回的proto: \n %s   \n", res)
	if err != nil {
		fmt.Println("生成非嵌套文件失败：：：", err)
		return "", err
	}

	//格式化proto
	format_res := ""
	format_res_scanner := strings.Split(res, "\n") //下标0是第一行，1是第二行...
	for index := lineCount - 1; index < len(format_res_scanner); index++ {
		temp_line := ""
		if !strings.Contains(format_res_scanner[index], "message") && !strings.Contains(format_res_scanner[index], "}") {
			//缩进保持一致
			format_res_scanner[index] = strings.TrimLeft(format_res_scanner[index], " ")
			format_res_scanner[index] = strings.TrimLeft(format_res_scanner[index], "\t")
			temp_line = "	" + strings.TrimLeft(format_res_scanner[index], " ")
		}
		if strings.Contains(format_res_scanner[index], "message") {
			temp_line = format_res_scanner[index]
			temp_line = strings.Replace(temp_line, "\t", "", -1)
		}
		if strings.Contains(format_res_scanner[index], "}") {
			temp_line = format_res_scanner[index]
		}

		format_res += temp_line + "\n"
	}

	//加上proto定义
	format_res =`syntax = "proto3";
package protobuf;

`+format_res

	return format_res, nil

}
