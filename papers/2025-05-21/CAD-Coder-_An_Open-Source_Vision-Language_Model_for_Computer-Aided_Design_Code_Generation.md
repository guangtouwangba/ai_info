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

这篇论文介绍了CAD-Coder，一个开源视觉语言模型（VLM），专门用于从视觉输入直接生成可编辑的CAD代码（CadQuery Python）。该模型基于新创建的GenCAD-Code数据集（包含超过16.3万组CAD模型图像和代码对），在生成有效语法和3D实体相似度方面表现优异，超越了GPT-4.5等先进模型。CAD-Coder还展现出一定的泛化能力，能够处理真实世界图像和未在训练中见过的CAD操作。这一成果展示了VLM在简化CAD工作流程中的潜力。项目已公开在GitHub上。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-21T23:02:12Z
- **目录日期**: 2025-05-21
