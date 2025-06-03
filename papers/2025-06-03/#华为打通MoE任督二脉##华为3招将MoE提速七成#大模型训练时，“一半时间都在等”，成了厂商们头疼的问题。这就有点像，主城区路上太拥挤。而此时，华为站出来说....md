# #华为打通MoE任督二脉##华为3招将MoE提速七成#大模型训练时，“一半时间都在等”，成了厂商们头疼的问题。这就有点像，主城区路上太拥挤。而此时，华为站出来说...

**URL**: https://weibo.com/6105753431/PuP69gLMe

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%8D%8E%E4%B8%BA%E6%89%93%E9%80%9AMoE%E4%BB%BB%E7%9D%A3%E4%BA%8C%E8%84%89%23&amp;extparam=%23%E5%8D%8E%E4%B8%BA%E6%89%93%E9%80%9AMoE%E4%BB%BB%E7%9D%A3%E4%BA%8C%E8%84%89%23" data-hide=""><span class="surl-text">#华为打通MoE任督二脉#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%8D%8E%E4%B8%BA3%E6%8B%9B%E5%B0%86MoE%E6%8F%90%E9%80%9F%E4%B8%83%E6%88%90%23&amp;extparam=%23%E5%8D%8E%E4%B8%BA3%E6%8B%9B%E5%B0%86MoE%E6%8F%90%E9%80%9F%E4%B8%83%E6%88%90%23" data-hide=""><span class="surl-text">#华为3招将MoE提速七成#</span></a><br><br>大模型训练时，“一半时间都在等”，成了厂商们头疼的问题。<br><br>这就有点像，主城区路上太拥挤。而此时，华为站出来说，我要建地下通道——他们推出了Adaptive Pipe &amp; EDPB优化方案，专治训练卡顿。<br><br>在了解他们具体怎么做的之前，咱们先来回顾一下技术背景。<br><br>大模型Scaling Law的利器——MoE（混合专家），训练难点主要有两块：<br><br>- 专家并行、All-to-All通信：机器间拉扯，计算老在等通信。<br><br>- 负载不均：热门专家被塞爆，冷门专家闲着；不同层计算量差距大。<br><br>于是，华为像建地下通道一样，将训练集群优化了：<br><br>- 通信掩盖技术（Adaptive Pipe）：把“通信”藏到“计算”中，几乎不让计算等通信，提升了带宽利用。<br><br>- 智能路由+专家动态迁移（EDPB）：热专家冷专家自动调配，负载均衡，让每台机器都不闲着。<br><br>这套系统依托DeployMind仿真平台，仅1小时就能模拟百万训练场景，选出最优并行策略。实测中，Adaptive Pipe可实现98%以上的EP通信掩盖，让计算引擎不受通信等待的束缚。详细请看：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FkOLEdpPDALM4IDu-sPXrVg" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">训练MoE足足提速70%！华为只用了3招</span></a><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i2252vqfrpj30sg0sg4qp.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i22538lmgcj30zk0kpti4.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

华为针对MoE(混合专家)大模型训练中的通信等待和负载不均问题，提出创新优化方案：1) Adaptive Pipe通信掩盖技术，将通信隐藏在计算中，提升带宽利用率；2) EDPB智能路由与专家动态迁移，实现负载均衡。通过DeployMind仿真平台快速优化并行策略，实测通信掩盖率达98%以上，整体训练速度提升70%。该方案有效解决了专家并行中的计算等待和资源分配不均问题。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-03T08:03:45Z
- **目录日期**: 2025-06-03
