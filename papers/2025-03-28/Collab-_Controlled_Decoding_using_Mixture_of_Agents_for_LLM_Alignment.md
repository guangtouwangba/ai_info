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

本文提出了一种基于多智能体协作的解码策略Collab，用于提升大语言模型(LLM)在推理阶段的对齐性能。该方法将现有对齐模型视为不同智能体，在解码时动态选择最适合当前token的模型策略，通过长时效用指标实现token级模型切换。理论分析证明了该方法在给定目标奖励下的最优性。实验表明，Collab在开源对齐模型上显著优于单智能体基线，平均奖励提升1.56倍，GPT-4评估胜率提高71.89%，成为当前最先进的解码策略。这种方法避免了昂贵的模型微调，实现了高效的多模型协作对齐。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-28T03:13:31Z
- **目录日期**: 2025-03-28
