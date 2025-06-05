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

DORI是一个专注于评估视觉语言模型物体方向感知能力的基准测试，包含四个维度：正面对齐、旋转变换、相对方向关系和标准方向理解。通过11个数据集的67个物体类别测试，发现当前最先进的15个模型表现有限，在粗粒度任务中最高准确率仅54.2%，细粒度方向判断降至33.0%，且在参考系转换和复合旋转任务中表现更差。研究表明现有模型缺乏精确角度估计、视角变化跟踪和复合旋转理解能力，揭示了其内部3D空间表征的局限性。DORI为改进机器人控制、3D场景重建和人机交互提供了重要参考。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-05T15:02:32Z
- **目录日期**: 2025-06-05
