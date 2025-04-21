# #微软发布首个原生1bit模型##0.4G内存单CPU跑大模型#微软发布了首个开源“原生1bit”大模型：BitNet b1.58 2B4T。2B参数规模，用的是三进制权重（{-1, 0, 1}），...

**URL**: https://weibo.com/6105753431/PohITnfFD

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%BE%AE%E8%BD%AF%E5%8F%91%E5%B8%83%E9%A6%96%E4%B8%AA%E5%8E%9F%E7%94%9F1bit%E6%A8%A1%E5%9E%8B%23&amp;extparam=%23%E5%BE%AE%E8%BD%AF%E5%8F%91%E5%B8%83%E9%A6%96%E4%B8%AA%E5%8E%9F%E7%94%9F1bit%E6%A8%A1%E5%9E%8B%23" data-hide=""><span class="surl-text">#微软发布首个原生1bit模型#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%230.4G%E5%86%85%E5%AD%98%E5%8D%95CPU%E8%B7%91%E5%A4%A7%E6%A8%A1%E5%9E%8B%23&amp;extparam=%230.4G%E5%86%85%E5%AD%98%E5%8D%95CPU%E8%B7%91%E5%A4%A7%E6%A8%A1%E5%9E%8B%23" data-hide=""><span class="surl-text">#0.4G内存单CPU跑大模型#</span></a><br><br>微软发布了首个开源“原生1bit”大模型：BitNet b1.58 2B4T。<br><br>2B参数规模，用的是三进制权重（{-1, 0, 1}），平均每个权重只占1.58bit存储空间，对比传统16bit浮点模型，直接把内存需求干到只有0.4GB！而且还能单CPU运行，速度居然能跟人类读书差不多快。<br><br>你没听错，不是边缘设备“能跑”，而是“跑得很溜”：<br><br>- 普通CPU每秒生成5~7个token；<br><br>- 解码延迟低至29ms；<br><br>- 单次生成的能耗只有0.028焦耳，简直是省电界的“环境守护神”。<br><br>理论上，无论是办公笔记本、边缘设备，还是树莓派爱好者，都能轻松拿下。比如在Mac M2芯片上运行效果也超稳。<br><br>更猛的是，它性能竟然能和全精度模型打得有来有回！尤其在数学推理（GSM8K）和常识推理（WinoGrande）任务上，全面碾压了Llama和Qwen等同类模型。请看文章： <a href="https://weibo.com/ttarticle/p/show?id=2309405157880033837215" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">微软开源“原生1bit”三进制LLM：2B参数，0.4GB内存/单CPU就能跑，性</span></a><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i0oevnnxikj30rs0fmta9.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

微软发布全球首个开源原生1bit大模型BitNet b1.58 2B4T，采用三进制权重（-1,0,1），平均每个权重仅占1.58bit，内存需求仅0.4GB，可在单CPU上流畅运行（每秒5-7个token，延迟29ms）。其能效比惊人（0.028焦耳/生成），适用于笔记本、树莓派等设备，在Mac M2芯片表现稳定。性能上，该2B参数模型在数学推理（GSM8K）和常识推理（WinoGrande）任务超越Llama等全精度模型，实现高压缩与高性能的突破。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-21T07:02:37Z
- **目录日期**: 2025-04-21
