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

为解决大模型生成长文本时存在的连贯性、逻辑一致性和质量下降问题，研究者提出了SuperWriter-Agent框架。该框架通过规划与优化阶段引入结构化思维，模拟专业作者的写作流程。基于此开发了7B参数的SuperWriter-LM模型，并创新性地采用分层直接偏好优化（DPO）方法，结合蒙特卡洛树搜索（MCTS）实现质量评估的层级传递。实验表明，该模型在多项基准测试中超越更大规模的基线模型，验证了分层DPO的有效性，同时证明结构化思维步骤能显著提升长文本生成质量。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-05T21:01:34Z
- **目录日期**: 2025-06-05
