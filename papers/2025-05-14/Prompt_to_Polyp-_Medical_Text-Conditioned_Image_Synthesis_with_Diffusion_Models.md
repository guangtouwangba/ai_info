# Prompt to Polyp: Medical Text-Conditioned Image Synthesis with Diffusion Models

**URL**: http://arxiv.org/abs/2505.05573v2

## 原始摘要

The generation of realistic medical images from text descriptions has
significant potential to address data scarcity challenges in healthcare AI
while preserving patient privacy. This paper presents a comprehensive study of
text-to-image synthesis in the medical domain, comparing two distinct
approaches: (1) fine-tuning large pre-trained latent diffusion models and (2)
training small, domain-specific models. We introduce a novel model named MSDM,
an optimized architecture based on Stable Diffusion that integrates a clinical
text encoder, variational autoencoder, and cross-attention mechanisms to better
align medical text prompts with generated images. Our study compares two
approaches: fine-tuning large pre-trained models (FLUX, Kandinsky) versus
training compact domain-specific models (MSDM). Evaluation across colonoscopy
(MedVQA-GI) and radiology (ROCOv2) datasets reveals that while large models
achieve higher fidelity, our optimized MSDM delivers comparable quality with
lower computational costs. Quantitative metrics and qualitative evaluations by
medical experts reveal strengths and limitations of each approach.


## AI 摘要

本文研究了医疗领域的文本到图像生成技术，比较了两种方法：(1)微调大型预训练潜在扩散模型和(2)训练小型领域专用模型。研究者提出了一种名为MSDM的新型优化架构，基于Stable Diffusion，整合了临床文本编码器、变分自编码器和交叉注意力机制，以更好地对齐医学文本提示与生成图像。评估显示，大型模型(如FLUX、Kandinsky)能获得更高保真度，而优化的MSDM在降低计算成本的同时提供了可比的质量。研究在结肠镜(MedVQA-GI)和放射学(ROCOv2)数据集上进行了定量指标和专家定性评估，揭示了各方法的优缺点。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-14T02:31:16Z
- **目录日期**: 2025-05-14
