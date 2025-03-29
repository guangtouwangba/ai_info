# Reimagining Memory Access for LLM Inference: Compression-Aware Memory Controller Design

**URL**: http://arxiv.org/abs/2503.18869v2

## 原始摘要

The efficiency of Large Language Model~(LLM) inference is often constrained
by substantial memory bandwidth and capacity demands. Existing techniques, such
as pruning, quantization, and mixture of experts/depth, reduce memory capacity
and/or bandwidth consumption at the cost of slight degradation in inference
quality. This paper introduces a design solution that further alleviates memory
bottlenecks by enhancing the on-chip memory controller in AI accelerators to
achieve two main objectives: (1) significantly reducing memory capacity and
bandwidth usage through lossless block compression~(e.g., LZ4 and ZSTD) of
model weights and key-value (KV) cache without compromising inference quality,
and (2) enabling memory bandwidth and energy consumption to scale
proportionally with context-dependent dynamic quantization. These goals are
accomplished by equipping the on-chip memory controller with mechanisms to
improve fine-grained bit-level accessibility and compressibility of weights and
KV cache through LLM-aware configuration of in-memory placement and
representation. Experimental results on publicly available LLMs demonstrate the
effectiveness of this approach, showing memory footprint reductions of 25.2\%
for model weights and 46.9\% for KV cache. In addition, our hardware prototype
at 4\,GHz and 32 lanes (7\,nm) achieves 8\,TB/s throughput with a modest area
overhead (under 3.8\,mm\(^2\)), which underscores the viability of LLM-aware
memory control as a key to efficient large-scale inference.


## AI 摘要

这篇论文提出了一种通过改进AI加速器的片上内存控制器来提升大语言模型(LLM)推理效率的方法。该方案采用无损块压缩(如LZ4和ZSTD)对模型权重和键值(KV)缓存进行压缩，在不影响推理质量的前提下，显著减少了内存占用和带宽消耗。实验显示，该方法可将模型权重内存占用减少25.2%，KV缓存减少46.9%。硬件原型在7nm工艺下实现了8TB/s的吞吐量，面积开销仅为3.8mm²，证明这种LLM感知的内存控制是实现高效大规模推理的关键。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-29T19:01:54Z
- **目录日期**: 2025-03-29
