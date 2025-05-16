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

该研究评估了大语言模型(LLM)在动态环境中的自适应能力，比较了不同规模模型在自反思、启发式突变和规划等提示技术下的表现。研究发现：1)大模型普遍优于小模型，但策略性提示可缩小差距；2)过长提示会损害小模型在基础任务中的表现；3)高级提示技术主要提升小模型在复杂任务中的表现，但对高性能大模型改善有限。研究指出当前LLM在规划、推理和空间协调等关键领域仍存在根本性缺陷，仅靠自反思提示难以完全克服。动态基准测试揭示了静态测试无法捕捉的推理能力局限性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-16T10:01:18Z
- **目录日期**: 2025-05-16
