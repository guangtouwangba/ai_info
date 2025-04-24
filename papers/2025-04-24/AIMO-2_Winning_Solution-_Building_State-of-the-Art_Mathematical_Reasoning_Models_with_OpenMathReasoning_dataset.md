# AIMO-2 Winning Solution: Building State-of-the-Art Mathematical Reasoning Models with OpenMathReasoning dataset

**URL**: http://arxiv.org/abs/2504.16891v1

## 原始摘要

This paper presents our winning submission to the AI Mathematical Olympiad -
Progress Prize 2 (AIMO-2) competition. Our recipe for building state-of-the-art
mathematical reasoning models relies on three key pillars. First, we create a
large-scale dataset comprising 540K unique high-quality math problems,
including olympiad-level problems, and their 3.2M long-reasoning solutions.
Second, we develop a novel method to integrate code execution with long
reasoning models through iterative training, generation, and quality filtering,
resulting in 1.7M high-quality Tool-Integrated Reasoning solutions. Third, we
create a pipeline to train models to select the most promising solution from
many candidates. We show that such generative solution selection (GenSelect)
can significantly improve upon majority voting baseline. Combining these ideas,
we train a series of models that achieve state-of-the-art results on
mathematical reasoning benchmarks. To facilitate further research, we release
our code, models, and the complete OpenMathReasoning dataset under a
commercially permissive license.


## AI 摘要

这篇论文介绍了在AI数学奥林匹克竞赛AIMO-2中获胜的方案，其核心基于三大创新：1)构建了包含54万高质量数学题(含奥数题)和320万详细解答的大规模数据集；2)开发了通过迭代训练、生成和质量筛选将代码执行与长推理模型结合的新方法，产生了170万工具集成推理解答；3)建立了从多个候选解中选择最优解的流程，证明生成式选择(GenSelect)能显著优于多数投票基准。最终训练出的模型在数学推理基准测试中达到SOTA水平。作者开源了代码、模型和完整数据集以促进后续研究。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-24T22:01:43Z
- **目录日期**: 2025-04-24
