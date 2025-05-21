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

这篇论文介绍了CAD-Coder，一个开源的视觉语言模型（VLM），专门用于从视觉输入直接生成可编辑的CAD代码（CadQuery Python）。基于作者构建的163k+ CAD模型图像与代码对数据集GenCAD-Code，CAD-Coder在语法有效性（100%）和3D实体相似度准确性上超越了GPT-4.5等先进模型。该模型展现出一定的泛化能力，能处理真实世界图像并执行未在训练中见过的CAD操作。CAD-Coder的优异表现表明，基于代码微调的VLM有望简化工程师的CAD工作流程。项目已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-21T08:02:43Z
- **目录日期**: 2025-05-21
