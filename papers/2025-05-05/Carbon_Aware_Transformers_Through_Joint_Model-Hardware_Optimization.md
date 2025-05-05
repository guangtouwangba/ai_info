# Carbon Aware Transformers Through Joint Model-Hardware Optimization

**URL**: http://arxiv.org/abs/2505.01386v1

## 原始摘要

The rapid growth of machine learning (ML) systems necessitates a more
comprehensive evaluation of their environmental impact, particularly their
carbon footprint, which comprises operational carbon from training and
inference execution and embodied carbon from hardware manufacturing and its
entire life-cycle. Despite the increasing importance of embodied emissions,
there is a lack of tools and frameworks to holistically quantify and optimize
the total carbon footprint of ML systems. To address this, we propose
CATransformers, a carbon-aware architecture search framework that enables
sustainability-driven co-optimization of ML models and hardware architectures.
By incorporating both operational and embodied carbon metrics into early design
space exploration of domain-specific hardware accelerators, CATransformers
demonstrates that optimizing for carbon yields design choices distinct from
those optimized solely for latency or energy efficiency. We apply our framework
to multi-modal CLIP-based models, producing CarbonCLIP, a family of CLIP models
achieving up to 17% reduction in total carbon emissions while maintaining
accuracy and latency compared to state-of-the-art edge small CLIP baselines.
This work underscores the need for holistic optimization methods to design
high-performance, environmentally sustainable AI systems.


## AI 摘要

针对机器学习系统环境影响评估不足的问题，本研究提出了CATransformers框架，通过碳感知架构搜索同时优化模型和硬件架构。该框架首次将运行碳排放（训练/推理）和隐含碳排放（硬件全生命周期）纳入设计空间探索，发现碳优化方案与传统时延/能效优化存在显著差异。应用于CLIP模型时，所开发的CarbonCLIP系列在保持精度和时延的同时，较先进边缘小模型基准实现了17%的总碳减排。这项工作凸显了需要整体优化方法来构建高性能且环境可持续的AI系统。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-05T02:30:34Z
- **目录日期**: 2025-05-05
