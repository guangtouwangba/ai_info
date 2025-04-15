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

研究人员提出了RealWebAssist基准测试，用于评估AI助手在真实场景中执行长期网页任务的能力。该测试基于真实用户的多步骤模糊指令，要求AI理解用户意图、跟踪其心理状态、适应个性化习惯，并准确操作网页界面。实验表明，当前最先进的AI模型在理解用户指令和界面操作方面仍面临重大挑战，难以满足现实世界中长期网页辅助的需求。该基准测试填补了现有评估工具的空白，强调了真实场景中模糊指令、动态需求和复杂交互对AI助手提出的更高要求。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-15T18:02:09Z
- **目录日期**: 2025-04-15
