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

MagicDec研究发现，推测解码(SD)技术能显著提升大语言模型(LLMs)在长上下文场景下的推理效率。通过采用稀疏KV缓存和智能草拟策略，该系统成功克服了传统SD技术在大批量处理时的瓶颈问题。实验表明，对于中等至长序列处理，在32-256批量范围内，Llama3.1-8B模型最高可获得2.51倍的加速效果，且不损失准确性。该研究突破了"SD仅适用于小批量处理"的传统认知，为高吞吐量场景下的长文本服务提供了有效的低延迟解决方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-27T13:08:45Z
- **目录日期**: 2025-03-27
