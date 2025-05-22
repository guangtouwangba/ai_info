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

SWE-smith是一个创新的自动化流程，用于大规模生成软件工程训练数据。它通过构建代码执行环境并自动合成数百至数千个破坏现有测试的任务实例，解决了现有数据集规模小（通常仅数千个实例）、构建复杂且存储需求高的问题。利用SWE-smith，研究者创建了包含5万个实例的数据集（来自128个GitHub仓库），规模远超以往。基于此训练的SWE-agent-LM-32B模型在SWE-bench基准测试中达到40.2%的Pass@1率，成为开源模型的新标杆。项目已开源，旨在降低自动化软件工程研究的门槛。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-22T22:02:32Z
- **目录日期**: 2025-05-22
