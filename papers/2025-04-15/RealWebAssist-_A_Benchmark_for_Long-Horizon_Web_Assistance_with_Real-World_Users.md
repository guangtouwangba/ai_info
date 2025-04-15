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

研究人员提出了RealWebAssist基准测试，用于评估AI助手在真实场景中处理长期网页任务的能力。该测试针对现实用户指令的模糊性、多变性以及需要持续理解用户心理状态等挑战而设计。数据集包含真实用户在多网站上的连续任务指令，要求AI能准确理解意图、跟踪用户心理变化、识别个人习惯，并将任务正确映射到GUI操作。实验表明，当前最先进的模型在理解用户指令和GUI操作方面仍存在困难，突显了长期网页辅助任务中的关键挑战。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-15T20:02:13Z
- **目录日期**: 2025-04-15
