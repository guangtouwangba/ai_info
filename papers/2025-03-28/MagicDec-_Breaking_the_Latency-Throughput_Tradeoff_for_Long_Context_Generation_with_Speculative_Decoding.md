# MagicDec: Breaking the Latency-Throughput Tradeoff for Long Context Generation with Speculative Decoding

**URL**: http://arxiv.org/abs/2408.11049v4

## 原始摘要

Large Language Models (LLMs) have become more prevalent in long-context
applications such as interactive chatbots, document analysis, and agent
workflows, but it is challenging to serve long-context requests with low
latency and high throughput. Speculative decoding (SD) is a widely used
technique to reduce latency losslessly, but the conventional wisdom suggests
that its efficacy is limited to small batch sizes. In MagicDec, we show that
surprisingly SD can achieve speedup even for a high throughput inference regime
for moderate to long sequences. More interestingly, an intelligent drafting
strategy can achieve better speedup with increasing batch size based on our
rigorous analysis. MagicDec first identifies the bottleneck shifts with
increasing batch size and sequence length, and uses these insights to deploy SD
more effectively for high throughput inference. We leverage draft model with
sparse KV cache to address the KV bottleneck, which scales with both sequence
length and batch size. Additionally, we propose a theoretical model to select
the optimal drafting strategy for maximum speedup. Our work highlights the
broad applicability of speculative decoding in long-context serving, as it can
enhance throughput and reduce latency without compromising accuracy. For
moderate to long sequences, we demonstrate up to 2.51x speedup for Llama3.1-8B
when serving batch sizes ranging from 32 to 256 on various types of hardware
and tasks.


## AI 摘要

MagicDec研究表明，在长上下文LLM推理中，投机解码(SD)技术能显著提升吞吐量并降低延迟，即使在大批量场景下。通过分析瓶颈转移规律，该研究采用稀疏KV缓存的草稿模型解决计算瓶颈，并提出理论模型优化草稿策略。实验显示，在Llama3.1-8B模型上，针对中等至长序列（批量32-256）可实现最高2.51倍加速，且不损失准确性，突破了传统认为SD仅适用于小批量的认知。这项工作拓展了SD在长上下文服务中的适用性，为高效推理提供了新思路。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-28T00:02:35Z
- **目录日期**: 2025-03-28
