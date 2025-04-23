# Plan-and-Act: Improving Planning of Agents for Long-Horizon Tasks

**URL**: http://arxiv.org/abs/2503.09572v3

## 原始摘要

Large language models (LLMs) have shown remarkable advancements in enabling
language agents to tackle simple tasks. However, applying them for complex,
multi-step, long-horizon tasks remains a challenge. Recent work have found
success by separating high-level planning from low-level execution, which
enables the model to effectively balance high-level planning objectives and
low-level execution details. However, generating accurate plans remains
difficult since LLMs are not inherently trained for this task. To address this,
we propose Plan-and-Act, a novel framework that incorporates explicit planning
into LLM-based agents and introduces a scalable method to enhance plan
generation through a novel synthetic data generation method. Plan-and-Act
consists of a Planner model which generates structured, high-level plans to
achieve user goals, and an Executor model that translates these plans into
environment-specific actions. To train the Planner effectively, we introduce a
synthetic data generation method that annotates ground-truth trajectories with
feasible plans, augmented with diverse and extensive examples to enhance
generalization. We evaluate Plan-and-Act using web navigation as a
representative long-horizon planning environment, demonstrating a
state-of-the-art 57.58% success rate on the WebArena-Lite benchmark as well as
a text-only state-of-the-art 81.36% success rate on WebVoyager.


## AI 摘要

本文提出了一种名为Plan-and-Act的新型框架，旨在解决大型语言模型（LLMs）在复杂多步骤任务中的规划难题。该框架通过分离高层次规划（Planner模型）和低层次执行（Executor模型），并采用创新的合成数据生成方法来增强规划能力。实验表明，Plan-and-Act在WebArena-Lite基准测试中达到57.58%的成功率，在纯文本环境下（WebVoyager）达到81.36%的成功率，均达到当前最优水平。该方法通过结构化规划和环境特定动作转换，显著提升了LLMs在长期任务中的表现。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-23T17:02:48Z
- **目录日期**: 2025-04-23
