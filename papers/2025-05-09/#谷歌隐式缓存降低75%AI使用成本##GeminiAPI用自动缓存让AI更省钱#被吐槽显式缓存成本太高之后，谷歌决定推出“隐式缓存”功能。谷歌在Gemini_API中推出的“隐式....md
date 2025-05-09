# #谷歌隐式缓存降低75%AI使用成本##GeminiAPI用自动缓存让AI更省钱#被吐槽显式缓存成本太高之后，谷歌决定推出“隐式缓存”功能。谷歌在Gemini API中推出的“隐式...

**URL**: https://weibo.com/6105753431/PqZZYCdrx

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E8%B0%B7%E6%AD%8C%E9%9A%90%E5%BC%8F%E7%BC%93%E5%AD%98%E9%99%8D%E4%BD%8E75%25AI%E4%BD%BF%E7%94%A8%E6%88%90%E6%9C%AC%23&amp;extparam=%23%E8%B0%B7%E6%AD%8C%E9%9A%90%E5%BC%8F%E7%BC%93%E5%AD%98%E9%99%8D%E4%BD%8E75%25AI%E4%BD%BF%E7%94%A8%E6%88%90%E6%9C%AC%23" data-hide=""><span class="surl-text">#谷歌隐式缓存降低75%AI使用成本#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23GeminiAPI%E7%94%A8%E8%87%AA%E5%8A%A8%E7%BC%93%E5%AD%98%E8%AE%A9AI%E6%9B%B4%E7%9C%81%E9%92%B1%23&amp;extparam=%23GeminiAPI%E7%94%A8%E8%87%AA%E5%8A%A8%E7%BC%93%E5%AD%98%E8%AE%A9AI%E6%9B%B4%E7%9C%81%E9%92%B1%23" data-hide=""><span class="surl-text">#GeminiAPI用自动缓存让AI更省钱#</span></a><br><br>被吐槽显式缓存成本太高之后，谷歌决定推出“隐式缓存”功能。<br><br>谷歌在Gemini API中推出的“隐式缓存”功能，适用于最新的Gemini 2.5 Pro和Gemini 2.5 Flash模型，能在“重复上下文”时节省75%的成本。<br><br>谷歌在官方博客中解释：“当你向Gemini 2.5系列模型发送请求时，如果该请求与之前的请求有共同前缀，就可能触发缓存命中，节省的成本会自动返还。”<br><br>触发隐式缓存的门槛较低：<br>- Gemini 2.5 Flash：至少1024个token<br>- Gemini 2.5 Pro：至少2048个token<br><br>缓存技术原本就是AI行业常用的降本手段，比如存储常见问题的答案，避免重复计算。<br><br>此前谷歌提供的显式缓存需要开发者手动设置高频提示词，而隐式缓存则是全自动运行，在Gemini 2.5系列模型中默认启用。<br><br>不过，TechCrunch等媒体也指出，鉴于谷歌此前关于缓存节省成本的承诺曾引发争议，这项新功能存在几点需要警惕之处：<br><br>1. 谷歌建议把重复内容放在请求开头，变动内容放末尾，以提高缓存命中率。<br>2. 目前没有第三方验证隐式缓存的实际节省效果，需观察早期用户反馈<br><br>因此，若开发者希望确保成本节约，仍可继续使用支持Gemini 2.5和2.0模型的显式缓存API。<img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i192cjidfaj318g0d8tfl.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i192clhokuj30wm0hmdls.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

谷歌在Gemini API中推出"隐式缓存"功能，适用于Gemini 2.5 Pro和2.5 Flash模型。该功能自动检测重复请求(最少1024/2048个token)，可节省75%成本。相比需要手动设置的显式缓存，隐式缓存默认开启且无需配置。但需注意：1)建议将重复内容放在请求开头；2)实际效果尚未经第三方验证。开发者仍可选择使用显式缓存API确保成本控制。这项技术通过避免重复计算来降低AI服务成本。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-09T03:19:50Z
- **目录日期**: 2025-05-09
