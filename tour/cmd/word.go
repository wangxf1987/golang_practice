package cmd

import (
	"github.com/spf13/cobra"
	"golang_practice/tour/internal/word"
	"log"
	"strings"
)

const (
	MODE_UPPER = iota +1
	MODE_LOWER
	MODE_UNDERSCORE_TO_UPPER_CAMELCALSS
	MODE_UNDERSCORE_TO_LOWER_CAMELCALSS
	MODE_CAMELCASE_TO_UNDERSCORE
)

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下：",
	"1: 全部单词转换为大写",
	"2: 全部单词转换为小写",
	"3: 下划线单词转换为大写驼峰单词",
	"4: 下划线单词转换为小写驼峰单词",
	"5: 驼峰单词转为下划线单词",
}, "\n")

var str string
var mode int8
var wordCmd = &cobra.Command{
	Use: "word",
	Short: "单词格式转换",
	Long: desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case MODE_UPPER:
			content = word.ToUpper(str)
		case MODE_LOWER:
			content = word.ToLower(str)
		case MODE_UNDERSCORE_TO_UPPER_CAMELCALSS:
			content = word.UnderscoreToUpperCamelClass(str)
		case MODE_UNDERSCORE_TO_LOWER_CAMELCALSS:
			content = word.UnderscoreToLowerCamelClass(str)
		case MODE_CAMELCASE_TO_UNDERSCORE:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该转换模式，请执行 help word 查看帮助文档")
		}

		log.Printf("输出结果: %s", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
}