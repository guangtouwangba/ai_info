# Collab: Controlled Decoding using Mixture of Agents for LLM Alignment

**URL**: http://arxiv.org/abs/2503.21720v1

## 原始摘要

Alignment of Large Language models (LLMs) is crucial for safe and trustworthy
deployment in applications. Reinforcement learning from human feedback (RLHF)
has emerged as an effective technique to align LLMs to human preferences and
broader utilities, but it requires updating billions of model parameters, which
is computationally expensive. Controlled Decoding, by contrast, provides a
mechanism for aligning a model at inference time without retraining. However,
single-agent decoding approaches often struggle to adapt to diverse tasks due
to the complexity and variability inherent in these tasks. To strengthen the
test-time performance w.r.t the target task, we propose a mixture of
agent-based decoding strategies leveraging the existing off-the-shelf aligned
LLM policies. Treating each prior policy as an agent in the spirit of mixture
of agent collaboration, we develop a decoding method that allows for
inference-time alignment through a token-level selection strategy among
multiple agents. For each token, the most suitable LLM is dynamically chosen
from a pool of models based on a long-term utility metric. This
policy-switching mechanism ensures optimal model selection at each step,
enabling efficient collaboration and alignment among LLMs during decoding.
Theoretical analysis of our proposed algorithm establishes optimal performance
with respect to the target task represented via a target reward for the given
off-the-shelf models. We conduct comprehensive empirical evaluations with
open-source aligned models on diverse tasks and preferences, which demonstrates
the merits of this approach over single-agent decoding baselines. Notably,
Collab surpasses the current SoTA decoding strategy, achieving an improvement
of up to 1.56x in average reward and 71.89% in GPT-4 based win-tie rate.


## AI 摘要

该研究提出了一种基于多智能体协作的解码方法（Collab），用于提升大语言模型（LLM）的推理时对齐能力。传统强化学习人类反馈（RLHF）需高成本微调，而单智能体解码难以适应多样化任务。Collab通过动态选择预训练对齐模型池中的最优策略，在token级别进行智能体切换，以长期效用指标实现高效协作。理论分析证明其目标奖励最优性，实验显示其显著优于单智能体基线：平均奖励提升1.56倍，GPT-4评估胜率提高71.89%，成为当前最佳解码策略。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-29T04:01:35Z
- **目录日期**: 2025-03-29
