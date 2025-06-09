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

稀疏自编码器(SAEs)被应用于视觉语言模型(VLMs)如CLIP，以提高其可解释性和可控性。研究提出了评估视觉表示中神经元单义性的框架，并通过大规模用户研究建立基准。实验表明，SAEs显著提升了VLMs神经元的单义性，其中稀疏性和宽潜在层是关键因素。特别地，在CLIP视觉编码器上应用SAE干预可直接引导多模态LLM(如LLaVA)的输出，而无需修改底层模型。这证实了SAEs作为无监督工具在增强VLMs可解释性和控制方面的实用性和有效性。代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-09T21:02:17Z
- **目录日期**: 2025-06-09
