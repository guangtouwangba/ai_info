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

随着机器学习系统的快速发展，评估其环境影响（尤其是碳足迹）变得至关重要。现有研究多关注运行碳排放（训练和推理），而忽略了硬件制造和全生命周期的隐含碳排放。为此，研究者提出CATransformers框架，首次将运行和隐含碳排放共同纳入ML模型与硬件架构的协同优化。该框架应用于多模态CLIP模型，开发出CarbonCLIP系列，在保持准确性和延迟的同时，总碳排放比现有最优边缘CLIP模型降低17%。研究表明，仅优化延迟或能效的设计选择与碳优化方案存在显著差异，凸显了可持续AI系统设计需要整体优化方法。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-05T12:01:11Z
- **目录日期**: 2025-05-05
