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

本文提出了一种名为EG-CFG（执行引导的无分类器指导）的新方法，将实时执行信号整合到语言模型的代码生成过程中。与现有方法不同，EG-CFG在生成代码时动态利用执行反馈，通过多阶段流程（候选程序采样、测试用例执行、信号整合）提供逐行指导。该方法保持了语法结构的一致性，并支持多代理并行探索不同推理路径。实验表明，EG-CFG在各种编码任务中显著提升了生成性能，从基础问题到复杂编程挑战均达到最先进水平。代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-13T06:02:00Z
- **目录日期**: 2025-06-13
