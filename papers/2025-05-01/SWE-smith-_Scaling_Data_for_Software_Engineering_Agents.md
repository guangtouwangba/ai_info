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

当前语言模型在软件工程中的应用面临训练数据收集困难的问题，现有数据集规模小且构建成本高。为解决这一问题，研究者提出了SWE-smith，一种自动化生成大规模软件工程训练数据的流程。该方法能基于任意Python代码库构建执行环境，并自动合成数百至数千个破坏现有测试的任务实例。利用该工具，团队创建了包含5万个实例的数据集，规模远超以往研究。基于此训练的SWE-agent-LM-32B模型在SWE-bench基准测试中达到40.2%的通过率，是目前开源模型中的最佳表现。所有资源已开源以促进相关研究。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-01T21:01:36Z
- **目录日期**: 2025-05-01
