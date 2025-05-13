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

该研究探讨了基于结果奖励的强化学习(RL)如何让大语言模型(LLMs)自主学会使用外部工具(如Python代码执行)进行数学推理。提出的ZeroTIR方法无需监督示例，就能训练模型自发生成并执行代码。研究发现，随着RL训练推进，代码执行频率、响应长度和任务准确率呈正相关增长，表明训练计算量与工具增强推理策略的涌现存在量化关系。实验显示，ZeroTIR在数学基准测试上显著优于非工具方法。该工作为理解自主工具使用的习得机制提供了基础，并建立了可复现的基准框架。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-13T11:02:02Z
- **目录日期**: 2025-05-13
