# #AI多轮对话易变人工智障##和AI对话最好一次说清楚#聊着聊着，大模型的智商像跳水一样越来越低？这不是你的错觉！来自微软研究院和Salesforce研究院的最新研究发...

**URL**: https://weibo.com/6105753431/PrUX6piLW

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23AI%E5%A4%9A%E8%BD%AE%E5%AF%B9%E8%AF%9D%E6%98%93%E5%8F%98%E4%BA%BA%E5%B7%A5%E6%99%BA%E9%9A%9C%23&amp;extparam=%23AI%E5%A4%9A%E8%BD%AE%E5%AF%B9%E8%AF%9D%E6%98%93%E5%8F%98%E4%BA%BA%E5%B7%A5%E6%99%BA%E9%9A%9C%23" data-hide=""><span class="surl-text">#AI多轮对话易变人工智障#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%92%8CAI%E5%AF%B9%E8%AF%9D%E6%9C%80%E5%A5%BD%E4%B8%80%E6%AC%A1%E8%AF%B4%E6%B8%85%E6%A5%9A%23&amp;extparam=%23%E5%92%8CAI%E5%AF%B9%E8%AF%9D%E6%9C%80%E5%A5%BD%E4%B8%80%E6%AC%A1%E8%AF%B4%E6%B8%85%E6%A5%9A%23" data-hide=""><span class="surl-text">#和AI对话最好一次说清楚#</span></a><br><br>聊着聊着，大模型的智商像跳水一样越来越低？这不是你的错觉！<br><br>来自微软研究院和Salesforce研究院的最新研究发现，大模型在多轮对话中的表现均显著低于单轮场景。【图2】<br><br>在六大生成任务（代码生成、数学解题、SQL编写、API调用、数据转文本和文档摘要）当中的平均性能下降甚至都能达到39%，即使是SoTA模型也不例外。【图3】<br><br>这是咋搞的呢？研究人员发现，这是由于LLMs常在对话早期做出假设，并过早尝试生成最终解决方案，且对此过度依赖。【图4】<br><br>再加上需要多轮对话的场景下，用户初始指令往往不完整，需要一步一步澄清。于是当我们意识到大模型在对话中偏离正确方向时，它们已经彻底迷失且无法自我修正。<br><br>研究人员将性能下降分解为两个因素：最佳情况下的性能下降和性能的波动范围。通过对20万+模拟对话的分析发现：【图5】<br><br>- 多轮对话的最佳情况下，模型的能力出现了15%小幅下降<br>- 多轮对话中，模型最优与最差回答的差距显著扩大，意味着其表现变得极不稳定。<br>- 单轮场景表现优异的模型，在多轮对话中与较小模型的不可靠程度相当。<br><br>这种性能下降很难通过“对话回顾”和“滚雪球”的策略缓解，降低生成随机性同样收效甚微。【图6】<br><br>不过，用户还是可以采用一些简单方法缓解，比如说：<br><br>- 尽量将所有需求整合到单次提示中，而非通过多轮对话逐步澄清。<br>- 当对话偏离正轨时，直接整合摘要开启新会话。<br><br>论文原文：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Farxiv.org%2Fabs%2F2505.06120" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i1g1rxyyzaj30zk0ogqe7.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i1g1rzhziij30zk0imk0b.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i1g1s1lgv7j30zk0hfto8.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i1g1s2jjv9j30zk0azdla.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i1g1s6b6clj30zk0luk5g.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i1g1s9l7lfj30ew0a4gow.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

微软和Salesforce的最新研究发现，大型语言模型（LLM）在多轮对话中的表现显著低于单轮场景，平均性能下降达39%。原因在于模型常在对话早期做出假设并过度依赖，而用户初始指令往往不完整，导致模型偏离方向后难以自我修正。多轮对话中模型表现不稳定，最优与最差回答差距扩大。缓解方法包括整合需求到单次提示中，或在对话偏离时开启新会话。研究强调多轮对话对模型性能的挑战，建议优化提示策略以提高效果。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-15T09:04:03Z
- **目录日期**: 2025-05-15
