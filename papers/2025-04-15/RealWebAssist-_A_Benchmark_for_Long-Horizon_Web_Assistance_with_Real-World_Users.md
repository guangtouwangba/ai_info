# RealWebAssist: A Benchmark for Long-Horizon Web Assistance with Real-World Users

**URL**: http://arxiv.org/abs/2504.10445v1

## 原始摘要

To achieve successful assistance with long-horizon web-based tasks, AI agents
must be able to sequentially follow real-world user instructions over a long
period. Unlike existing web-based agent benchmarks, sequential instruction
following in the real world poses significant challenges beyond performing a
single, clearly defined task. For instance, real-world human instructions can
be ambiguous, require different levels of AI assistance, and may evolve over
time, reflecting changes in the user's mental state. To address this gap, we
introduce RealWebAssist, a novel benchmark designed to evaluate sequential
instruction-following in realistic scenarios involving long-horizon
interactions with the web, visual GUI grounding, and understanding ambiguous
real-world user instructions. RealWebAssist includes a dataset of sequential
instructions collected from real-world human users. Each user instructs a
web-based assistant to perform a series of tasks on multiple websites. A
successful agent must reason about the true intent behind each instruction,
keep track of the mental state of the user, understand user-specific routines,
and ground the intended tasks to actions on the correct GUI elements. Our
experimental results show that state-of-the-art models struggle to understand
and ground user instructions, posing critical challenges in following
real-world user instructions for long-horizon web assistance.


## AI 摘要

研究人员提出了RealWebAssist基准测试，用于评估AI助手在真实场景中执行长期网页任务的能力。该基准测试包含来自真实用户的连续指令数据集，要求AI理解模糊指令、跟踪用户心理状态、适应个性化习惯并准确操作网页界面。研究表明，当前最先进的AI模型在理解用户意图、执行长期网页交互任务方面仍面临重大挑战，难以满足现实世界中持续演变的用户需求。该测试填补了现有网页代理评估的空白，突出了真实场景中顺序指令跟随的复杂性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-15T09:01:33Z
- **目录日期**: 2025-04-15
