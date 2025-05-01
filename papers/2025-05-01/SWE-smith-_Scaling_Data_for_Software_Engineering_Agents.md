# SWE-smith: Scaling Data for Software Engineering Agents

**URL**: http://arxiv.org/abs/2504.21798v1

## 原始摘要

Despite recent progress in Language Models (LMs) for software engineering,
collecting training data remains a significant pain point. Existing datasets
are small, with at most 1,000s of training instances from 11 or fewer GitHub
repositories. The procedures to curate such datasets are often complex,
necessitating hundreds of hours of human labor; companion execution
environments also take up several terabytes of storage, severely limiting their
scalability and usability. To address this pain point, we introduce SWE-smith,
a novel pipeline for generating software engineering training data at scale.
Given any Python codebase, SWE-smith constructs a corresponding execution
environment, then automatically synthesizes 100s to 1,000s of task instances
that break existing test(s) in the codebase. Using SWE-smith, we create a
dataset of 50k instances sourced from 128 GitHub repositories, an order of
magnitude larger than all previous works. We train SWE-agent-LM-32B, achieving
40.2% Pass@1 resolve rate on the SWE-bench Verified benchmark, state of the art
among open source models. We open source SWE-smith (collection procedure, task
instances, trajectories, models) to lower the barrier of entry for research in
LM systems for automated software engineering. All assets available at
https://swesmith.com.


## AI 摘要

当前语言模型在软件工程中的应用面临训练数据收集困难的问题。现有数据集规模小（最多数千条）、来源有限（≤11个GitHub仓库），且构建过程复杂耗时。为此，研究者提出SWE-smith自动化管道，能基于任意Python代码库生成执行环境并自动合成数百至数千个破坏现有测试的任务实例。利用该工具，团队创建了包含128个仓库5万实例的大规模数据集（比现有工作大一个数量级），并训练出SWE-agent-LM-32B模型，在SWE-bench基准测试中达到40.2%的Pass@1解决率（开源模型最佳）。所有资源已开源以降低研究门槛。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-01T08:01:38Z
- **目录日期**: 2025-05-01
