# Towards a Deeper Understanding of Reasoning Capabilities in Large Language Models

**URL**: http://arxiv.org/abs/2505.10543v1

## 原始摘要

While large language models demonstrate impressive performance on static
benchmarks, the true potential of large language models as self-learning and
reasoning agents in dynamic environments remains unclear. This study
systematically evaluates the efficacy of self-reflection, heuristic mutation,
and planning as prompting techniques to test the adaptive capabilities of
agents. We conduct experiments with various open-source language models in
dynamic environments and find that larger models generally outperform smaller
ones, but that strategic prompting can close this performance gap. Second, a
too-long prompt can negatively impact smaller models on basic reactive tasks,
while larger models show more robust behaviour. Third, advanced prompting
techniques primarily benefit smaller models on complex games, but offer less
improvement for already high-performing large language models. Yet, we find
that advanced reasoning methods yield highly variable outcomes: while capable
of significantly improving performance when reasoning and decision-making
align, they also introduce instability and can lead to big performance drops.
Compared to human performance, our findings reveal little evidence of true
emergent reasoning. Instead, large language model performance exhibits
persistent limitations in crucial areas such as planning, reasoning, and
spatial coordination, suggesting that current-generation large language models
still suffer fundamental shortcomings that may not be fully overcome through
self-reflective prompting alone. Reasoning is a multi-faceted task, and while
reasoning methods like Chain of thought improves multi-step reasoning on math
word problems, our findings using dynamic benchmarks highlight important
shortcomings in general reasoning capabilities, indicating a need to move
beyond static benchmarks to capture the complexity of reasoning.


## AI 摘要

这项研究评估了大型语言模型(LLM)在动态环境中的自适应能力，测试了自我反思、启发式突变和规划等提示技术的效果。实验发现：1)大模型通常优于小模型，但策略性提示可缩小差距；2)过长提示会损害小模型的基本反应能力，而大模型表现更稳定；3)高级提示技术对小模型在复杂任务中帮助更大。研究揭示了LLM在规划、推理和空间协调等关键领域仍存在根本性局限，当前的推理方法(如思维链)虽能改进多步数学推理，但在动态环境中表现出不稳定性，表明需要超越静态基准来全面评估推理能力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-16T17:01:12Z
- **目录日期**: 2025-05-16
