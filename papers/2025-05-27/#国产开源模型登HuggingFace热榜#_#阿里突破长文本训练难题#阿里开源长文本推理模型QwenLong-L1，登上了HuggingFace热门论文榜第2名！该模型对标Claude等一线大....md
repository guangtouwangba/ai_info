# #国产开源模型登HuggingFace热榜# #阿里突破长文本训练难题#阿里开源长文本推理模型QwenLong-L1，登上了HuggingFace热门论文榜第2名！该模型对标Claude等一线大...

**URL**: https://weibo.com/6105753431/PtLTVEX0i

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%9B%BD%E4%BA%A7%E5%BC%80%E6%BA%90%E6%A8%A1%E5%9E%8B%E7%99%BBHuggingFace%E7%83%AD%E6%A6%9C%23&amp;extparam=%23%E5%9B%BD%E4%BA%A7%E5%BC%80%E6%BA%90%E6%A8%A1%E5%9E%8B%E7%99%BBHuggingFace%E7%83%AD%E6%A6%9C%23" data-hide=""><span class="surl-text">#国产开源模型登HuggingFace热榜#</span></a> <a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E9%98%BF%E9%87%8C%E7%AA%81%E7%A0%B4%E9%95%BF%E6%96%87%E6%9C%AC%E8%AE%AD%E7%BB%83%E9%9A%BE%E9%A2%98%23&amp;extparam=%23%E9%98%BF%E9%87%8C%E7%AA%81%E7%A0%B4%E9%95%BF%E6%96%87%E6%9C%AC%E8%AE%AD%E7%BB%83%E9%9A%BE%E9%A2%98%23" data-hide=""><span class="surl-text">#阿里突破长文本训练难题#</span></a><br><br>阿里开源长文本推理模型QwenLong-L1，登上了HuggingFace热门论文榜第2名！<br><br>该模型对标Claude等一线大模型，主打看得懂、想得深，还能自我反思。<br><br>长文本对于LLM来说，内容太多、干扰太多，模型容易跑偏。比如在金融文档推理中，传统模型会被无关信息绕晕，答题不靠谱。<br><br>而QwenLong-L1可以像人一样“回头看看”，通过验证机制筛掉噪音，抓住核心要点，给出正确答案。<br><br>来看它怎么做到的：<br><br>- 两阶段训练流程：先用高质量数据“热身”，再上强化学习（RL）精修。<br><br>- 分阶段强化学习：不是一口气喂128K token，而是先练短文，再逐步扩展上下文长度，最大支持到128K。<br><br>- 混合奖励机制：既检查答案是否正确（规则），又请另一个模型判断答得是否合理（语义），精准又灵活。<br><br>- 难题反复练：每阶段都会把“最难的题”留下来继续训练，不让模型忘了怎么应对复杂推理。<br><br>性能方面，QwenLong-L1的32B版本跑到了70.7分，超越了OpenAI-o3-mini和Qwen3-235B-A22B，跟Claude-3.7-Sonnet-Thinking打成平手。而它的14B 版本也表现不俗，比基础模型平均提升4.1分。<br><br>关键是，团队发现光靠监督学习，只学到表面套路，推理能力不稳定。只有用强化学习，才能让LLM学会像人一样“设子目标”、“回头查证”“排除干扰”。 <a href="https://weibo.com/ttarticle/p/show?id=2309405170956481462286" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">阿里开源QwenLong-L1，突破长文本训练难题，登HuggingFace热榜！</span></a><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i1u58a7thej30m90cjgnx.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

阿里开源的长文本推理模型QwenLong-L1登上HuggingFace热榜第二，突破长文本训练难题。该模型通过两阶段训练（预训练+强化学习）和分阶段扩展上下文（最大128K token），结合混合奖励机制和难题强化训练，显著提升长文本理解能力。其32B版本以70.7分超越部分主流模型，14B版本性能提升4.1分。研究显示强化学习能有效培养模型"反思验证"等类人推理能力，解决传统模型在金融文档等长文本中易受干扰的问题。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-27T19:04:01Z
- **目录日期**: 2025-05-27
