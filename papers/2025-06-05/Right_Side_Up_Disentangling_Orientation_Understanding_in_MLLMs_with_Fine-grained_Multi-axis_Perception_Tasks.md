# Right Side Up? Disentangling Orientation Understanding in MLLMs with Fine-grained Multi-axis Perception Tasks

**URL**: http://arxiv.org/abs/2505.21649v4

## 原始摘要

Object orientation understanding represents a fundamental challenge in visual
perception critical for applications like robotic manipulation and augmented
reality. Current vision-language benchmarks fail to isolate this capability,
often conflating it with positional relationships and general scene
understanding. We introduce DORI (Discriminative Orientation Reasoning
Intelligence), a comprehensive benchmark establishing object orientation
perception as a primary evaluation target. DORI assesses four dimensions of
orientation comprehension: frontal alignment, rotational transformations,
relative directional relationships, and canonical orientation understanding.
Through carefully curated tasks from 11 datasets spanning 67 object categories
across synthetic and real-world scenarios, DORI provides insights on how
multi-modal systems understand object orientations. Our evaluation of 15
state-of-the-art vision-language models reveals critical limitations: even the
best models achieve only 54.2% accuracy on coarse tasks and 33.0% on granular
orientation judgments, with performance deteriorating for tasks requiring
reference frame shifts or compound rotations. These findings demonstrate the
need for dedicated orientation representation mechanisms, as models show
systematic inability to perform precise angular estimations, track orientation
changes across viewpoints, and understand compound rotations - suggesting
limitations in their internal 3D spatial representations. As the first
diagnostic framework specifically designed for orientation awareness in
multimodal systems, DORI offers implications for improving robotic control, 3D
scene reconstruction, and human-AI interaction in physical environments. DORI
data: https://huggingface.co/datasets/appledora/DORI-Benchmark


## AI 摘要

研究人员开发了DORI基准测试，专门评估视觉语言模型对物体朝向的理解能力。该测试包含四个维度：正面对齐、旋转变换、相对方向关系和标准朝向理解，涵盖11个数据集和67种物体。评估15个先进模型发现，即使在简单任务中最佳模型准确率仅54.2%，精细判断任务中降至33.0%，且模型在参考系转换和复合旋转任务中表现更差。结果表明现有模型缺乏精确角度估计、视角变化追踪和复合旋转理解能力，凸显其内部3D空间表征的局限性。DORI为改进机器人控制、3D场景重建和人机交互提供了诊断框架。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-05T11:02:22Z
- **目录日期**: 2025-06-05
