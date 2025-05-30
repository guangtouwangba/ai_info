# Artificial Scientific Discovery

**URL**: http://arxiv.org/abs/2411.11672v2

## 原始摘要

Rooted in the explosion of deep learning over the past decade, this thesis
spans from AlphaGo to ChatGPT to empirically examine the fundamental concepts
needed to realize the vision of an artificial scientist: a machine with the
capacity to autonomously generate original research and contribute to the
expansion of human knowledge. The investigation begins with Olivaw, an AlphaGo
Zero-like agent that discovers Othello knowledge from scratch but is unable to
communicate it. This realization leads to the development of the Explanatory
Learning (EL) framework, a formalization of the problem faced by a scientist
when trying to explain a new phenomenon to their peers. The effective EL
prescriptions allow us to crack Zendo, a popular board game simulating the
scientific endeavor. This success comes with a fundamental insight: an
artificial scientist must develop its own interpretation of the language used
to explain its findings, and not rely on a rigid existing interpreter.
Questioning the very process of learning an interpreter, we turn our attention
to the inner functioning of modern multimodal models. This culminates in a
simple idea to build CLIP-like models where interpretation and perception are
explicitly disentangled: a cost-effective approach that couples two unimodal
models using little multimodal data and no further training. Finally, we
discuss what ChatGPT and its siblings are still missing to become artificial
scientists, and introduce the Big-Bench Symbol Interpretation Task, a benchmark
about interpreting Zendo-like explanations that sees LLMs going no further than
random chance while being instead fully solved by humans.


## AI 摘要

这篇论文探讨了实现"人工科学家"愿景所需的核心概念，从AlphaGo到ChatGPT展开实证研究。研究始于类似AlphaGo Zero的Olivaw智能体，它虽能自主发现知识但无法表达，由此提出了"解释性学习"(EL)框架。通过破解科学模拟游戏Zendo，研究发现人工科学家需要自主构建解释语言而非依赖固定解释器。作者进一步分析多模态模型，提出将感知与解释显式分离的CLIP类模型构建方法。最后指出当前大语言模型(如ChatGPT)与人工科学家的差距，并引入Big-Bench符号解释任务作为基准测试，显示人类表现远超随机水平而大模型仅达随机水平。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-05T01:29:38Z
- **目录日期**: 2025-05-05
