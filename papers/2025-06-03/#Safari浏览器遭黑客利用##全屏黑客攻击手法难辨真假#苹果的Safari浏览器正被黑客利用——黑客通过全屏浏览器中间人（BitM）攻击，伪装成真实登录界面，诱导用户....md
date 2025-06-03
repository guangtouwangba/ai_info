# #Safari浏览器遭黑客利用##全屏黑客攻击手法难辨真假#苹果的Safari浏览器正被黑客利用——黑客通过全屏浏览器中间人（BitM）攻击，伪装成真实登录界面，诱导用户...

**URL**: https://weibo.com/6105753431/PuPt9jq1t

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23Safari%E6%B5%8F%E8%A7%88%E5%99%A8%E9%81%AD%E9%BB%91%E5%AE%A2%E5%88%A9%E7%94%A8%23&amp;extparam=%23Safari%E6%B5%8F%E8%A7%88%E5%99%A8%E9%81%AD%E9%BB%91%E5%AE%A2%E5%88%A9%E7%94%A8%23" data-hide=""><span class="surl-text">#Safari浏览器遭黑客利用#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%85%A8%E5%B1%8F%E9%BB%91%E5%AE%A2%E6%94%BB%E5%87%BB%E6%89%8B%E6%B3%95%E9%9A%BE%E8%BE%A8%E7%9C%9F%E5%81%87%23&amp;extparam=%23%E5%85%A8%E5%B1%8F%E9%BB%91%E5%AE%A2%E6%94%BB%E5%87%BB%E6%89%8B%E6%B3%95%E9%9A%BE%E8%BE%A8%E7%9C%9F%E5%81%87%23" data-hide=""><span class="surl-text">#全屏黑客攻击手法难辨真假#</span></a><br><br>苹果的Safari浏览器正被黑客利用——<br><br>黑客通过全屏浏览器中间人（BitM）攻击，伪装成真实登录界面，诱导用户输入账号密码，已经有爱用全屏浏览模式的网友中招了！<br><br>该攻击的核心就是，全屏显示的浏览器窗口，看不见地址栏，从而用假网址欺骗用户。<br><br>具体操作方式是这样的：<br><br>- 攻击者诱导用户点击伪造链接（例如伪装成Figma或Steam的登录页面）；<br><br>- 一旦用户进入该页面并点击“登录”，隐藏的黑客浏览器窗口会被激活；<br><br>- 攻击者利用Fullscreen API让这个窗口进入全屏模式，遮住浏览器地址栏；<br><br>- Safari浏览器在进入全屏时，唯一的提示仅是一段容易被忽略的动画，这使得攻击更加隐蔽。<br><br>更麻烦的是，这种攻击不会触发大多数安全软件的预警。即使在Chrome或Firefox中，这类攻击也能运行，但这些浏览器至少会弹出“正在进入全屏”的提示，而Safari的窗口是滑动动画，非常容易被忽略。<br><br>SquareX团队已向Apple通报了该问题，但苹果方面回应为“wontfix”，理由是“动画已经足够提示用户”。<br><br>目前，相关攻击已被观察到在社交媒体广告、评论区等渠道活跃传播。提醒Safari用户格外警惕，遇到需要登录的重要页面时，务必核对网址，不要轻易信任任何看似正常的界面。<img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i226q2yuwmj31ac0og13r.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i226q4m7htj30zk0b446i.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

Safari浏览器正遭遇新型全屏网络钓鱼攻击。黑客通过伪造登录页面，利用Fullscreen API隐藏地址栏，使用户难以辨别真假。攻击手法在社交媒体广告等渠道传播，诱导用户输入账号密码。相比Chrome等浏览器的明确提示，Safari仅以易被忽略的滑动动画作为全屏提示，安全风险更高。苹果公司以"动画已足够提示"为由拒绝修复该漏洞。专家建议用户登录时务必核对网址，警惕全屏页面。该攻击手法对Safari用户威胁较大，需提高防范意识。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-03T20:03:43Z
- **目录日期**: 2025-06-03
