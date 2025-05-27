# #Claude4被诱导窃取个人隐私##GitHub官方MCP泄露私有代码库#被选为GitHub Copilot官方模型后，Claude 4直接被诱导出bug了！一家瑞士网络安全公司发现，GitHub官...

**URL**: https://weibo.com/6105753431/PtLjAdAED

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23Claude4%E8%A2%AB%E8%AF%B1%E5%AF%BC%E7%AA%83%E5%8F%96%E4%B8%AA%E4%BA%BA%E9%9A%90%E7%A7%81%23&amp;extparam=%23Claude4%E8%A2%AB%E8%AF%B1%E5%AF%BC%E7%AA%83%E5%8F%96%E4%B8%AA%E4%BA%BA%E9%9A%90%E7%A7%81%23" data-hide=""><span class="surl-text">#Claude4被诱导窃取个人隐私#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23GitHub%E5%AE%98%E6%96%B9MCP%E6%B3%84%E9%9C%B2%E7%A7%81%E6%9C%89%E4%BB%A3%E7%A0%81%E5%BA%93%23&amp;extparam=%23GitHub%E5%AE%98%E6%96%B9MCP%E6%B3%84%E9%9C%B2%E7%A7%81%E6%9C%89%E4%BB%A3%E7%A0%81%E5%BA%93%23" data-hide=""><span class="surl-text">#GitHub官方MCP泄露私有代码库#</span></a><br><br>被选为GitHub Copilot官方模型后，Claude 4直接被诱导出bug了！<br><br>一家瑞士网络安全公司发现，GitHub官方MCP服务器正在面临新型攻击——<br><br>通过在公共仓库的正常内容中隐藏恶意指令，可以诱导AI Agent自动将私有仓库的敏感数据泄露至公共仓库。【图1】<br><br>就是说，当用户使用集成了GitHub MCP的Claude 4 ，用户的私人敏感数据可能遭到泄露。【图2】<br><br>更可怕的是，GitLab Duo近期也曝出类似漏洞（由以色列安全服务商Legit Security披露），也是和提示注入及HTML注入相关，攻击者利用漏洞操控AI Agent，最终导致私有代码泄露。【图3】<br><br>瑞士的这家公司表示，这并非传统意义上的GitHub平台漏洞，而是AI Agent工作流的设计缺陷。<br><br>这也引发了人们关于MCP是否应该存在的讨论。【图4】<br><br>具体咋回事儿？下面详细展开：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FnnSMzqDHwLUB3WSl8aAPNw" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">Claude 4被诱导窃取个人隐私！GitHub官方MCP服务器安全漏洞曝光</span></a><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i1u2o0doo2j30zk0nxtj0.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i1u2o1r1mfj30qm0zknbh.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i1u2o4p2hzj30zk0e7792.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i1u2o6v19bj30xq0d87da.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

瑞士网络安全公司发现GitHub官方MCP服务器存在新型AI安全漏洞：攻击者通过在公共仓库植入恶意指令，可诱导集成Claude 4的AI Agent自动泄露用户私有仓库的敏感数据。类似漏洞也出现在GitLab Duo，攻击者利用提示注入和HTML注入操控AI导致代码泄露。该问题源于AI工作流设计缺陷，非传统平台漏洞，已引发关于MCP安全性的讨论。事件揭示了AI辅助工具在隐私保护方面的潜在风险，需重新评估其安全架构。（98字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-27T08:03:06Z
- **目录日期**: 2025-05-27
