# Affogato: Learning Open-Vocabulary Affordance Grounding with Automated Data Generation at Scale

**URL**: http://arxiv.org/abs/2506.12009v1

## 原始摘要

Affordance grounding-localizing object regions based on natural language
descriptions of interactions-is a critical challenge for enabling intelligent
agents to understand and interact with their environments. However, this task
remains challenging due to the need for fine-grained part-level localization,
the ambiguity arising from multiple valid interaction regions, and the scarcity
of large-scale datasets. In this work, we introduce Affogato, a large-scale
benchmark comprising 150K instances, annotated with open-vocabulary text
descriptions and corresponding 3D affordance heatmaps across a diverse set of
objects and interactions. Building on this benchmark, we develop simple yet
effective vision-language models that leverage pretrained part-aware vision
backbones and a text-conditional heatmap decoder. Our models trained with the
Affogato dataset achieve promising performance on the existing 2D and 3D
benchmarks, and notably, exhibit effectiveness in open-vocabulary cross-domain
generalization. The Affogato dataset is shared in public:
https://huggingface.co/datasets/project-affogato/affogato


## AI 摘要

Affordance grounding（功能定位）旨在通过自然语言描述定位物体交互区域，是智能体理解环境的关键挑战，但面临细粒度定位、多解歧义和数据稀缺等问题。研究团队推出Affogato大规模基准数据集，包含15万实例，涵盖多样物体和交互的开放词汇文本描述及3D功能热图。基于此，他们开发了高效的视觉-语言模型，结合预训练部件感知视觉主干和文本条件热图解码器。该模型在现有2D/3D基准测试中表现优异，并展现出优秀的开放词汇跨域泛化能力。数据集已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-16T02:33:34Z
- **目录日期**: 2025-06-16
