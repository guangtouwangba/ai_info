# 7B Fully Open Source Moxin-LLM/VLM -- From Pretraining to GRPO-based Reinforcement Learning Enhancement

**URL**: http://arxiv.org/abs/2412.06845v5

## 原始摘要

Recently, Large Language Models (LLMs) have undergone a significant
transformation, marked by a rapid rise in both their popularity and
capabilities. Leading this evolution are proprietary LLMs like GPT-4 and
GPT-o1, which have captured widespread attention in the AI community due to
their remarkable performance and versatility. Simultaneously, open-source LLMs,
such as LLaMA, have made great contributions to the ever-increasing popularity
of LLMs due to the ease to customize and deploy the models across diverse
applications. Although open-source LLMs present unprecedented opportunities for
innovation and research, the commercialization of LLMs has raised concerns
about transparency, reproducibility, and safety. Many open-source LLMs fail to
meet fundamental transparency requirements by withholding essential components
like training code and data, which may hinder further innovations on LLMs. To
mitigate this issue, we introduce Moxin 7B, a fully open-source LLM developed,
adhering to principles of open science, open source, open data, and open
access. We release the pre-training code and configurations, training and
fine-tuning datasets, and intermediate and final checkpoints, aiming to make
continuous commitments to fully open-source LLMs. After pre-training the base
model, we finetune the Moxin Base model with SOTA post-training framework and
instruction data to obtain Moxin Instruct model. To improve the reasoning
capability, we further finetune our Instruct model with chain-of-thought data
distilled from DeepSeek R1, and then use Group Relative Policy Optimization
(GRPO) following DeepSeek R1 to finetune our model, leading to the Moxin
Reasoning model. Moreover, we develop our vision language model based on our
Moxin model. Experiments show that our models achieve superior performance in
various evaluations such as zero-shot evaluation, few-shot evaluation, and CoT
evaluation.


## AI 摘要

近年来，大语言模型（LLMs）快速发展，以GPT-4为代表的专有模型和以LLaMA为代表的开源模型推动了这一趋势。然而，开源LLMs的商业化引发了透明度、可复现性和安全性的担忧。为此，研究团队推出了完全开源的Moxin 7B模型，公开了预训练代码、配置、数据集及检查点。该模型经过基础训练后，通过SOTA后训练框架和指令数据微调获得Moxin Instruct模型，并进一步使用思维链数据和GRPO方法优化推理能力，形成Moxin Reasoning模型。实验表明，该模型在零样本、少样本和思维链评估中表现优异。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-13T01:30:13Z
- **目录日期**: 2025-06-13
