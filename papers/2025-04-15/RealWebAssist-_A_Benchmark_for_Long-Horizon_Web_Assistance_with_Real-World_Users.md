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

研究人员提出了RealWebAssist基准测试，用于评估AI助手在真实网络场景中执行连续指令的能力。该基准包含从真实用户收集的序列指令数据集，要求AI在多网站环境中完成长周期任务，同时解决模糊指令理解、用户心理状态跟踪、个性化习惯识别和GUI元素精准定位等挑战。实验表明，当前最先进的AI模型在理解用户真实意图、保持交互一致性以及准确执行网络操作方面仍存在显著不足。该研究突出了开发能适应动态用户需求的网络助手的核心难点，为未来AI助手的发展指明了关键方向。（99字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-15T16:01:34Z
- **目录日期**: 2025-04-15
