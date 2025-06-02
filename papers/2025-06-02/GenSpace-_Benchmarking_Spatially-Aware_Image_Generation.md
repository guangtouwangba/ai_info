# GenSpace: Benchmarking Spatially-Aware Image Generation

**URL**: http://arxiv.org/abs/2505.24870v1

## 原始摘要

Humans can intuitively compose and arrange scenes in the 3D space for
photography. However, can advanced AI image generators plan scenes with similar
3D spatial awareness when creating images from text or image prompts? We
present GenSpace, a novel benchmark and evaluation pipeline to comprehensively
assess the spatial awareness of current image generation models. Furthermore,
standard evaluations using general Vision-Language Models (VLMs) frequently
fail to capture the detailed spatial errors. To handle this challenge, we
propose a specialized evaluation pipeline and metric, which reconstructs 3D
scene geometry using multiple visual foundation models and provides a more
accurate and human-aligned metric of spatial faithfulness. Our findings show
that while AI models create visually appealing images and can follow general
instructions, they struggle with specific 3D details like object placement,
relationships, and measurements. We summarize three core limitations in the
spatial perception of current state-of-the-art image generation models: 1)
Object Perspective Understanding, 2) Egocentric-Allocentric Transformation and
3) Metric Measurement Adherence, highlighting possible directions for improving
spatial intelligence in image generation.


## AI 摘要

GenSpace是一个评估AI图像生成模型3D空间感知能力的新基准和评估流程。研究发现，尽管现有模型能生成视觉吸引人的图像并遵循一般指令，但在处理3D空间细节（如物体位置、关系和尺寸）时表现不佳。研究提出了专门评估流程和指标，通过多个视觉基础模型重建3D场景几何，提供更准确的空间保真度度量。研究总结了当前最先进图像生成模型的三大空间感知局限：1)物体视角理解，2)自我中心-他中心转换，3)度量标准遵循，为提升图像生成的空间智能指明了改进方向。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-02T22:01:17Z
- **目录日期**: 2025-06-02
