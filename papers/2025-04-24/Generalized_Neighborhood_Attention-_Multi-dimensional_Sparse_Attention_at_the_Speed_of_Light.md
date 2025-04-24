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

本文提出了一种广义邻域注意力机制(GNA)，用于优化稀疏注意力计算性能。研究针对现有稀疏注意力方法(如邻域注意力)难以稳定超越标准自注意力计算速度的问题，通过分析局部性稀疏模式，建立了更精确的性能分析模型。GNA可统一描述滑动窗口、跨步窗口和分块注意力等多种稀疏模式，并开发了模拟器预测实际加速上限。研究在NVIDIA Blackwell架构上实现了基于CUTLASS的高效FMHA内核，在理想分块稀疏情况下达到1.3 petaFLOP/s的FP16计算利用率。实验表明，GNA在Cosmos-7B等模型上可实现28%-46%的端到端加速，相关工具将通过NATTEN项目开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-24T11:01:01Z
- **目录日期**: 2025-04-24
