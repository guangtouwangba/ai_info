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

DORI（Discriminative Orientation Reasoning Intelligence）是一个专注于评估多模态系统物体方向感知能力的基准测试，填补了现有视觉语言基准在隔离方向理解能力上的空白。该测试从四个维度（正面对齐、旋转变换、相对方向关系和标准方向理解）评估模型性能，涵盖11个数据集的67个物体类别。对15个先进模型的测试显示：最佳模型在粗粒度任务中仅达54.2%准确率，细粒度方向判断降至33.0%，且在参考系转换和复合旋转任务中表现更差。结果表明现有模型在内部3D空间表征和精确角度估计方面存在系统性缺陷。DORI为改进机器人控制、3D场景重建等应用提供了诊断框架。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-05T03:23:18Z
- **目录日期**: 2025-06-05
