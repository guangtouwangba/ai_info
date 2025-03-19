# PLAY2PROMPT: Zero-shot Tool Instruction Optimization for LLM Agents via Tool Play

**URL**: http://arxiv.org/abs/2503.14432v1

## 原始摘要

Large language models (LLMs) are increasingly integrated with specialized
external tools, yet many tasks demand zero-shot tool usage with minimal or
noisy documentation. Existing solutions rely on manual rewriting or labeled
data for validation, making them inapplicable in true zero-shot settings. To
address these challenges, we propose PLAY2PROMPT, an automated framework that
systematically "plays" with each tool to explore its input-output behaviors.
Through this iterative trial-and-error process, PLAY2PROMPT refines tool
documentation and generates usage examples without any labeled data. These
examples not only guide LLM inference but also serve as validation to further
enhance tool utilization. Extensive experiments on real-world tasks demonstrate
that PLAY2PROMPT significantly improves zero-shot tool performance across both
open and closed models, offering a scalable and effective solution for
domain-specific tool integration.


## AI 摘要

大型语言模型（LLMs）越来越多地与专用外部工具集成，但许多任务需要在零样本设置下使用工具，且文档可能不完整或存在噪声。现有方法依赖手动重写或标注数据进行验证，难以适用于真正的零样本场景。为此，研究者提出了PLAY2PROMPT框架，通过自动化“试错”过程探索工具的输入输出行为，生成工具使用示例并优化文档，无需标注数据。实验表明，PLAY2PROMPT显著提升了开放和封闭模型在零样本工具使用中的表现，为领域专用工具集成提供了可扩展且高效的解决方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-19T15:03:02Z
- **目录日期**: 2025-03-19
