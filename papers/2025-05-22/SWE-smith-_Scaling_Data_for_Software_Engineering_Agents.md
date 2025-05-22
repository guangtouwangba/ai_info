# SWE-smith: Scaling Data for Software Engineering Agents

**URL**: http://arxiv.org/abs/2504.21798v2

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

该论文介绍了SWE-smith，一种新型软件工程训练数据生成管道，解决了现有数据集规模小（通常仅数千条数据）、构建复杂且存储需求大的问题。该工具能自动为任意Python代码库构建执行环境，并生成数百至数千个破坏现有测试的任务实例。基于此，研究团队创建了包含5万实例（来自128个GitHub仓库）的数据集，规模远超以往。训练出的SWE-agent-LM-32B模型在SWE-bench基准测试中达到40.2%的Pass@1解决率，成为开源模型的新标杆。所有资源已开源以促进自动化软件工程研究。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-22T20:02:37Z
- **目录日期**: 2025-05-22
