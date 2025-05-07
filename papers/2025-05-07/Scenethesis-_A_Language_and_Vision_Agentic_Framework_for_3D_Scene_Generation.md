# Scenethesis: A Language and Vision Agentic Framework for 3D Scene Generation

**URL**: http://arxiv.org/abs/2505.02836v1

## 原始摘要

Synthesizing interactive 3D scenes from text is essential for gaming, virtual
reality, and embodied AI. However, existing methods face several challenges.
Learning-based approaches depend on small-scale indoor datasets, limiting the
scene diversity and layout complexity. While large language models (LLMs) can
leverage diverse text-domain knowledge, they struggle with spatial realism,
often producing unnatural object placements that fail to respect common sense.
Our key insight is that vision perception can bridge this gap by providing
realistic spatial guidance that LLMs lack. To this end, we introduce
Scenethesis, a training-free agentic framework that integrates LLM-based scene
planning with vision-guided layout refinement. Given a text prompt, Scenethesis
first employs an LLM to draft a coarse layout. A vision module then refines it
by generating an image guidance and extracting scene structure to capture
inter-object relations. Next, an optimization module iteratively enforces
accurate pose alignment and physical plausibility, preventing artifacts like
object penetration and instability. Finally, a judge module verifies spatial
coherence. Comprehensive experiments show that Scenethesis generates diverse,
realistic, and physically plausible 3D interactive scenes, making it valuable
for virtual content creation, simulation environments, and embodied AI
research.


## AI 摘要

本文提出Scenethesis框架，通过结合大语言模型(LLM)和视觉感知技术，从文本生成交互式3D场景。传统方法受限于小规模数据集或缺乏空间合理性，而Scenethesis首先用LLM生成粗略布局，再通过视觉模块优化物体间关系，最后用优化模块确保姿态对齐和物理合理性，避免物体穿透等问题。验证模块会检查空间一致性。实验表明，该方法能生成多样、真实且物理合理的3D场景，适用于虚拟内容创作、仿真环境和具身AI研究。该框架无需训练，有效解决了现有方法在场景多样性和空间合理性方面的不足。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-07T01:28:34Z
- **目录日期**: 2025-05-07
