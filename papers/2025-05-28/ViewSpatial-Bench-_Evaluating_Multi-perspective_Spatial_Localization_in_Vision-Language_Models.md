# ViewSpatial-Bench: Evaluating Multi-perspective Spatial Localization in Vision-Language Models

**URL**: http://arxiv.org/abs/2505.21500v1

## 原始摘要

Vision-language models (VLMs) have demonstrated remarkable capabilities in
understanding and reasoning about visual content, but significant challenges
persist in tasks requiring cross-viewpoint understanding and spatial reasoning.
We identify a critical limitation: current VLMs excel primarily at egocentric
spatial reasoning (from the camera's perspective) but fail to generalize to
allocentric viewpoints when required to adopt another entity's spatial frame of
reference. We introduce ViewSpatial-Bench, the first comprehensive benchmark
designed specifically for multi-viewpoint spatial localization recognition
evaluation across five distinct task types, supported by an automated 3D
annotation pipeline that generates precise directional labels. Comprehensive
evaluation of diverse VLMs on ViewSpatial-Bench reveals a significant
performance disparity: models demonstrate reasonable performance on
camera-perspective tasks but exhibit reduced accuracy when reasoning from a
human viewpoint. By fine-tuning VLMs on our multi-perspective spatial dataset,
we achieve an overall performance improvement of 46.24% across tasks,
highlighting the efficacy of our approach. Our work establishes a crucial
benchmark for spatial intelligence in embodied AI systems and provides
empirical evidence that modeling 3D spatial relationships enhances VLMs'
corresponding spatial comprehension capabilities.


## AI 摘要

当前视觉语言模型（VLM）在自我中心视角（相机视角）的空间推理表现良好，但在跨视角（如他人视角）的空间理解上存在显著不足。为此，研究者提出了首个多视角空间定位基准测试ViewSpatial-Bench，包含5类任务，并通过自动化3D标注生成精确方向标签。测试发现VLM在相机视角任务中表现尚可，但在人类视角推理时准确率明显下降。通过多视角空间数据微调后，模型整体性能提升46.24%，证明建模3D空间关系能有效增强VLM的空间理解能力。该研究为具身AI系统的空间智能评估提供了重要基准。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-28T04:02:50Z
- **目录日期**: 2025-05-28
