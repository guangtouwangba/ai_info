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

多模态大语言模型(MLLMs)在多视角理解方面仍面临重大挑战，特别是在几何一致性和跨视角对应关系上。研究者提出了All-Angles Bench基准测试，包含90个真实场景中的2,100个人工标注问答对，通过6项任务评估模型的几何对应能力和跨视角信息对齐能力。测试27个主流MLLMs(包括Gemini-2.0、Claude-3.7和GPT-4o)后发现，当前模型在部分遮挡视角的跨视角对应和粗略相机位姿估计方面表现不佳，与人类水平存在显著差距。该研究为提升MLLMs的多视角理解能力提供了重要参考。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-22T03:14:58Z
- **目录日期**: 2025-04-22
