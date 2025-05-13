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

这篇论文研究了大型语言模型(LLMs)在数学推理任务中的表现，提出了一种名为ZeroTIR的强化学习方法，让模型能自主生成并执行Python代码来解决数学问题，而无需监督学习工具使用示例。研究发现，随着强化学习训练的进行，关键指标(如代码执行频率、响应长度和任务准确率)呈现可预测的增长趋势。实验表明，ZeroTIR在数学基准测试上显著优于不使用工具的方法。该研究为理解自主工具使用的习得机制提供了基础，并建立了一个可复现的基准。相关代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-13T05:02:04Z
- **目录日期**: 2025-05-13
