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

本文提出广义邻域注意力机制(GNA)，通过分析局部稀疏注意力的性能改进模型，开发了更精确的模拟器预测速度上限。研究在NVIDIA Blackwell架构上实现GNA，基于CUTLASS优化的FMHA内核，在块稀疏情况下达到理论最大加速，FP16计算效率达1.3 petaFLOPs/秒。实验表明，GNA在Cosmos-7B等生成模型中无需微调即可带来28%-46%的端到端加速。该工作将为NATTEN项目开源模拟器和Blackwell内核实现，为突破注意力机制O(n²)复杂度提供可靠方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-24T09:00:57Z
- **目录日期**: 2025-04-24
