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

这篇论文介绍了在AI数学奥林匹克竞赛（AIMO-2）中获胜的解决方案。其核心方法基于三大支柱：1）构建包含54万道高质量数学题（含奥赛题）及其320万条长推理解答的大规模数据集；2）通过迭代训练、生成和质量筛选，开发出融合代码执行与长推理模型的新方法，产生170万条工具集成推理解答；3）建立流程训练模型从多个候选解中选出最优解，证明生成式选择（GenSelect）显著优于多数投票基线。最终训练出的模型在数学推理基准测试中达到SOTA，并开源了代码、模型和OpenMathReasoning数据集。（99字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-24T14:01:49Z
- **目录日期**: 2025-04-24
