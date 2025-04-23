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

多视角理解能力（即在不同视角下协调视觉信息以实现导航、操作和3D场景理解）是多模态大语言模型（MLLMs）作为具身智能体的关键挑战。尽管MLLMs在高层次推理和规划方面取得进展，但在多视角几何一致性和跨视角对应关系上仍表现不足。为此，研究者提出了All-Angles Bench基准测试，包含90个真实场景中2,100个人工标注的多视角问答对，涵盖6项任务（计数、属性识别、相对距离/方向、物体操作和相机姿态估计）。测试27个主流MLLMs（如GPT-4o、Gemini-2.0-Flash）后发现，模型在遮挡视角对应和粗略相机姿态估计方面显著落后于人类水平，凸显了增强多视角感知模块的必要性。该基准已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-23T00:01:17Z
- **目录日期**: 2025-04-23
