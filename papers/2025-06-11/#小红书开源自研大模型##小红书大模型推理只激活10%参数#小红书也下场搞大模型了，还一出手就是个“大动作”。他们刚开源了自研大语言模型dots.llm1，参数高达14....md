# #小红书开源自研大模型##小红书大模型推理只激活10%参数#小红书也下场搞大模型了，还一出手就是个“大动作”。他们刚开源了自研大语言模型dots.llm1，参数高达14...

**URL**: https://weibo.com/6105753431/Pw0Chev89

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%B0%8F%E7%BA%A2%E4%B9%A6%E5%BC%80%E6%BA%90%E8%87%AA%E7%A0%94%E5%A4%A7%E6%A8%A1%E5%9E%8B%23&amp;extparam=%23%E5%B0%8F%E7%BA%A2%E4%B9%A6%E5%BC%80%E6%BA%90%E8%87%AA%E7%A0%94%E5%A4%A7%E6%A8%A1%E5%9E%8B%23" data-hide=""><span class="surl-text">#小红书开源自研大模型#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%B0%8F%E7%BA%A2%E4%B9%A6%E5%A4%A7%E6%A8%A1%E5%9E%8B%E6%8E%A8%E7%90%86%E5%8F%AA%E6%BF%80%E6%B4%BB10%25%E5%8F%82%E6%95%B0%23&amp;extparam=%23%E5%B0%8F%E7%BA%A2%E4%B9%A6%E5%A4%A7%E6%A8%A1%E5%9E%8B%E6%8E%A8%E7%90%86%E5%8F%AA%E6%BF%80%E6%B4%BB10%25%E5%8F%82%E6%95%B0%23" data-hide=""><span class="surl-text">#小红书大模型推理只激活10%参数#</span></a><br><br>小红书也下场搞大模型了，还一出手就是个“大动作”。<br><br>他们刚开源了自研大语言模型dots.llm1，参数高达1420亿，但推理时只激活10%——约140亿参数，做到了“省钱不降质”。在中文任务上，甚至干过了阿里家的Qwen2.5，还顺手比DeepSeek系列新模型分数更高。<br><br>以下是dots.llm1的重点信息，看看这个“小红书大模型”有多能打：<br><br>- 类型：MoE（专家混合模型），结构上是decoder-only Transformer；<br>    <br>- 参数量：总共1420亿，但每次推理只用到140亿，大大降低算力成本；<br>    <br>- 架构来源：在DeepSeekMoE的基础上改进。<br>    <br>在中文、数学和代码任务上，dots.llm1都小胜Qwen2.5-72B：<br><br>- 中文任务：91.3分（领先约1分）<br>- 数学任务：78.3分 vs 77.3<br>- 代码任务：59.6分 vs 59.0<br>- 英文任务稍弱：75.7 vs 76.3（但成本更低）<br>  <br>Github：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fgithub.com%2Frednote-hilab%2Fdots.llm1" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>论文：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fwww.arxiv.org%2Fabs%2F2506.05767" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i2b5nn3mjij313c0swn8l.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i2b5nnxv8oj30x60j2wi8.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

小红书开源了自研大模型dots.llm1，采用MoE架构，总参数量1420亿但推理时仅激活140亿（10%），显著降低算力成本。该模型基于DeepSeekMoE改进，在中文任务（91.3分）、数学（78.3分）和代码（59.6分）上均小幅领先阿里的Qwen2.5-72B，英文任务稍弱（75.7分）但成本更低。其高效设计实现了"低成本高性能"，相关代码和论文已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-11T02:33:33Z
- **目录日期**: 2025-06-11
