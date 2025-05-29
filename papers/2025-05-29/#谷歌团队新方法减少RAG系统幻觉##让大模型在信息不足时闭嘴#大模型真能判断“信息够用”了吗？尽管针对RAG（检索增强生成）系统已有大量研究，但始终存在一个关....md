# #谷歌团队新方法减少RAG系统幻觉##让大模型在信息不足时闭嘴#大模型真能判断“信息够用”了吗？尽管针对RAG（检索增强生成）系统已有大量研究，但始终存在一个关...

**URL**: https://weibo.com/6105753431/Pu2At6J4Y

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E8%B0%B7%E6%AD%8C%E5%9B%A2%E9%98%9F%E6%96%B0%E6%96%B9%E6%B3%95%E5%87%8F%E5%B0%91RAG%E7%B3%BB%E7%BB%9F%E5%B9%BB%E8%A7%89%23&amp;extparam=%23%E8%B0%B7%E6%AD%8C%E5%9B%A2%E9%98%9F%E6%96%B0%E6%96%B9%E6%B3%95%E5%87%8F%E5%B0%91RAG%E7%B3%BB%E7%BB%9F%E5%B9%BB%E8%A7%89%23" data-hide=""><span class="surl-text">#谷歌团队新方法减少RAG系统幻觉#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E8%AE%A9%E5%A4%A7%E6%A8%A1%E5%9E%8B%E5%9C%A8%E4%BF%A1%E6%81%AF%E4%B8%8D%E8%B6%B3%E6%97%B6%E9%97%AD%E5%98%B4%23&amp;extparam=%23%E8%AE%A9%E5%A4%A7%E6%A8%A1%E5%9E%8B%E5%9C%A8%E4%BF%A1%E6%81%AF%E4%B8%8D%E8%B6%B3%E6%97%B6%E9%97%AD%E5%98%B4%23" data-hide=""><span class="surl-text">#让大模型在信息不足时闭嘴#</span></a><br><br>大模型真能判断“信息够用”了吗？<br><br>尽管针对RAG（检索增强生成）系统已有大量研究，但始终存在一个关键谜题：当大模型回答出错时，究竟是因为没用好给它的信息，还是因为信息本身就不够回答问题？<br><br>为了搞清楚这个问题，来自谷歌的研究团队提出了“充分上下文”的概念，来评估提供给大模型的信息是否足够回答问题。<br><br>他们开发了一个基于Gemini 1.5 Pro的自动标注器，用于标记实例是否具有充足上下文，准确率高达93%，无需真实答案即可进行大规模评估。【图1】<br><br>研究团队在对主流数据集（如HotPotQA、Musique）进行分析时发现，即使是人工筛选过或在“理想检索”的条件下，仍有甚至超过一半的问题缺乏充足上下文。【图2】<br><br>研究发现：【图3】<br><br>- 基线性能较高的大型模型（如Gemini 1.5 Pro、GPT-4o和Claude 3.5）在上下文充足时表现出色，但上下文不足时宁可编错答案也不愿说“不知道”。<br><br>- 基线性能较低的小型模型（如Mistral 3和Gemma 2）即便上下文充足，也频繁放弃回答或胡编乱造。<br><br>为了改进模型的表现，研究团队提出了一种“选择性生成”方法。<br><br>这种方法会结合模型自身的“自信程度”以及上下文充分性来做决定，如果模型不确定，或者信息不够，它就会选择不回答。<br><br>这种方法相当有效，使用选择性生成方法后，模型在回答问题时的正确率提高了2-10%。【图4】<br><br>论文地址：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Farxiv.org%2Fabs%2F2411.06037" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>代码仓库：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fgithub.com%2Fhljoren%2Fsufficientcontext" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i1w6x36ymoj30um0bowjc.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i1w6x5icm2j30uc0ky182.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i1w6x8fukej30gu09eq4f.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i1w6xb34u8j30va0n24a2.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

谷歌团队提出新方法减少RAG系统幻觉，通过"充分上下文"概念评估信息是否足够回答问题。他们开发了基于Gemini 1.5 Pro的自动标注器(准确率93%)，分析发现即使理想检索下，过半问题仍缺乏足够信息。研究发现大模型在信息不足时宁可编造也不说"不知道"，而小模型则频繁放弃或胡编。团队提出的"选择性生成"方法结合模型自信度和上下文充分性，使正确率提升2-10%。该方法能有效让模型在信息不足时保持沉默，减少错误回答。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-29T06:03:12Z
- **目录日期**: 2025-05-29
