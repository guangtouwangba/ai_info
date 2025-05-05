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

本文提出CATransformers框架，通过碳感知架构搜索技术，首次实现机器学习系统全生命周期碳排放（运行碳排放+硬件隐含碳排放）的协同优化。该框架在专用硬件加速器设计阶段即引入碳指标，证明碳优化方案与单纯追求延迟/能效的设计存在显著差异。应用于多模态CLIP模型时，生成的CarbonCLIP系列在保持准确性和延迟的同时，较现有边缘小型CLIP基线降低总碳排放达17%。研究强调需要开发整体优化方法，以构建高性能且环境可持续的AI系统。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-05T22:01:11Z
- **目录日期**: 2025-05-05
