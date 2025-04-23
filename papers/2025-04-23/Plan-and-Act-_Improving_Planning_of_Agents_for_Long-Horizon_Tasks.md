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

大型语言模型（LLMs）在简单任务中表现优异，但处理复杂多步骤任务仍具挑战。研究提出“规划-执行”（Plan-and-Act）框架，通过分离高层规划与底层执行来优化任务处理。该框架包含生成结构化高层计划的Planner模型和执行环境特定动作的Executor模型。为提高Planner的准确性，研究者开发了一种合成数据生成方法，通过标注真实轨迹数据并增强多样性来提升泛化能力。实验表明，该框架在WebArena-Lite和WebVoyager基准测试中分别达到57.58%和81.36%的成功率，表现优异。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-23T21:02:28Z
- **目录日期**: 2025-04-23
