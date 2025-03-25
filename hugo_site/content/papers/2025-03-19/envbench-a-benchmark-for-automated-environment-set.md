---
title: "EnvBench: A Benchmark for Automated Environment Setup"
date: 2025-03-19T23:02:00+00:00
source_url: "http://arxiv.org/abs/2503.14443v1"
categories: ["2025-03-19"]
tags: ["Benchmark", "Automated", "Environment", "Setup", "EnvBench"]
summary: "近期大型语言模型（LLMs）的进展促使研究者关注软件工程领域的实际仓库级任务。本文聚焦于自动化软件仓库环境设置任务，提出了一个全面的环境设置基准EnvBench，涵盖329个Python和665个JVM（Java、Kotlin）仓库，重点关注具有真实配置挑战的项目。我们实现了两种自动评估指标：Python的静态分析检查和JVM语言的编译检查。通过评估三种环境设置方法，包括简单的零样本基线和两种代理工作流，使用GPT-4o和GPT-4o-mini进行测试，结果表明EnvBench对当前方法仍具挑战性。基准套件公开于https://github.com/JetBrains-Research/EnvBench。"
---

# EnvBench: A Benchmark for Automated Environment Setup

**原始链接**: [查看原文](http://arxiv.org/abs/2503.14443v1)

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

近期大型语言模型（LLMs）的进展促使研究者关注软件工程领域的实际仓库级任务。本文聚焦于自动化软件仓库环境设置任务，提出了一个全面的环境设置基准EnvBench，涵盖329个Python和665个JVM（Java、Kotlin）仓库，重点关注具有真实配置挑战的项目。我们实现了两种自动评估指标：Python的静态分析检查和JVM语言的编译检查。通过评估三种环境设置方法，包括简单的零样本基线和两种代理工作流，使用GPT-4o和GPT-4o-mini进行测试，结果表明EnvBench对当前方法仍具挑战性。基准套件公开于https://github.com/JetBrains-Research/EnvBench。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-19T23:02:03Z
- **目录日期**: 2025-03-19
