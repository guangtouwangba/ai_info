# #一行代码让iPhone变砖##一行代码让iPhone无限重启#只需一行代码，就能让iPhone变砖，进入无限重启状态？【图1】这是安全员Rambo曝光的一个iOS漏洞，问题出在一...

**URL**: https://weibo.com/6105753431/PptvvjEHu

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E4%B8%80%E8%A1%8C%E4%BB%A3%E7%A0%81%E8%AE%A9iPhone%E5%8F%98%E7%A0%96%23&amp;extparam=%23%E4%B8%80%E8%A1%8C%E4%BB%A3%E7%A0%81%E8%AE%A9iPhone%E5%8F%98%E7%A0%96%23" data-hide=""><span class="surl-text">#一行代码让iPhone变砖#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E4%B8%80%E8%A1%8C%E4%BB%A3%E7%A0%81%E8%AE%A9iPhone%E6%97%A0%E9%99%90%E9%87%8D%E5%90%AF%23&amp;extparam=%23%E4%B8%80%E8%A1%8C%E4%BB%A3%E7%A0%81%E8%AE%A9iPhone%E6%97%A0%E9%99%90%E9%87%8D%E5%90%AF%23" data-hide=""><span class="surl-text">#一行代码让iPhone无限重启#</span></a><br><br>只需一行代码，就能让iPhone变砖，进入无限重启状态？【图1】<br><br>这是安全员Rambo曝光的一个iOS漏洞，问题出在一个冷门且历史悠久的系统API——Darwin Notifications。<br><br>Darwin Notifications是苹果系统内部用于进程间通信的机制，特点是：<br><br>- 不需要特殊权限即可发送消息；<br>- 没有发送者验证机制；<br>- 公开可用，第三方App也能调用。<br><br>正常来说，这些通知是系统内部自己用的，比如告诉桌面“屏幕要锁了”。但因为没有认证机制，任何App都能伪造它们。关键代码就一行：<br><br>notify_post("com.apple.MobileSync.BackupAgent.RestoreStarted")<br><br>为演示漏洞威力，Rambo开发了一个名为VeryEvilNotify的小型App，安装后，手机会不断弹出恢复提示，重启无效，只能彻底刷机才能恢复。【图2】<br><br>好消息是，Apple已在iOS 18.3中修复了这个问题。现在，发送敏感Darwin通知需要特殊权限（entitlements），普通App无法随意调用了。<br><br>该事件已被官方加入了漏洞库，代号CVE-2025-24091，Rambo还获得了1.75万美元的赏金。<br><br>感兴趣的小伙伴可以点击原文：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Frambo.codes%2Fposts%2F2025-04-24-how-a-single-line-of-code-could-brick-your-iphone" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i0xh6s8yajg30ik0acu11.gif" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i0xh6qpw7og306u0diwki.gif" referrerpolicy="no-referrer"><br><br>

## AI 摘要

安全研究员Rambo发现iOS存在高危漏洞，通过Darwin Notifications系统API（无权限验证机制），仅需一行代码"notify_post("com.apple.MobileSync.BackupAgent.RestoreStarted")"即可导致iPhone无限重启变砖。该漏洞允许任意App发送伪造系统通知，其开发的VeryEvilNotify应用可触发持续恢复提示，需刷机才能修复。苹果已在iOS 18.3修复该漏洞（CVE-2025-24091），新增权限限制，并向发现者支付1.75万美元赏金。该API此前存在三大风险点：免权限调用、无发送者验证、第三方应用可用。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-29T07:03:45Z
- **目录日期**: 2025-04-29
