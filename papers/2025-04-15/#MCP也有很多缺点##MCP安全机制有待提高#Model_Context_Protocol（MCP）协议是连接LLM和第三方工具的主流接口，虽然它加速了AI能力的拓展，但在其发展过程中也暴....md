# #MCP也有很多缺点##MCP安全机制有待提高#Model Context Protocol（MCP）协议是连接LLM和第三方工具的主流接口，虽然它加速了AI能力的拓展，但在其发展过程中也暴...

**URL**: https://weibo.com/6105753431/Pnop5s9JX

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23MCP%E4%B9%9F%E6%9C%89%E5%BE%88%E5%A4%9A%E7%BC%BA%E7%82%B9%23&amp;extparam=%23MCP%E4%B9%9F%E6%9C%89%E5%BE%88%E5%A4%9A%E7%BC%BA%E7%82%B9%23" data-hide=""><span class="surl-text">#MCP也有很多缺点#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23MCP%E5%AE%89%E5%85%A8%E6%9C%BA%E5%88%B6%E6%9C%89%E5%BE%85%E6%8F%90%E9%AB%98%23&amp;extparam=%23MCP%E5%AE%89%E5%85%A8%E6%9C%BA%E5%88%B6%E6%9C%89%E5%BE%85%E6%8F%90%E9%AB%98%23" data-hide=""><span class="surl-text">#MCP安全机制有待提高#</span></a><br><br>Model Context Protocol（MCP）协议是连接LLM和第三方工具的主流接口，虽然它加速了AI能力的拓展，但在其发展过程中也暴露出不少隐患。<br><br>开发者Shrivu Shankar的文章写道，MCP能让模型可以自主调用调试代码、发社交动态、控制设备等操作。<br><br>然而，随着使用场景的增加，几个关键风险也逐渐浮现：<br><br>1、安全机制薄弱：MCP第一版规范并未统一认证机制，导致各个MCP服务端自行实现认证流程，甚至出现完全不校验权限的情况。<br><br>这不仅增加数据泄露的可能，还可能被恶意方利用，形成“第四方攻击”——即攻击者通过MCP工具所信任的第三方数据源植入恶意内容，绕过防护。<br><br>2、执行环境过于开放：MCP支持本地运行工具服务，这使得接入第三方代码门槛极低，几乎等于“默认信任”。如果工具开发者在代码中隐藏异常逻辑，模型执行时用户根本无法察觉，一旦被触发，后果难以控制。<br><br>3、人机边界模糊：工具名称可以被随意修改、调用过程无需用户确认，且系统权限划分不明确，导致模型“误操作”的风险大大增加。比如不经确认就删除文件、私自发布内容等，都可能在用户毫不知情的情况下发生，像是在YOLO模式下把电脑交给一个“理解力有限”的智能体。<br><br>4、数据传输开销大：MCP接口大多以纯文本方式传递数据，信息传输成本高，延迟明显。一些Cursor用户已在Reddit上吐槽相关账单无端上涨，说明这一问题对普通用户也产生了真实影响。<br><br>5、多工具环境性能降低：当LLM需要处理极长的上下文、多源数据和大量工具时，反而更容易“胡说八道”。智能助手越复杂，出错的可能越高，尤其在用户不了解工具运作机制的情况下，错误更难被发现和纠正。<br><br>6、敏感数据泄露风险高：MCP的便利性让模型能自动访问和处理大量用户数据。例如，用户只是想“连上Google Drive写点东西”，模型却可能读取了医疗记录，还“贴心”地自动发布了出去。<br><br>作者表示，MCP协议本身并无问题，它统一了AI调用外部能力的接口，让开发者能快速扩展模型功能。<br><br>协议应以安全为默认前提，应用需清晰设定权限边界，用户也必须理解模型真正能做什么、不能做什么。<br><br>否则，再强大的AI助手，也可能在“连上数据”的那一刻，变成一个不可控的Bug制造机。<br><br>感兴趣的小伙伴可以点击：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fblog.sshh.io%2Fp%2Feverything-wrong-with-mcp" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i0hmmicszqj317k18a4eg.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

MCP协议作为连接大语言模型和第三方工具的主流接口，虽然加速了AI能力拓展，但存在多个安全隐患：1) 安全机制薄弱，缺乏统一认证；2) 执行环境过于开放，易被植入恶意代码；3) 人机边界模糊，存在误操作风险；4) 数据传输效率低；5) 多工具环境下性能下降；6) 敏感数据易泄露。作者指出MCP协议本身有价值，但需加强默认安全设置，明确权限边界，并提高用户对模型能力的认知，否则可能成为不可控的风险源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-15T21:03:52Z
- **目录日期**: 2025-04-15
