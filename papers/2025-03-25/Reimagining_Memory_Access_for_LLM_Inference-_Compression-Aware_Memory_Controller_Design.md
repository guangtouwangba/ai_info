# Reimagining Memory Access for LLM Inference: Compression-Aware Memory Controller Design

**URL**: http://arxiv.org/abs/2503.18869v1

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

该研究提出了一种通过改进AI加速器的片上内存控制器来提升大语言模型(LLM)推理效率的方法。主要创新点包括：(1)采用无损块压缩技术(LZ4/ZSTD)压缩模型权重和KV缓存，在不影响推理质量的前提下，将内存占用分别降低25.2%和46.9%；(2)通过动态量化使内存带宽和能耗随上下文自适应调整。硬件原型(7nm工艺，4GHz，32通道)实现了8TB/s吞吐量，面积开销仅3.8mm²。实验证明这种LLM感知的内存控制方案能有效缓解内存瓶颈，为大规模推理提供高效解决方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-25T23:01:38Z
- **目录日期**: 2025-03-25
