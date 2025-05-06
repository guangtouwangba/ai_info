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

Scenethesis是一个基于文本生成交互式3D场景的训练免费框架，结合了大型语言模型（LLM）的场景规划和视觉引导的布局优化。LLM首先生成粗略布局，视觉模块通过图像引导和场景结构提取优化对象关系。优化模块迭代调整姿态和物理合理性，避免物体穿透和不稳定问题，最后由验证模块确保空间一致性。实验表明，Scenethesis能生成多样、真实且物理合理的3D场景，适用于虚拟内容创作、模拟环境和具身AI研究，弥补了现有方法在空间真实性和布局复杂性上的不足。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-06T21:00:54Z
- **目录日期**: 2025-05-06
