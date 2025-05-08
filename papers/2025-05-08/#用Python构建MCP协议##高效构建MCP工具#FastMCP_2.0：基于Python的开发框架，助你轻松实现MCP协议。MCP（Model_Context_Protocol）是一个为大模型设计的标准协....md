# #用Python构建MCP协议##高效构建MCP工具#FastMCP 2.0：基于Python的开发框架，助你轻松实现MCP协议。MCP（Model Context Protocol）是一个为大模型设计的标准协...

**URL**: https://weibo.com/6105753431/PqSCS4ksD

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E7%94%A8Python%E6%9E%84%E5%BB%BAMCP%E5%8D%8F%E8%AE%AE%23&amp;extparam=%23%E7%94%A8Python%E6%9E%84%E5%BB%BAMCP%E5%8D%8F%E8%AE%AE%23" data-hide=""><span class="surl-text">#用Python构建MCP协议#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E9%AB%98%E6%95%88%E6%9E%84%E5%BB%BAMCP%E5%B7%A5%E5%85%B7%23&amp;extparam=%23%E9%AB%98%E6%95%88%E6%9E%84%E5%BB%BAMCP%E5%B7%A5%E5%85%B7%23" data-hide=""><span class="surl-text">#高效构建MCP工具#</span></a><br><br>FastMCP 2.0：基于Python的开发框架，助你轻松实现MCP协议。<br><br>MCP（Model Context Protocol）是一个为大模型设计的标准协议，目的是让模型可以安全、统一地访问外部数据和功能，有点像「AI 的 USB-C 接口」。<br><br>而MCP本身偏底层，实现起来不算轻松。FastMCP用Python封装了一套高层接口，帮你快速搭MCP服务端或客户端。<br><br>这次2.0版本，是FastMCP的全新迭代，核心亮点有：<br><br>客户端能力上线：不仅能写server，现在也能轻松写MCP client，可以用代码直接访问远程MCP服务，也可以用本地in-memory的方式无缝调试。<br><br>服务代理能力更强：支持构建proxy server，把一个远程服务封装成本地服务用，也支持多个MCP服务合并成一个复合服务，适合做插件系统或多功能助手。<br><br>模块接口自动生成：可以从现有的OpenAPI或FastAPI项目直接生成MCP服务，接入老项目几乎零改动，兼容性好。<br><br>开发体验更轻量：服务定义用Python装饰器就能搞定，不用再纠结协议细节、请求格式或权限处理，适合原型验证和小团队开发。<br><br>官方文档和示例很全面，几乎覆盖了从“Hello World”到线上部署的整个流程。<br><br>写法也很 Pythonic，比如：<br><br>`<a href="https://weibo.com/n/mcp">@mcp</a>.tool() def add(a: int, b: int) -&gt; int:     return a + b`<br><br>装饰器一加，LLM就能调用这个函数。<br><br>对熟悉大模型工具的开发人员来说，这套框架既快又干净，非常适合做快速原型，也能接到正式服务里。<br><br>项目地址：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fgithub.com%2Fjlowin%2Ffastmcp" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>文档站：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fgofastmcp.com%2F" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i185ss7pdoj31p21cs1kx.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

FastMCP 2.0是一个基于Python的MCP（Model Context Protocol）协议开发框架，旨在简化大模型与外部数据/功能的交互。新版本支持客户端和服务端开发，提供代理服务、多服务合并、模块接口自动生成（兼容OpenAPI/FastAPI）等功能。通过Python装饰器即可快速定义服务（如`@mcp.tool()`），无需处理底层协议细节，适合快速原型开发和部署。该项目文档完善，具有Pythonic风格，显著降低了MCP协议的实现门槛。GitHub地址和文档站已公开。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-08T19:03:33Z
- **目录日期**: 2025-05-08
