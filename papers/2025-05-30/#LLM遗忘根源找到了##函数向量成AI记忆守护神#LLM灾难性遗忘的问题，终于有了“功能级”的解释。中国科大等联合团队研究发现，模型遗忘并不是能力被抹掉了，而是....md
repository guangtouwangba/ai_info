# #LLM遗忘根源找到了##函数向量成AI记忆守护神#LLM灾难性遗忘的问题，终于有了“功能级”的解释。中国科大等联合团队研究发现，模型遗忘并不是能力被抹掉了，而是...

**URL**: https://weibo.com/6105753431/PudJn88ms

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23LLM%E9%81%97%E5%BF%98%E6%A0%B9%E6%BA%90%E6%89%BE%E5%88%B0%E4%BA%86%23&amp;extparam=%23LLM%E9%81%97%E5%BF%98%E6%A0%B9%E6%BA%90%E6%89%BE%E5%88%B0%E4%BA%86%23" data-hide=""><span class="surl-text">#LLM遗忘根源找到了#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%87%BD%E6%95%B0%E5%90%91%E9%87%8F%E6%88%90AI%E8%AE%B0%E5%BF%86%E5%AE%88%E6%8A%A4%E7%A5%9E%23&amp;extparam=%23%E5%87%BD%E6%95%B0%E5%90%91%E9%87%8F%E6%88%90AI%E8%AE%B0%E5%BF%86%E5%AE%88%E6%8A%A4%E7%A5%9E%23" data-hide=""><span class="surl-text">#函数向量成AI记忆守护神#</span></a><br><br>LLM灾难性遗忘的问题，终于有了“功能级”的解释。<br><br>中国科大等联合团队研究发现，模型遗忘并不是能力被抹掉了，而是新任务激活了“偏”的功能，导致原有能力失效。换句话说，模型没真忘，只是没用对地方。<br><br>为此，团队引入了“函数向量”分析方法，把模型对特定任务的功能抽象成向量，并跟踪其变化，发现功能向量的偏移正是遗忘的核心症状。<br><br>实验还发现：遗忘现象有模型差异，生成任务比分类任务更易遗忘，且早期的性能下降可能是“假性”——后续会恢复。<br><br>研究进一步提出了一种函数向量引导训练法（FVG），通过限制功能激活的偏移来保护模型记忆，不仅减轻遗忘，还增强上下文学习能力。论文已被ICLR 2025 oral接收，代码已开源。 <a href="https://weibo.com/ttarticle/p/show?id=2309405172026255212626" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">函数向量对齐技术，让大模型持续学习不“失忆”丨ICLR 2025</span></a><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i1xk43cdeej30kc0bgta1.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

中国科大团队研究发现，大语言模型(LLM)的"遗忘"并非能力丢失，而是新任务激活了错误功能区域。研究通过"函数向量"分析，将任务功能抽象为向量并追踪其偏移，揭示了遗忘机制。实验显示：生成任务比分类任务更易遗忘，早期性能下降可能是暂时现象。团队提出函数向量引导训练法(FVG)，通过控制功能偏移来保护记忆，不仅减少遗忘还提升上下文学习能力。该成果获ICLR 2025 oral接收并开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-30T18:03:57Z
- **目录日期**: 2025-05-30
