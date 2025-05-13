# Agent RL Scaling Law: Agent RL with Spontaneous Code Execution for Mathematical Problem Solving

**URL**: http://arxiv.org/abs/2505.07773v1

## 原始摘要

Large Language Models (LLMs) often struggle with mathematical reasoning tasks
requiring precise, verifiable computation. While Reinforcement Learning (RL)
from outcome-based rewards enhances text-based reasoning, understanding how
agents autonomously learn to leverage external tools like code execution
remains crucial. We investigate RL from outcome-based rewards for
Tool-Integrated Reasoning, ZeroTIR, training base LLMs to spontaneously
generate and execute Python code for mathematical problems without supervised
tool-use examples. Our central contribution is we demonstrate that as RL
training progresses, key metrics scale predictably. Specifically, we observe
strong positive correlations where increased training steps lead to increases
in the spontaneous code execution frequency, the average response length, and,
critically, the final task accuracy. This suggests a quantifiable relationship
between computational effort invested in training and the emergence of
effective, tool-augmented reasoning strategies. We implement a robust framework
featuring a decoupled code execution environment and validate our findings
across standard RL algorithms and frameworks. Experiments show ZeroTIR
significantly surpasses non-tool ZeroRL baselines on challenging math
benchmarks. Our findings provide a foundational understanding of how autonomous
tool use is acquired and scales within Agent RL, offering a reproducible
benchmark for future studies. Code is released at
\href{https://github.com/Anonymize-Author/AgentRL}{https://github.com/Anonymize-Author/AgentRL}.


## AI 摘要

这篇论文研究了基于强化学习（RL）的大语言模型（LLMs）如何自主学会使用外部工具（如Python代码执行）来解决数学推理问题。提出的ZeroTIR方法无需监督示例，通过RL训练使LLM能自发生成并执行代码。研究发现训练步数与代码执行频率、响应长度及任务准确率呈正相关，表明计算投入与工具增强推理策略之间存在量化关系。实验显示ZeroTIR在数学基准测试上显著优于非工具基线，为自主工具使用的学习机制提供了可复现的基准。代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-13T23:01:52Z
- **目录日期**: 2025-05-13
