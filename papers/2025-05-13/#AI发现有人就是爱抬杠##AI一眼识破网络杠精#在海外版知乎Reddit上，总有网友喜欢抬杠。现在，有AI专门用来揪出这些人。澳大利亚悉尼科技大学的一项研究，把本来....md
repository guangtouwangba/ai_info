# #AI发现有人就是爱抬杠##AI一眼识破网络杠精#在海外版知乎Reddit上，总有网友喜欢抬杠。现在，有AI专门用来揪出这些人。澳大利亚悉尼科技大学的一项研究，把本来...

**URL**: https://weibo.com/6105753431/PrDMjCag0

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23AI%E5%8F%91%E7%8E%B0%E6%9C%89%E4%BA%BA%E5%B0%B1%E6%98%AF%E7%88%B1%E6%8A%AC%E6%9D%A0%23&amp;extparam=%23AI%E5%8F%91%E7%8E%B0%E6%9C%89%E4%BA%BA%E5%B0%B1%E6%98%AF%E7%88%B1%E6%8A%AC%E6%9D%A0%23" data-hide=""><span class="surl-text">#AI发现有人就是爱抬杠#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23AI%E4%B8%80%E7%9C%BC%E8%AF%86%E7%A0%B4%E7%BD%91%E7%BB%9C%E6%9D%A0%E7%B2%BE%23&amp;extparam=%23AI%E4%B8%80%E7%9C%BC%E8%AF%86%E7%A0%B4%E7%BD%91%E7%BB%9C%E6%9D%A0%E7%B2%BE%23" data-hide=""><span class="surl-text">#AI一眼识破网络杠精#</span></a><br><br>在海外版知乎Reddit上，总有网友喜欢抬杠。现在，有AI专门用来揪出这些人。<br><br>澳大利亚悉尼科技大学的一项研究，把本来用于自动驾驶的AI技术——逆强化学习（Inverse Reinforcement Learning，IRL），用在了网络行为分析上。<br><br>研究团队扒了6年，足足590万条Reddit互动数据，终于搞清楚了哪些人最喜欢在网上抬杠。<br><br>具体研究过程，大致可以理解为三步：<br><br>1. 把Reddit用户当成“智能体”建模：Reddit用户的每一次评论、回帖、对话路径都被当作一种“行为选择”。研究团队设定了一个“12个状态+6种动作”的用户模型，模拟他们是如何决定发帖、回帖、反驳的。<br><br>2. 用IRL推断背后的动机策略：IRL根据行为，推测“这个人到底图啥”。通过观察每个用户在不同情境下的决策方式，学习他们的“策略”。比如有人总是第一时间跳出来反对那AI就能识别出这个人有“争议式”的倾向。<br><br>3. 计算行为相似度：AI最后将用户分成5类行为人格<br><br>- 起话题型（Thread Creators）<br>- 只评论第一条评论的（Root Only）<br>- 偏好第一条评论+选定回帖的（Root Favored）<br>- 均衡参与的（Balanced）<br>- 专爱抬杠的（Disagreers）<br><br>为了比较用户间的“行为相似度”，研究引入了一个新指标：SWKL对称加权KL散度，用来度量两个用户“行为策略”的差距。<br><br>再借助LLM，研究人员把每条评论分成“同意 / 中立 / 反对”三类，帮助识别哪些人是“爱抬杠”的人格。<br><br>研究发现，这种“抬杠人格”不是某个话题带出来的，而是受“版块”影响。<br><br>比如足球版块r/soccer和电竞版块r/leagueoflegends的用户，大多都爱站队、都爱对喷。<br><br>这项研究表明，即使在匿名环境下，用户说出的语言，也能被AI看出ta是“哪类人”。<br><br>感兴趣的小伙伴可以阅读原文：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Farxiv.org%2Fabs%2F2502.02943" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i1dxyu7u2uj316a11sak8.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i1dxyv56caj316c15qneu.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

澳大利亚悉尼科技大学研究团队利用自动驾驶领域的逆强化学习技术（IRL），分析了Reddit平台590万条互动数据，开发出能识别"网络杠精"的AI系统。通过建立用户行为模型，AI将用户分为5类人格（包括专爱抬杠的"Disagreers"），并采用SWKL指标量化行为差异。研究发现，抬杠行为与特定版块（如足球、电竞）高度相关，而非话题本身。该技术能通过语言模式在匿名环境中识别用户类型，相关论文已发布于arXiv平台。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-13T09:03:01Z
- **目录日期**: 2025-05-13
