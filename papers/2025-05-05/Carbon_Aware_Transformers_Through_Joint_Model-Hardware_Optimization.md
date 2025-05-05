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

随着机器学习系统的快速发展，评估其环境影响（包括运行碳排放和硬件制造的全生命周期隐含碳排放）变得至关重要。针对缺乏量化与优化机器学习系统总碳足迹工具的问题，研究者提出了CATransformers框架，通过碳感知架构搜索实现可持续性驱动的模型与硬件协同优化。该框架将运行和隐含碳排放纳入硬件加速器设计空间探索，发现碳优化方案与单纯优化延迟或能效的选择不同。应用于多模态CLIP模型时，生成的CarbonCLIP系列在保持准确性和延迟的同时，总碳排放比现有边缘小型CLIP基准降低达17%。该研究强调需要整体优化方法来设计高性能且环保的AI系统。（100字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-05T17:01:08Z
- **目录日期**: 2025-05-05
