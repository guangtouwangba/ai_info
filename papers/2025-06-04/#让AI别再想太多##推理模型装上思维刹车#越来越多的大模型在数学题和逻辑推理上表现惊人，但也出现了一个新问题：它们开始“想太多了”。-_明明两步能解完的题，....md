# #让AI别再想太多##推理模型装上思维刹车#越来越多的大模型在数学题和逻辑推理上表现惊人，但也出现了一个新问题：它们开始“想太多了”。- 明明两步能解完的题，...

**URL**: https://weibo.com/6105753431/PuPKhkqKa

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E8%AE%A9AI%E5%88%AB%E5%86%8D%E6%83%B3%E5%A4%AA%E5%A4%9A%23&amp;extparam=%23%E8%AE%A9AI%E5%88%AB%E5%86%8D%E6%83%B3%E5%A4%AA%E5%A4%9A%23" data-hide=""><span class="surl-text">#让AI别再想太多#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E6%8E%A8%E7%90%86%E6%A8%A1%E5%9E%8B%E8%A3%85%E4%B8%8A%E6%80%9D%E7%BB%B4%E5%88%B9%E8%BD%A6%23&amp;extparam=%23%E6%8E%A8%E7%90%86%E6%A8%A1%E5%9E%8B%E8%A3%85%E4%B8%8A%E6%80%9D%E7%BB%B4%E5%88%B9%E8%BD%A6%23" data-hide=""><span class="surl-text">#推理模型装上思维刹车#</span></a><br><br>越来越多的大模型在数学题和逻辑推理上表现惊人，但也出现了一个新问题：它们开始“想太多了”。<br><br>- 明明两步能解完的题，要绕七八步；<br><br>- 简单结论非得用复杂语言；<br><br>- 输出动辄上千tokens，拖慢速度还可能出错。<br><br>这类“过度思考”不仅浪费算力，还可能让答案跑偏。<br><br>为了给模型“踩刹车”，浙大、天大和MSRA团队提出了一种新框架：Self-Braking Tuning（SBT）。<br><br>它的目标很明确——让模型在该停的时候停下来，“少想一点”，但准确率不打折。<br><br>具体怎么做？<br><br>- 刹车信号机制：训练时加入提示，让模型学会判断“信息够了”，该停了；<br><br>- 多任务微调：模型不仅学会怎么答题，还学会“何时打住”；<br><br>- 自调节制动：对于推理中冗余的部分用mask掩盖，不纳入训练损失；<br><br>- 自然语言提醒：用“我是不是说太多了？”这类提示词，引导模型停下来。<br><br>为了衡量效果，团队还设计了两个指标：<br><br>1. 推理效率比：解出第一个答案用了多少步骤，占总步骤多少；<br><br>2. 过度推理标记比：分析模型语言中出现“Wait”“However”等提示“想太多了”的词频。<br><br>最终结果很亮眼——  <br><br>在多个数学推理基准测试中，SBT让模型的token输出减少了62.8%，准确率依然高达94.1%。<br><br>换句话说，AI终于学会了“不绕弯子”，想得少、答得好。<br><br>项目已开源，有兴趣的可以看看，项目主页：ZJU-REAL.github.io/SBT <a href="https://weibo.com/ttarticle/p/show?id=2309405173487538143363" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">推理“刹不住车”？新框架让DeepSeek-R1们告别过度思考，已开源</span></a><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i227xjwlv6j30np0dcgoj.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

浙江大学、天津大学和微软亚洲研究院团队提出Self-Braking Tuning（SBT）框架，解决大模型在数学推理中"过度思考"问题。该框架通过刹车信号机制、多任务微调等方法，训练模型在获得足够信息时自动停止推理。实验显示，SBT使模型输出token减少62.8%，同时保持94.1%的准确率。团队还设计了推理效率比和过度推理标记比两个评估指标。该项目已开源，有效提升了AI的推理效率和精准度。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-04T00:02:10Z
- **目录日期**: 2025-06-04
