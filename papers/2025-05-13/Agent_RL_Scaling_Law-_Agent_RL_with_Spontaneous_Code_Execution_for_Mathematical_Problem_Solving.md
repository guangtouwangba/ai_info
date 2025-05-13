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

该研究探索了基于结果奖励的强化学习(RL)如何让大语言模型(LLMs)自主学会使用外部工具(如Python代码)进行数学推理(ZeroTIR方法)。研究发现，随着RL训练推进，关键指标呈现可预测的线性增长：代码执行频率、响应长度和任务准确率同步提升。该方法在数学基准测试上显著优于不使用工具的基线模型。研究实现了一个包含独立代码执行环境的框架，验证了工具使用能力与训练计算量之间的量化关系，为自主工具学习机制提供了可复现的基准。代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-13T12:02:08Z
- **目录日期**: 2025-05-13
