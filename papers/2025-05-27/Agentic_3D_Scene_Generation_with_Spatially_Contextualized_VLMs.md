# Agentic 3D Scene Generation with Spatially Contextualized VLMs

**URL**: http://arxiv.org/abs/2505.20129v1

## 原始摘要

Despite recent advances in multimodal content generation enabled by
vision-language models (VLMs), their ability to reason about and generate
structured 3D scenes remains largely underexplored. This limitation constrains
their utility in spatially grounded tasks such as embodied AI, immersive
simulations, and interactive 3D applications. We introduce a new paradigm that
enables VLMs to generate, understand, and edit complex 3D environments by
injecting a continually evolving spatial context. Constructed from multimodal
input, this context consists of three components: a scene portrait that
provides a high-level semantic blueprint, a semantically labeled point cloud
capturing object-level geometry, and a scene hypergraph that encodes rich
spatial relationships, including unary, binary, and higher-order constraints.
Together, these components provide the VLM with a structured, geometry-aware
working memory that integrates its inherent multimodal reasoning capabilities
with structured 3D understanding for effective spatial reasoning. Building on
this foundation, we develop an agentic 3D scene generation pipeline in which
the VLM iteratively reads from and updates the spatial context. The pipeline
features high-quality asset generation with geometric restoration, environment
setup with automatic verification, and ergonomic adjustment guided by the scene
hypergraph. Experiments show that our framework can handle diverse and
challenging inputs, achieving a level of generalization not observed in prior
work. Further results demonstrate that injecting spatial context enables VLMs
to perform downstream tasks such as interactive scene editing and path
planning, suggesting strong potential for spatially intelligent systems in
computer graphics, 3D vision, and embodied applications.


## AI 摘要

本文提出了一种新方法，通过为视觉语言模型(VLM)注入动态演进的空间上下文，使其能够生成、理解和编辑复杂的3D场景。该上下文包含三个组件：高层语义蓝图(场景画像)、带语义标签的点云(几何信息)和编码丰富空间关系的场景超图(包括一元、二元和高阶约束)。基于此，研究者开发了一个自主3D场景生成流程，VLM可迭代更新空间上下文，实现高质量资产生成、环境设置和人体工学调整。实验表明，该方法能处理多样化输入，实现前所未有的泛化能力，并支持交互式场景编辑和路径规划等下游任务，在计算机图形学、3D视觉和具身应用领域具有重要潜力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-27T10:01:35Z
- **目录日期**: 2025-05-27
