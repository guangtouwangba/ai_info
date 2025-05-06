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

Scenethesis是一个无需训练的智能框架，用于从文本生成交互式3D场景。它结合了基于大语言模型(LLM)的场景规划和视觉引导的布局优化。首先，LLM根据文本提示生成粗略布局；然后视觉模块通过图像引导提取场景结构，优化物体间关系；接着优化模块迭代调整物体姿态，确保物理合理性，避免穿透或不稳定；最后验证模块检查空间一致性。实验表明，Scenethesis能生成多样、真实且物理合理的3D场景，适用于虚拟内容创作、模拟环境和具身AI研究，解决了现有方法在多样性和空间合理性上的不足。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-06T12:01:00Z
- **目录日期**: 2025-05-06
