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

该研究提出了一种新范式，通过动态空间上下文增强视觉语言模型（VLM）的3D场景生成与理解能力。该上下文包含三层结构：高层语义蓝图（场景画像）、带语义标签的点云（几何信息）和编码多元空间关系的场景超图。这种结构化记忆将VLM的多模态推理与3D空间理解相结合，支持迭代式的智能场景生成流程，包括几何修复的资产生成、自动验证的环境布置和基于超图的人体工学调整。实验表明该方法能处理多样化输入，并支持交互式场景编辑和路径规划等下游任务，为空间智能系统开辟了新可能。（99字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-27T21:01:26Z
- **目录日期**: 2025-05-27
