# CAD-Coder: An Open-Source Vision-Language Model for Computer-Aided Design Code Generation

**URL**: http://arxiv.org/abs/2505.14646v1

## 原始摘要

Efficient creation of accurate and editable 3D CAD models is critical in
engineering design, significantly impacting cost and time-to-market in product
innovation. Current manual workflows remain highly time-consuming and demand
extensive user expertise. While recent developments in AI-driven CAD generation
show promise, existing models are limited by incomplete representations of CAD
operations, inability to generalize to real-world images, and low output
accuracy. This paper introduces CAD-Coder, an open-source Vision-Language Model
(VLM) explicitly fine-tuned to generate editable CAD code (CadQuery Python)
directly from visual input. Leveraging a novel dataset that we
created--GenCAD-Code, consisting of over 163k CAD-model image and code
pairs--CAD-Coder outperforms state-of-the-art VLM baselines such as GPT-4.5 and
Qwen2.5-VL-72B, achieving a 100% valid syntax rate and the highest accuracy in
3D solid similarity. Notably, our VLM demonstrates some signs of
generalizability, successfully generating CAD code from real-world images and
executing CAD operations unseen during fine-tuning. The performance and
adaptability of CAD-Coder highlights the potential of VLMs fine-tuned on code
to streamline CAD workflows for engineers and designers. CAD-Coder is publicly
available at: https://github.com/anniedoris/CAD-Coder.


## AI 摘要

本文介绍了CAD-Coder，一种基于视觉语言模型（VLM）的开源工具，能够直接从视觉输入生成可编辑的CAD代码（CadQuery Python）。通过使用包含16.3万个CAD模型图像和代码对的新数据集GenCAD-Code，CAD-Coder在语法正确率和3D实体相似度上超越了GPT-4.5等先进模型，并展现出一定的泛化能力，能处理真实世界图像和未在训练中见过的CAD操作。该工具显著提升了工程设计中CAD模型创建的效率和准确性，有望简化工程师和设计师的工作流程。项目已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-21T18:02:28Z
- **目录日期**: 2025-05-21
