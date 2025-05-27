# #扩散语言模型推理加速9倍##插件式dLLM推理加速器#扩散语言模型（dLLMs）火了，但一个现实问题是：推理过程迭代步数太多，影响了实际落地效率。针对这个问题，上...

**URL**: https://weibo.com/6105753431/PtL7jdqWT

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E6%89%A9%E6%95%A3%E8%AF%AD%E8%A8%80%E6%A8%A1%E5%9E%8B%E6%8E%A8%E7%90%86%E5%8A%A0%E9%80%9F9%E5%80%8D%23&amp;extparam=%23%E6%89%A9%E6%95%A3%E8%AF%AD%E8%A8%80%E6%A8%A1%E5%9E%8B%E6%8E%A8%E7%90%86%E5%8A%A0%E9%80%9F9%E5%80%8D%23" data-hide=""><span class="surl-text">#扩散语言模型推理加速9倍#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E6%8F%92%E4%BB%B6%E5%BC%8FdLLM%E6%8E%A8%E7%90%86%E5%8A%A0%E9%80%9F%E5%99%A8%23&amp;extparam=%23%E6%8F%92%E4%BB%B6%E5%BC%8FdLLM%E6%8E%A8%E7%90%86%E5%8A%A0%E9%80%9F%E5%99%A8%23" data-hide=""><span class="surl-text">#插件式dLLM推理加速器#</span></a><br><br>扩散语言模型（dLLMs）火了，但一个现实问题是：推理过程迭代步数太多，影响了实际落地效率。<br><br>针对这个问题，上交大EPIC Lab提出了一种完全免训练、即插即用的推理加速方法：dLLM-Cache。<br><br>亮点速览：<br><br>1. 无需训练、不改模型，即插即用；<br><br>2. 支持主流dLLM架构（LLaDA、Dream等）；<br><br>3. 在LLaDA 8B上实现最高9.1倍加速，生成质量几乎无损；<br><br>4. 让dLLMs的推理速度首次超过配有KV-Cache的LLaMA3。<br><br>如何判断哪些tokens“变了”？<br><br>作者提出了V-verify机制，通过比较当前步和前一步中Value向量的余弦相似度，来筛出“最不一样”的那一批（比如25%）进行更新。这个机制既精确又轻量，效果比传统的L2距离更好。<br><br>【图2】展示了dLLM-Cache的核心流程：<br><br>- 左侧是标准推理流程，所有token在每一层都重新计算；<br><br>- 右侧是启用缓存后的流程，黄色部分表示复用特征，蓝色表示需要更新的部分；<br><br>- V-verify模块则计算Value向量的余弦相似度，根据相似度低（变化大）来选择需要更新的tokens。<br><br>这项工作由两位本科生刘知远、杨奕存完成，在CVPR2025已有满分论文，背后指导老师是张林峰助理教授。<br><br>扩散模型通过多步去噪摆脱了自回归限制，能更好处理逆逻辑和长依赖，dLLM-Cache则让这条路线更具工程落地价值。<br><br>论文地址：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fgithub.com%2Fmaomaocun%2FdLLM-cache%2Fblob%2Fmain%2Fasset%2Fpaper.pdf" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>开源代码：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fgithub.com%2Fmaomaocun%2FdLLM-Cache" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>更多解读：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FMMVBTI6OHE2wUKLcTNAfmQ" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">扩散语言模型九倍推理加速！上海交大：KV Cache并非自回归模型的专属技巧</span></a><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i1u1mkiz8aj30wk0gwtgb.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i1u1u86tcbj30v50mkqes.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

上海交大EPIC Lab提出dLLM-Cache方法，显著加速扩散语言模型推理。该方案无需训练、即插即用，通过V-verify机制（基于Value向量余弦相似度）动态筛选25%变化最大的token更新，其余复用缓存。在LLaDA 8B模型上实现9.1倍加速且质量无损，推理速度首次超越带KV-Cache的LLaMA3。由本科生团队研发，已获CVPR2025满分论文，使扩散模型在逆逻辑和长依赖任务中更具实用价值。论文与代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-27T07:02:41Z
- **目录日期**: 2025-05-27
