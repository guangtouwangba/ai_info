---
title: "SWEET-RL: Training Multi-Turn LLM Agents on Collaborative Reasoning Tasks"
date: 2025-03-20T13:01:00+00:00
source_url: "http://arxiv.org/abs/2503.15478v1"
categories: ["2025-03-20"]
tags: ["Training", "Multi-Turn", "LLM", "Agents", "Collaborative"]
summary: "大型语言模型（LLM）代理在现实任务中需进行多轮交互，但现有多轮强化学习（RL）算法难以有效分配多轮信用并利用LLM的泛化能力。为此，研究团队提出了新基准ColBench，模拟LLM代理与人类协作解决后端编程和前端设计任务。基于此，他们开发了SWEET-RL算法，通过训练时信息优化目标训练评论家模型，提供步骤级奖励以改进策略模型。实验表明，SWEET-RL在ColBench上成功率提升6%，使Llama-3.1-8B在协作内容创作中表现媲美或超越GPT4-o。"
---

# SWEET-RL: Training Multi-Turn LLM Agents on Collaborative Reasoning Tasks

**原始链接**: [查看原文](http://arxiv.org/abs/2503.15478v1)

## 原始摘要

Large language model (LLM) agents need to perform multi-turn interactions in
real-world tasks. However, existing multi-turn RL algorithms for optimizing LLM
agents fail to perform effective credit assignment over multiple turns while
leveraging the generalization capabilities of LLMs and it remains unclear how
to develop such algorithms. To study this, we first introduce a new benchmark,
ColBench, where an LLM agent interacts with a human collaborator over multiple
turns to solve realistic tasks in backend programming and frontend design.
Building on this benchmark, we propose a novel RL algorithm, SWEET-RL (RL with
Step-WisE Evaluation from Training-time information), that uses a carefully
designed optimization objective to train a critic model with access to
additional training-time information. The critic provides step-level rewards
for improving the policy model. Our experiments demonstrate that SWEET-RL
achieves a 6% absolute improvement in success and win rates on ColBench
compared to other state-of-the-art multi-turn RL algorithms, enabling
Llama-3.1-8B to match or exceed the performance of GPT4-o in realistic
collaborative content creation.

## AI 摘要

大型语言模型（LLM）代理在现实任务中需进行多轮交互，但现有多轮强化学习（RL）算法难以有效分配多轮信用并利用LLM的泛化能力。为此，研究团队提出了新基准ColBench，模拟LLM代理与人类协作解决后端编程和前端设计任务。基于此，他们开发了SWEET-RL算法，通过训练时信息优化目标训练评论家模型，提供步骤级奖励以改进策略模型。实验表明，SWEET-RL在ColBench上成功率提升6%，使Llama-3.1-8B在协作内容创作中表现媲美或超越GPT4-o。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-20T13:01:40+08:00
- **目录日期**: 2025-03-20
