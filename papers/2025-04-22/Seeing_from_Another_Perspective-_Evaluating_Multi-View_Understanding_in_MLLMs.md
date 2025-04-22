# Seeing from Another Perspective: Evaluating Multi-View Understanding in MLLMs

**URL**: http://arxiv.org/abs/2504.15280v1

## 原始摘要

Multi-view understanding, the ability to reconcile visual information across
diverse viewpoints for effective navigation, manipulation, and 3D scene
comprehension, is a fundamental challenge in Multi-Modal Large Language Models
(MLLMs) to be used as embodied agents. While recent MLLMs have shown impressive
advances in high-level reasoning and planning, they frequently fall short when
confronted with multi-view geometric consistency and cross-view correspondence.
To comprehensively evaluate the challenges of MLLMs in multi-view scene
reasoning, we propose All-Angles Bench, a benchmark of over 2,100 human
carefully annotated multi-view question-answer pairs across 90 diverse
real-world scenes. Our six tasks (counting, attribute identification, relative
distance, relative direction, object manipulation, and camera pose estimation)
specifically test model's geometric correspondence and the capacity to align
information consistently across views. Our extensive experiments, benchmark on
27 representative MLLMs including Gemini-2.0-Flash, Claude-3.7-Sonnet, and
GPT-4o against human evaluators reveals a substantial performance gap,
indicating that current MLLMs remain far from human-level proficiency. Through
in-depth analysis, we show that MLLMs are particularly underperforming under
two aspects: (1) cross-view correspondence for partially occluded views and (2)
establishing the coarse camera poses. These findings highlight the necessity of
domain-specific refinements or modules that embed stronger multi-view
awareness. We believe that our All-Angles Bench offers valuable insights and
contribute to bridging the gap between MLLMs and human-level multi-view
understanding. The project and benchmark are publicly available at
https://danielchyeh.github.io/All-Angles-Bench/.


## AI 摘要

多模态大语言模型(MLLMs)在跨视角理解能力上仍存在显著不足。研究者提出All-Angles Bench基准测试，包含90个真实场景的2,100多组人工标注多视角问答对，通过6项任务(计数、属性识别、相对距离/方向、物体操控、相机姿态估计)系统评估模型的多视角推理能力。测试27个主流MLLMs(包括Gemini-2.0、Claude-3.7和GPT-4o)发现，模型在部分遮挡视角的对应关系和粗略相机姿态估计方面表现尤其不佳，与人类水平存在明显差距。该研究为提升MLLMs的多视角理解能力提供了重要参考。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-22T04:01:00Z
- **目录日期**: 2025-04-22
