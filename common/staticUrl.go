package common

var (
	// SYUrl = "<a href=\"([^\"]+/)\" title=\"([^\"]+)\">([^\"]+)</a>"
	SYUrl = "<li><a href=\"(.*)\" title=\"(.*)\">(.*)</a></li>"
	TypeInfoUrl = "<option value='(.*html)' (.*)>(.*)</option>"
	BookInfoRule = `<div class="s">作者：(.*)<br />大小：(.*)MB<br>等级：<em class="lstar(.*)"></em><br>更新：(.*)</div>`
	BookInfoRuleAll = `
<div class="s">作者：(.*)<br />大小：(.*)MB<br>等级：<em class="lstar(.*)"></em><br>更新：(.*)</div>
    <a href="(.*)"><img src="(.*).jpg" onerror="this.src='(.*).jpg'">《(.*)》全集</a>
`
	TBookInfoRuleAll =  `<div class="s">作者：(.*)<br />大小：(.*)MB<br>等级：<em class="lstar(.*)"></em><br>更新：(.*)</div>(.*)
    <a href="(.*)"><img src="(.*).jpg" onerror="this.src='(.*).jpg'">《(.*)》全集</a>`
	// 因为获取的信息不是完整的Url，所以需要拼接
	SplicingUrl = "https://www.kankezw.com"
	BookRule = `<li>
    <div class="s">作者：(.*)<br />大小：(.*)MB<br>等级：<em class="lstar(.*)"></em><br>更新：(.*)</div>
    <a href="(.*)"><img src="(.*).jpg" onerror="this.src='(.*).jpg'">《(.*)》全集</a>
    <div class="u"> (.*)...</div>
    <div><a style="font-weight: normal;" href="(.*)">最新章节：(.*)</a></div>
</li>`
)