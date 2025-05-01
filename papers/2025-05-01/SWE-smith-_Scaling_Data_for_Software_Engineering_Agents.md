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

研究人员开发了SWE-smith工具，用于自动生成大规模软件工程训练数据。该工具能构建代码执行环境，并自动生成数千个破坏现有测试的任务实例。基于此，团队创建了包含5万个实例的数据集（来自128个GitHub仓库），规模远超现有数据集。训练出的SWE-agent-LM-32B模型在SWE-bench测试中达到40.2%的通过率，是目前开源模型中的最佳表现。所有资源已开源（包括收集流程、任务实例、轨迹和模型），以降低自动化软件工程研究的门槛。项目网址：https://swesmith.com。（99字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-01T15:01:35Z
- **目录日期**: 2025-05-01
