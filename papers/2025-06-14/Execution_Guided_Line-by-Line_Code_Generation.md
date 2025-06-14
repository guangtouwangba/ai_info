# Execution Guided Line-by-Line Code Generation

**URL**: http://arxiv.org/abs/2506.10948v1

## 原始摘要

We present a novel approach to neural code generation that incorporates
real-time execution signals into the language model generation process. While
large language models (LLMs) have demonstrated impressive code generation
capabilities, they typically do not utilize execution feedback during
inference, a critical signal that human programmers regularly leverage. Our
method, Execution-Guided Classifier-Free Guidance (EG-CFG), dynamically
incorporates execution signals as the model generates code, providing
line-by-line feedback that guides the generation process toward executable
solutions. EG-CFG employs a multi-stage process: first, we conduct beam search
to sample candidate program completions for each line; second, we extract
execution signals by executing these candidates against test cases; and
finally, we incorporate these signals into the prompt during generation. By
maintaining consistent signals across tokens within the same line and
refreshing signals at line boundaries, our approach provides coherent guidance
while preserving syntactic structure. Moreover, the method naturally supports
native parallelism at the task level in which multiple agents operate in
parallel, exploring diverse reasoning paths and collectively generating a broad
set of candidate solutions. Our experiments across diverse coding tasks
demonstrate that EG-CFG significantly improves code generation performance
compared to standard approaches, achieving state-of-the-art results across
various levels of complexity, from foundational problems to challenging
competitive programming tasks. Our code is available at:
https://github.com/boazlavon/eg_cfg


## AI 摘要

本文提出了一种名为"执行引导的无分类器指导"(EG-CFG)的新型神经代码生成方法。该方法在语言模型生成过程中实时整合执行反馈信号，通过多阶段流程：首先生成候选代码行，执行测试用例获取反馈，然后将这些信号动态融入生成过程。EG-CFG在保持语法结构的同时提供连贯指导，并天然支持多智能体并行探索不同推理路径。实验表明，该方法显著提升了代码生成性能，在各种复杂度的编程任务中均达到最先进水平，从基础问题到具有挑战性的竞赛编程任务都表现出色。代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-14T03:20:05Z
- **目录日期**: 2025-06-14
