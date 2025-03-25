---
title: "Value Profiles for Encoding Human Variation"
date: 2025-03-20T13:01:00+00:00
source_url: "http://arxiv.org/abs/2503.15484v1"
categories: ["2025-03-20"]
tags: ["Encoding", "Human", "Variation", "Value", "Profiles"]
summary: "本研究提出了一种通过价值档案（自然语言描述）和可调控解码器模型来建模人类评分任务中的个体差异的方法。价值档案通过上下文演示压缩生成，能够有效保留超过70%的有用信息。研究发现，演示包含最多信息，其次是价值档案和人口统计信息。价值档案在可审查性、可解释性和可调控性方面具有优势，且通过聚类价值档案能更好地解释评分者差异。解码器模型能够根据语义差异调整评分，具有良好的校准性，并能模拟注释者群体以解释实例级分歧。价值档案为描述个体差异提供了新的预测方法。"
---

# Value Profiles for Encoding Human Variation

**原始链接**: [查看原文](http://arxiv.org/abs/2503.15484v1)

## 原始摘要

Modelling human variation in rating tasks is crucial for enabling AI systems
for personalization, pluralistic model alignment, and computational social
science. We propose representing individuals using value profiles -- natural
language descriptions of underlying values compressed from in-context
demonstrations -- along with a steerable decoder model to estimate ratings
conditioned on a value profile or other rater information. To measure the
predictive information in rater representations, we introduce an
information-theoretic methodology. We find that demonstrations contain the most
information, followed by value profiles and then demographics. However, value
profiles offer advantages in terms of scrutability, interpretability, and
steerability due to their compressed natural language format. Value profiles
effectively compress the useful information from demonstrations (&gt;70%
information preservation). Furthermore, clustering value profiles to identify
similarly behaving individuals better explains rater variation than the most
predictive demographic groupings. Going beyond test set performance, we show
that the decoder models interpretably change ratings according to semantic
profile differences, are well-calibrated, and can help explain instance-level
disagreement by simulating an annotator population. These results demonstrate
that value profiles offer novel, predictive ways to describe individual
variation beyond demographics or group information.

## AI 摘要

本研究提出了一种通过价值档案（自然语言描述）和可调控解码器模型来建模人类评分任务中的个体差异的方法。价值档案通过上下文演示压缩生成，能够有效保留超过70%的有用信息。研究发现，演示包含最多信息，其次是价值档案和人口统计信息。价值档案在可审查性、可解释性和可调控性方面具有优势，且通过聚类价值档案能更好地解释评分者差异。解码器模型能够根据语义差异调整评分，具有良好的校准性，并能模拟注释者群体以解释实例级分歧。价值档案为描述个体差异提供了新的预测方法。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-20T13:01:29+08:00
- **目录日期**: 2025-03-20
