# Is the Reversal Curse a Binding Problem? Uncovering Limitations of Transformers from a Basic Generalization Failure

**URL**: http://arxiv.org/abs/2504.01928v1

## 原始摘要

Despite their impressive capabilities, LLMs exhibit a basic generalization
failure known as the Reversal Curse, where they struggle to learn reversible
factual associations. Understanding why this occurs could help identify
weaknesses in current models and advance their generalization and robustness.
In this paper, we conjecture that the Reversal Curse in LLMs is a manifestation
of the long-standing binding problem in cognitive science, neuroscience and AI.
Specifically, we identify two primary causes of the Reversal Curse stemming
from transformers' limitations in conceptual binding: the inconsistency and
entanglements of concept representations. We perform a series of experiments
that support these conjectures. Our exploration leads to a model design based
on JEPA (Joint-Embedding Predictive Architecture) that for the first time
breaks the Reversal Curse without side-stepping it with specialized data
augmentation or non-causal masking, and moreover, generalization could be
further improved by incorporating special memory layers that support
disentangled concept representations. We demonstrate that the skill of reversal
unlocks a new kind of memory integration that enables models to solve
large-scale arithmetic reasoning problems via parametric forward-chaining,
outperforming frontier LLMs based on non-parametric memory and prolonged
explicit reasoning.


## AI 摘要

大型语言模型(LLMs)存在"逆转诅咒"(Reversal Curse)的泛化缺陷，即难以学习可逆的事实关联。研究表明，这一现象源于Transformer在概念绑定上的局限性，表现为概念表示的不一致性和纠缠性。通过实验验证后，研究者提出基于JEPA(联合嵌入预测架构)的新模型设计，首次在不依赖数据增强或非因果掩码的情况下解决了该问题。进一步引入支持解耦概念表示的特殊记忆层可提升泛化能力。突破逆转诅咒使模型能通过参数化前向链解决大规模算术推理问题，性能超越前沿LLMs。该发现为改进模型泛化和鲁棒性提供了新思路。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-04T00:02:00Z
- **目录日期**: 2025-04-04
