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

这篇论文介绍了在AI数学奥林匹克竞赛AIMO-2中获胜的方案。其核心方法包括：1) 构建包含54万道高质量数学题(含奥赛题)和320万条详细解答的大规模数据集；2) 开发新方法，通过迭代训练、生成和质量筛选，将代码执行与长推理模型结合，产生170万条工具辅助的高质量解答；3) 建立流程训练模型从多个候选解答中选择最优解，证明生成式选择(GenSelect)能显著优于多数投票基准。综合这些方法训练出的模型在数学推理基准测试中达到最先进水平。作者开源了代码、模型和完整数据集。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-24T23:01:31Z
- **目录日期**: 2025-04-24
