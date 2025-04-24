# Generalized Neighborhood Attention: Multi-dimensional Sparse Attention at the Speed of Light

**URL**: http://arxiv.org/abs/2504.16922v1

## 原始摘要

Many sparse attention mechanisms such as Neighborhood Attention have
typically failed to consistently deliver speedup over the self attention
baseline. This is largely due to the level of complexity in attention
infrastructure, and the rapid evolution of AI hardware architecture. At the
same time, many state-of-the-art foundational models, particularly in computer
vision, are heavily bound by attention, and need reliable sparsity to escape
the O(n^2) complexity. In this paper, we study a class of promising sparse
attention mechanisms that focus on locality, and aim to develop a better
analytical model of their performance improvements. We first introduce
Generalized Neighborhood Attention (GNA), which can describe sliding window,
strided sliding window, and blocked attention. We then consider possible design
choices in implementing these approaches, and create a simulator that can
provide much more realistic speedup upper bounds for any given setting.
Finally, we implement GNA on top of a state-of-the-art fused multi-headed
attention (FMHA) kernel designed for the NVIDIA Blackwell architecture in
CUTLASS. Our implementation can fully realize the maximum speedup theoretically
possible in many perfectly block-sparse cases, and achieves an effective
utilization of 1.3 petaFLOPs/second in FP16. In addition, we plug various GNA
configurations into off-the-shelf generative models, such as Cosmos-7B,
HunyuanVideo, and FLUX, and show that it can deliver 28% to 46% end-to-end
speedup on B200 without any fine-tuning. We will open source our simulator and
Blackwell kernels directly through the NATTEN project.


## AI 摘要

本文提出了一种名为广义邻域注意力(GNA)的新型稀疏注意力机制，通过研究基于局部性的稀疏注意力，改进了性能分析模型。GNA能统一描述滑动窗口、跨步滑动窗口和分块注意力等多种形式。研究者开发了性能模拟器来预测实际加速上限，并在NVIDIA Blackwell架构上基于CUTLASS实现了优化的FMHA内核，在理想分块稀疏情况下达到1.3 petaFLOP/s的FP16计算效率。实验表明，GNA在Cosmos-7B等生成模型中可实现28%-46%的端到端加速，且无需微调。相关模拟器和Blackwell内核将通过NATTEN项目开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-24T15:01:07Z
- **目录日期**: 2025-04-24
