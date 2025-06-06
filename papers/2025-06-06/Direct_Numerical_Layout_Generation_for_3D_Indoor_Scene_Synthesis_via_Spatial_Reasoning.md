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

DirectLayout是一个基于大语言模型（LLM）的3D室内场景生成框架，通过文本描述直接生成数值化3D布局。它将生成过程分为三个阶段：鸟瞰图（BEV）布局生成、3D空间提升和物体位置优化。该框架利用3D-Front数据集进行显式空间推理，并通过思维链（CoT）激活帮助模型掌握物体摆放原则，同时设计了CoT-Grounded生成式布局奖励以增强泛化能力。推理时通过上下文学习进行迭代式资产-布局对齐，解决了资产与布局不匹配的问题。实验表明，DirectLayout在语义一致性、泛化能力和物理合理性方面表现优异。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-06T07:01:09Z
- **目录日期**: 2025-06-06
