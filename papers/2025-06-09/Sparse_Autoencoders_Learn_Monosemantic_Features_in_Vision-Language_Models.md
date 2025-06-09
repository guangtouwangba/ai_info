# Sparse Autoencoders Learn Monosemantic Features in Vision-Language Models

**URL**: http://arxiv.org/abs/2504.02821v2

## 原始摘要

Given that interpretability and steerability are crucial to AI safety, Sparse
Autoencoders (SAEs) have emerged as a tool to enhance them in Large Language
Models (LLMs). In this work, we extend the application of SAEs to
Vision-Language Models (VLMs), such as CLIP, and introduce a comprehensive
framework for evaluating monosemanticity at the neuron-level in vision
representations. To ensure that our evaluation aligns with human perception, we
propose a benchmark derived from a large-scale user study. Our experimental
results reveal that SAEs trained on VLMs significantly enhance the
monosemanticity of individual neurons, with sparsity and wide latents being the
most influential factors. Notably, we demonstrate that applying SAE
interventions on CLIP's vision encoder directly steers multimodal LLM outputs
(e.g., LLaVA), without any modifications to the underlying model. These
findings emphasize the practicality and efficacy of SAEs as an unsupervised
tool for enhancing both interpretability and control of VLMs. Code is available
at https://github.com/ExplainableML/sae-for-vlm.


## AI 摘要

稀疏自编码器（SAEs）被用于提升视觉语言模型（VLMs，如CLIP）的可解释性和可控性。本研究提出了一种评估视觉表征中神经元单义性的框架，并通过大规模用户研究验证其与人类感知的一致性。实验表明，SAEs显著提高了神经元的单义性，其中稀疏性和宽潜在层是关键因素。此外，直接在CLIP视觉编码器上应用SAEs干预可引导多模态大模型（如LLaVA）的输出，无需修改原模型。这证实了SAEs作为无监督工具在增强VLM可解释性和控制方面的实用性与有效性。代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-09T22:02:15Z
- **目录日期**: 2025-06-09
