package utils

import (
	"regexp"
	"strings"
)

// HTMLToPlainText 将HTML转换为纯文本
func HTMLToPlainText(html string) string {
	replacements := map[string]string{
		`<p[^>]*>`:      "",
		`</p>`:          "\n",
		`<h1[^>]*>`:     "",
		`</h1>`:         "\n",
		`<h2[^>]*>`:     "",
		`</h2>`:         "\n",
		`<h3[^>]*>`:     "",
		`</h3>`:         "\n",
		`<ul[^>]*>`:     "",
		`</ul>`:         "",
		`<li[^>]*>`:     "- ",
		`</li>`:         "\n",
		`<strong[^>]*>`: "**",
		`</strong>`:     "**",
		`<em[^>]*>`:     "*",
		`</em>`:         "*",
		`<a[^>]*>`:      "",
		`</a>`:          "",
		`<br[^>]*>`:     "\n",
	}

	// 执行替换
	for pattern, replacement := range replacements {
		re := regexp.MustCompile(pattern)
		html = re.ReplaceAllString(html, replacement)
	}

	// 清理多余的换行和空格
	plainText := strings.TrimSpace(html)
	plainText = strings.ReplaceAll(plainText, "\n\n", "\n")

	return plainText
}
