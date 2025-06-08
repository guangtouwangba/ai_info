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

这篇论文提出了DirectLayout框架，利用大语言模型（LLMs）的空间推理能力直接从文本描述生成3D室内场景布局。该方法将生成过程分为三个阶段：鸟瞰图布局生成、3D空间提升和物体位置优化。通过基于3D-Front数据集的思维链（CoT）激活和设计的CoT-Grounded生成式布局奖励，增强了模型的泛化能力和空间规划。推理阶段采用上下文学习的迭代资产-布局对齐来解决资产与布局不匹配问题。实验表明，DirectLayout在语义一致性、泛化性和物理合理性方面表现优异，克服了现有方法在开放词汇和细粒度控制上的局限。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-08T00:01:11Z
- **目录日期**: 2025-06-08
