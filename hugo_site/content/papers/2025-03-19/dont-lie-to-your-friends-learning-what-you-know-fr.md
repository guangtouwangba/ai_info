---
title: "Don't lie to your friends: Learning what you know from collaborative self-play"
date: 2025-03-19T23:01:00+00:00
source_url: "http://arxiv.org/abs/2503.14481v1"
categories: ["2025-03-19"]
tags: ["Don't", "Learning"]
summary: "本文提出了一种名为“协作自玩”的新方法，用于教导AI代理了解自身能力和限制。通过构建多代理协作环境，团队因共同得出正确答案而获得奖励，从而促使代理发展出必要的元知识。这种方法特别适用于拥有不同工具（如特定语料库检索）的小型代理群体，它们必须协作以最大化成功并最小化努力。实验表明，多代理群体的集体奖励能够诱导出在单独部署时也能改善工具使用和选择性预测的策略。这种方法避免了监督微调中难以构建反映代理特定能力的示例的问题。"
---

# Don't lie to your friends: Learning what you know from collaborative self-play

**原始链接**: [查看原文](http://arxiv.org/abs/2503.14481v1)

## 原始摘要

To be helpful assistants, AI agents must be aware of their own capabilities
and limitations. This includes knowing when to answer from parametric knowledge
versus using tools, when to trust tool outputs, and when to abstain or hedge.
Such capabilities are hard to teach through supervised fine-tuning because they
require constructing examples that reflect the agent's specific capabilities.
We therefore propose a radically new approach to teaching agents what they
know: \emph{collaborative self-play}. We construct multi-agent collaborations
in which the group is rewarded for collectively arriving at correct answers.
The desired meta-knowledge emerges from the incentives built into the structure
of the interaction. We focus on small societies of agents that have access to
heterogeneous tools (corpus-specific retrieval), and therefore must collaborate
to maximize their success while minimizing their effort. Experiments show that
group-level rewards for multi-agent communities can induce policies that
\emph{transfer} to improve tool use and selective prediction in settings where
individual agents are deployed in isolation.

## AI 摘要

本文提出了一种名为“协作自玩”的新方法，用于教导AI代理了解自身能力和限制。通过构建多代理协作环境，团队因共同得出正确答案而获得奖励，从而促使代理发展出必要的元知识。这种方法特别适用于拥有不同工具（如特定语料库检索）的小型代理群体，它们必须协作以最大化成功并最小化努力。实验表明，多代理群体的集体奖励能够诱导出在单独部署时也能改善工具使用和选择性预测的策略。这种方法避免了监督微调中难以构建反映代理特定能力的示例的问题。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-19T23:01:52Z
- **目录日期**: 2025-03-19
