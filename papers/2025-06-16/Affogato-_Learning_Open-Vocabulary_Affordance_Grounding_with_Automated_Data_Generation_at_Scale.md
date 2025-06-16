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

这篇论文介绍了Affogato，一个大规模基准数据集，包含15万个实例，标注了开放词汇文本描述和对应的3D功能热图，涵盖多样化的物体和交互。为解决功能定位任务中的细粒度局部定位、多有效交互区域模糊性及数据稀缺等挑战，研究团队开发了基于预训练视觉主干和文本条件热图解码器的视觉语言模型。该模型在现有2D和3D基准测试中表现优异，并展现出开放词汇跨领域泛化的有效性。Affogato数据集已公开共享，地址为：https://huggingface.co/datasets/project-affogato/affogato。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-16T03:25:42Z
- **目录日期**: 2025-06-16
