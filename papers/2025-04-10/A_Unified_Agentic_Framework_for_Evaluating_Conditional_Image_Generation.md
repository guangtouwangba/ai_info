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

本文提出了CIGEval框架，利用大型多模态模型（LMMs）作为核心，结合多功能工具箱和细粒度评估体系，为条件图像生成任务提供统一、可靠的自动化评估方案。实验表明，GPT-4o版本的CIGEval与人工评估的相关系数达0.4625（接近人工评估者间0.47的相关性），而仅用2.3K训练轨迹微调的7B开源LMM版本甚至超越基于GPT-4o的现有最优方法。案例研究证实该框架能识别图像生成中主体一致性和控制引导遵循度等细微问题，展现出接近人类水平的自动化评估潜力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-10T21:01:33Z
- **目录日期**: 2025-04-10
