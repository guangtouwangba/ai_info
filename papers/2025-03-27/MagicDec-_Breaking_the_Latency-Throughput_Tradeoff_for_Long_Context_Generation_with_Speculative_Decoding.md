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

MagicDec研究发现，推测解码(SD)技术可显著提升大语言模型(LLMs)在长上下文场景下的推理效率。通过分析不同批量和序列长度下的性能瓶颈，该方法采用稀疏KV缓存来缓解KV瓶颈，并提出理论模型选择最优推测策略。实验显示，在中等至长序列场景下，Llama3.1-8B模型在32-256批量范围内最高可获得2.51倍加速，且不影响准确性。这项研究突破了传统认为SD仅适用于小批量的认知，证明其在高吞吐量推理中同样有效，为聊天机器人、文档分析等长上下文应用提供了高效解决方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-27T20:02:06Z
- **目录日期**: 2025-03-27
