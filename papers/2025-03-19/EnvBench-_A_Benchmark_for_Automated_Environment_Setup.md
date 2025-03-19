# EnvBench: A Benchmark for Automated Environment Setup

**URL**: http://arxiv.org/abs/2503.14443v1

## 原始摘要

Recent advances in Large Language Models (LLMs) have enabled researchers to
focus on practical repository-level tasks in software engineering domain. In
this work, we consider a cornerstone task for automating work with software
repositories-environment setup, i.e., a task of configuring a
repository-specific development environment on a system. Existing studies on
environment setup introduce innovative agentic strategies, but their evaluation
is often based on small datasets that may not capture the full range of
configuration challenges encountered in practice. To address this gap, we
introduce a comprehensive environment setup benchmark EnvBench. It encompasses
329 Python and 665 JVM-based (Java, Kotlin) repositories, with a focus on
repositories that present genuine configuration challenges, excluding projects
that can be fully configured by simple deterministic scripts. To enable further
benchmark extension and usage for model tuning, we implement two automatic
metrics: a static analysis check for missing imports in Python and a
compilation check for JVM languages. We demonstrate the applicability of our
benchmark by evaluating three environment setup approaches, including a simple
zero-shot baseline and two agentic workflows, that we test with two powerful
LLM backbones, GPT-4o and GPT-4o-mini. The best approach manages to
successfully configure 6.69% repositories for Python and 29.47% repositories
for JVM, suggesting that EnvBench remains challenging for current approaches.
Our benchmark suite is publicly available at
https://github.com/JetBrains-Research/EnvBench. The dataset and experiment
trajectories are available at https://jb.gg/envbench.


## AI 摘要

近期大型语言模型（LLMs）的进展使研究者能够专注于软件工程领域的实际仓库级任务。本文提出EnvBench，一个全面的环境设置基准，涵盖329个Python和665个基于JVM（Java、Kotlin）的仓库，专注于真实配置挑战。EnvBench引入两种自动评估指标：Python的静态分析检查和JVM语言的编译检查。通过评估三种环境设置方法（包括零样本基线和两种代理工作流），使用GPT-4o和GPT-4o-mini进行测试，最佳方法成功配置了6.69%的Python仓库和29.47%的JVM仓库，表明EnvBench对现有方法仍具挑战性。数据集和实验轨迹公开于https://jb.gg/envbench。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-19T11:49:37Z
- **目录日期**: 2025-03-19
