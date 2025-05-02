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

当前软件工程语言模型面临训练数据稀缺问题，现有数据集规模小（通常≤1000条/11个仓库）、构建成本高（数百小时人工+TB级存储）。为此，研究团队提出SWE-smith自动化流水线，能基于任意Python代码库构建执行环境并自动生成数百至数千个破坏原有测试的任务实例。利用该工具创建了含5万实例/128个仓库的新数据集（规模达先前10倍），并训练出SWE-agent-LM-32B模型，在SWE-bench测试中实现40.2%的Pass@1解析率（开源模型最佳）。所有资源已开源以降低研究门槛。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-02T01:28:55Z
- **目录日期**: 2025-05-02
