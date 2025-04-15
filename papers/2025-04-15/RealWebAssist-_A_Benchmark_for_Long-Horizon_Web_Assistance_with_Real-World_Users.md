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

这篇论文介绍了RealWebAssist，一个评估AI代理在现实场景中执行长期网页任务能力的新基准。该基准针对真实用户指令的模糊性、多阶段演变和不同辅助需求等挑战，要求代理理解用户意图、跟踪心理状态、识别用户习惯并准确定位GUI操作。基于真实用户指令收集的数据集显示，当前最先进模型在理解用户指令和界面操作方面仍存在困难，突显了长期网页辅助任务的关键挑战。该研究填补了现有网页代理测试在连续指令执行和真实场景适应性方面的空白。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-15T11:01:27Z
- **目录日期**: 2025-04-15
