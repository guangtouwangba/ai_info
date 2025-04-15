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

RealWebAssist是一个评估AI代理在真实场景中执行长期网页任务能力的新基准。与现有基准不同，它聚焦于处理模糊、动态变化的用户指令，要求代理理解用户真实意图、跟踪其心理状态并准确执行多网站任务。该基准包含来自真实用户的连续指令数据集，挑战包括视觉GUI定位、模糊指令理解和长期交互。实验表明，当前最先进模型在理解用户意图和GUI操作上仍存在困难，突显了实现长期网页辅助的关键挑战。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-15T07:01:36Z
- **目录日期**: 2025-04-15
