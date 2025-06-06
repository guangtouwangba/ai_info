# SuperWriter: Reflection-Driven Long-Form Generation with Large Language Models

**URL**: http://arxiv.org/abs/2506.04180v1

## 原始摘要

Long-form text generation remains a significant challenge for large language
models (LLMs), particularly in maintaining coherence, ensuring logical
consistency, and preserving text quality as sequence length increases. To
address these limitations, we propose SuperWriter-Agent, an agent-based
framework designed to enhance the quality and consistency of long-form text
generation. SuperWriter-Agent introduces explicit structured thinking-through
planning and refinement stages into the generation pipeline, guiding the model
to follow a more deliberate and cognitively grounded process akin to that of a
professional writer. Based on this framework, we construct a supervised
fine-tuning dataset to train a 7B SuperWriter-LM. We further develop a
hierarchical Direct Preference Optimization (DPO) procedure that uses Monte
Carlo Tree Search (MCTS) to propagate final quality assessments and optimize
each generation step accordingly. Empirical results across diverse benchmarks
demonstrate that SuperWriter-LM achieves state-of-the-art performance,
surpassing even larger-scale baseline models in both automatic evaluation and
human evaluation. Furthermore, comprehensive ablation studies demonstrate the
effectiveness of hierarchical DPO and underscore the value of incorporating
structured thinking steps to improve the quality of long-form text generation.


## AI 摘要

本文提出了SuperWriter-Agent框架，通过结构化思维（规划与优化阶段）提升大语言模型的长文本生成质量。基于该框架，作者构建了监督微调数据集训练7B参数的SuperWriter-LM模型，并开发了分层直接偏好优化（DPO）方法，利用蒙特卡洛树搜索（MCTS）传播质量评估以优化生成过程。实验表明，SuperWriter-LM在多项基准测试中表现优异，自动和人工评估均超越更大规模的基线模型。消融研究验证了分层DPO的有效性，证实结构化思维步骤能显著提升长文本生成质量。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-06T02:31:13Z
- **目录日期**: 2025-06-06
