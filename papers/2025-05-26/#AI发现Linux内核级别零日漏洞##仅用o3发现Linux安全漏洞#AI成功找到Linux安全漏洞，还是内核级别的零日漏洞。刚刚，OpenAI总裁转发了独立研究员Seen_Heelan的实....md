# #AI发现Linux内核级别零日漏洞##仅用o3发现Linux安全漏洞#AI成功找到Linux安全漏洞，还是内核级别的零日漏洞。刚刚，OpenAI总裁转发了独立研究员Seen Heelan的实...

**URL**: https://weibo.com/6105753431/Pts6y3L0V

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23AI%E5%8F%91%E7%8E%B0Linux%E5%86%85%E6%A0%B8%E7%BA%A7%E5%88%AB%E9%9B%B6%E6%97%A5%E6%BC%8F%E6%B4%9E%23&amp;extparam=%23AI%E5%8F%91%E7%8E%B0Linux%E5%86%85%E6%A0%B8%E7%BA%A7%E5%88%AB%E9%9B%B6%E6%97%A5%E6%BC%8F%E6%B4%9E%23" data-hide=""><span class="surl-text">#AI发现Linux内核级别零日漏洞#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E4%BB%85%E7%94%A8o3%E5%8F%91%E7%8E%B0Linux%E5%AE%89%E5%85%A8%E6%BC%8F%E6%B4%9E%23&amp;extparam=%23%E4%BB%85%E7%94%A8o3%E5%8F%91%E7%8E%B0Linux%E5%AE%89%E5%85%A8%E6%BC%8F%E6%B4%9E%23" data-hide=""><span class="surl-text">#仅用o3发现Linux安全漏洞#</span></a><br><br>AI成功找到Linux安全漏洞，还是内核级别的零日漏洞。<br><br>刚刚，OpenAI总裁转发了独立研究员Seen Heelan的实验成果：用o3模型找到了Linux内核SMB实现中的一个远程零日漏洞。【图1】<br><br>更让人惊讶的是，整个过程中没有用到任何复杂的工具——没有脚手架、没有智能体框架、没有工具调用，仅仅是o3 API本身。<br><br>这个漏洞被编号为CVE-2025-37899，是SMB”注销”命令处理程序中的一个释放后使用（use-after-free）漏洞。<br><br>据作者透露，这是首次公开讨论的由大模型发现的此类漏洞。【图2】<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FJ_oaAnGjfiDUyv7jxTvZ3w" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">历史首次！o3找到Linux内核零日漏洞，12000行代码看100遍揪出，无需调用任何工具</span></a><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3ly1i1rptbzgghj30xa0kaq8t.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3ly1i1rpth4kg4j31lo0jqq93.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

AI首次独立发现Linux内核级零日漏洞（CVE-2025-37899）。研究员Seen Heelan仅使用o3模型API（无需复杂工具）在Linux内核SMB实现中识别出一个释放后使用漏洞，涉及"注销"命令处理程序。该漏洞通过分析12000行代码被发现，是首个公开报道由大模型发现的重大安全漏洞。OpenAI总裁转发了这一突破性成果，突显了AI在代码审计领域的潜力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-26T04:06:47Z
- **目录日期**: 2025-05-26
