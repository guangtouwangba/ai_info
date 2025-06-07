# Direct Numerical Layout Generation for 3D Indoor Scene Synthesis via Spatial Reasoning

**URL**: http://arxiv.org/abs/2506.05341v1

## 原始摘要

Realistic 3D indoor scene synthesis is vital for embodied AI and digital
content creation. It can be naturally divided into two subtasks: object
generation and layout generation. While recent generative models have
significantly advanced object-level quality and controllability, layout
generation remains challenging due to limited datasets. Existing methods either
overfit to these datasets or rely on predefined constraints to optimize
numerical layout that sacrifice flexibility. As a result, they fail to generate
scenes that are both open-vocabulary and aligned with fine-grained user
instructions. We introduce DirectLayout, a framework that directly generates
numerical 3D layouts from text descriptions using generalizable spatial
reasoning of large language models (LLMs). DirectLayout decomposes the
generation into three stages: producing a Bird's-Eye View (BEV) layout, lifting
it into 3D space, and refining object placements. To enable explicit spatial
reasoning and help the model grasp basic principles of object placement, we
employ Chain-of-Thought (CoT) Activation based on the 3D-Front dataset.
Additionally, we design CoT-Grounded Generative Layout Reward to enhance
generalization and spatial planning. During inference, DirectLayout addresses
asset-layout mismatches via Iterative Asset-Layout Alignment through in-context
learning. Extensive experiments demonstrate that DirectLayout achieves
impressive semantic consistency, generalization and physical plausibility.


## AI 摘要

这篇论文提出了DirectLayout框架，利用大语言模型的空间推理能力直接从文本生成3D室内场景布局。该方法分三阶段：首先生成鸟瞰图布局，然后提升到3D空间，最后优化物体摆放。通过3D-Front数据集激活思维链机制，帮助模型掌握物体摆放原则，并设计了基于思维链的生成布局奖励机制增强泛化能力。推理时采用上下文学习进行迭代式资产-布局对齐。实验表明，DirectLayout在语义一致性、泛化能力和物理合理性方面表现优异，解决了现有方法在开放词汇和细粒度控制方面的不足。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-07T19:00:50Z
- **目录日期**: 2025-06-07
