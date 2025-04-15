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

这篇论文介绍了RealWebAssist基准测试，用于评估AI代理在真实网络环境中执行长期连续指令的能力。与现有基准不同，它专注于处理现实场景中的模糊指令、用户心理状态变化和多网站任务。该基准包含真实用户提供的连续指令数据集，要求AI理解用户意图、跟踪心理状态、适应个人习惯并准确操作GUI元素。实验表明，当前先进模型在理解指令和GUI操作方面仍面临重大挑战，突显了长期网络辅助任务中的关键难点。该研究填补了真实场景下连续指令跟随评估的空白。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-15T14:01:46Z
- **目录日期**: 2025-04-15
