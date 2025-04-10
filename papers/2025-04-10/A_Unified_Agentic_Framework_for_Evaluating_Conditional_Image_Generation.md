# A Unified Agentic Framework for Evaluating Conditional Image Generation

**URL**: http://arxiv.org/abs/2504.07046v1

## 原始摘要

Conditional image generation has gained significant attention for its ability
to personalize content. However, the field faces challenges in developing
task-agnostic, reliable, and explainable evaluation metrics. This paper
introduces CIGEval, a unified agentic framework for comprehensive evaluation of
conditional image generation tasks. CIGEval utilizes large multimodal models
(LMMs) as its core, integrating a multi-functional toolbox and establishing a
fine-grained evaluation framework. Additionally, we synthesize evaluation
trajectories for fine-tuning, empowering smaller LMMs to autonomously select
appropriate tools and conduct nuanced analyses based on tool outputs.
Experiments across seven prominent conditional image generation tasks
demonstrate that CIGEval (GPT-4o version) achieves a high correlation of 0.4625
with human assessments, closely matching the inter-annotator correlation of
0.47. Moreover, when implemented with 7B open-source LMMs using only 2.3K
training trajectories, CIGEval surpasses the previous GPT-4o-based
state-of-the-art method. Case studies on GPT-4o image generation highlight
CIGEval's capability in identifying subtle issues related to subject
consistency and adherence to control guidance, indicating its great potential
for automating evaluation of image generation tasks with human-level
reliability.


## AI 摘要

本文提出CIGEval框架，利用多模态大模型（LMMs）评估条件图像生成任务。该框架整合多功能工具包，建立细粒度评估体系，并通过微调使小型LMMs能自主选择工具进行精准分析。实验表明，GPT-4o版本的CIGEval与人工评估相关性达0.4625（接近人工评估者间0.47的相关性）。仅用2.3K训练轨迹的7B开源LMMs实现时，其性能超越基于GPT-4o的现有最佳方法。案例研究证实CIGEval能识别主题一致性和控制引导等细微问题，展现自动化评估图像生成任务的潜力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-10T18:01:34Z
- **目录日期**: 2025-04-10
