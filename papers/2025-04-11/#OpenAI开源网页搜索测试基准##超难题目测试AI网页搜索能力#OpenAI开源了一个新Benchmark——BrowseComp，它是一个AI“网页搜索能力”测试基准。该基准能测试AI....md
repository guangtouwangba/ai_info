# #OpenAI开源网页搜索测试基准##超难题目测试AI网页搜索能力#OpenAI开源了一个新Benchmark——BrowseComp，它是一个AI“网页搜索能力”测试基准。该基准能测试AI...

**URL**: https://weibo.com/6105753431/PmLdVwQ0T

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23OpenAI%E5%BC%80%E6%BA%90%E7%BD%91%E9%A1%B5%E6%90%9C%E7%B4%A2%E6%B5%8B%E8%AF%95%E5%9F%BA%E5%87%86%23&amp;extparam=%23OpenAI%E5%BC%80%E6%BA%90%E7%BD%91%E9%A1%B5%E6%90%9C%E7%B4%A2%E6%B5%8B%E8%AF%95%E5%9F%BA%E5%87%86%23" data-hide=""><span class="surl-text">#OpenAI开源网页搜索测试基准#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E8%B6%85%E9%9A%BE%E9%A2%98%E7%9B%AE%E6%B5%8B%E8%AF%95AI%E7%BD%91%E9%A1%B5%E6%90%9C%E7%B4%A2%E8%83%BD%E5%8A%9B%23&amp;extparam=%23%E8%B6%85%E9%9A%BE%E9%A2%98%E7%9B%AE%E6%B5%8B%E8%AF%95AI%E7%BD%91%E9%A1%B5%E6%90%9C%E7%B4%A2%E8%83%BD%E5%8A%9B%23" data-hide=""><span class="surl-text">#超难题目测试AI网页搜索能力#</span></a><br><br>OpenAI开源了一个新Benchmark——BrowseComp，它是一个AI“网页搜索能力”测试基准。<br><br>该基准能测试AI模型的联网能力，评估模型网络搜索、筛选和信息整合能力，看看AI能否“找得准、找得全”。<br><br>BrowseComp中的题目覆盖范围极广，涉及电影、音乐、体育、历史、科技等多个领域，就是为了拟高难度的信息检索任务。【图1】<br><br>让我们看看BrowseComp中的具体例题，感受一下“高难度信息挖掘”场景：<br><br>案例1: 请找出一个虚构角色，他很幽默，会和观众打破第四面墙互动，有过与修行者相关的背景设定，并出演了一部1960年代至1980年代之间播出、集数不到50集的电视节目。<br><br>正确答案：Plastic Man<br><br>这个题目关键词杂乱，时间范围宽泛，还需要比对老剧集中的角色设定。AI必须综合剧情资料与历史信息，极具挑战性。<br><br>案例2: 请找出一篇在2023年6月前发表的学术论文，它提及了文化传统、科学过程和烹饪方法。作者共三人，其中一位曾是西孟加拉邦的助理教授，另一位持有博士学位。<br><br>正确答案：The Fundamentals of Bread Making: The Science of Bread<br><br>此题横跨人文与科学，线索包含具体时间、学术背景与作者身份，需要模型使用学术搜索引擎进行深度检索，复杂程度直逼“炼狱模式”。<br><br>案例3: 请找出一位作家兼传记作者的笔名，这人写了多本书，包括自己的自传，并在1980年写过父亲的传记。这位作家曾在1940年代离婚并再婚，曾爱上某位哲学家的兄弟，而该哲学家在家中排行第八。<br><br>正确答案：Esther Wyndham<br><br>该题综合了人物情感经历、时间线、家族关系和出版信息，是一道典型的拼图式问题。稍有偏差，搜索方向就会完全跑偏。<br><br>由此，BrowseComp测试集具有以下特点——<br><br>- 任务设计精巧：“难找但可验证”。每道题都设有标准答案，避免开放式解释，确保评估的客观性与一致性。<br>- 搜索引擎难以直接解决：题目均由训练师严格验证，确保无法通过搜索引擎第一页轻松查到答案。<br>- 连人类都头疼：训练师解题成功率仅29.2%，超70%的题目就算查两个小时也可能无果。<br><br>OpenAI还借此测试集评估了自家模型的实际搜索能力，从结果来看【图2】，有两点值得关注：<br><br>- 通用联网模型表现普遍一般，即使具备搜索功能，在面对复杂题目时依然难以给出准确答案；<br>- 不具联网功能但推理能力更强的o1，反倒取得更好的成绩，说明推理对复杂检索任务同样关键。<br><br>而表现亮眼的，当然是Deep Research，它专门为“上网找资料”训练，并支持多轮采样和灵活调整搜索策略。<br><br>当一次搜索失败时，Deep Research不会死磕关键词，而是主动尝试替代方案，从而显著提升了搜索的准确率。<br><br>目前，BrowseComp的数据集和论文都已开源，感兴趣的小伙伴可以点击：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fgithub.com%2Fopenai%2Fevals" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i0cu9a6o42j30rg0q2adu.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i0cu9b84x8j30wy0gwwfo.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

OpenAI开源了网页搜索测试基准BrowseComp，用于评估AI模型的联网搜索和信息整合能力。该基准包含多领域高难度题目，如查找特定虚构角色、学术论文或作家笔名，要求综合复杂线索进行深度检索。题目设计确保无法通过简单搜索解决，人类解题成功率仅29.2%。测试显示，专用搜索模型Deep Research表现最佳，因其支持多轮策略调整。通用联网模型表现一般，而强推理模型也有一定优势。数据集和论文已开源，推动AI搜索能力研究。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-11T08:03:27Z
- **目录日期**: 2025-04-11
