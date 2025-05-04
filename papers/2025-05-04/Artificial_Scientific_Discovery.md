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

这篇论文探讨了如何实现"人工科学家"的愿景，即机器能自主开展原创研究并扩展人类知识。研究从AlphaGo类智能体Olivaw出发，发现其缺乏解释能力，进而提出"解释性学习"(EL)框架来模拟科学家解释新现象的过程。通过破解科学模拟游戏Zendo，研究发现人工科学家必须自主构建解释语言的理解机制。作者进一步分析多模态模型，提出解耦感知与解释的CLIP类模型构建方法。最后指出ChatGPT等模型在解释Zendo类规则时表现仅达随机水平，与人类存在显著差距，为此设计了Big-Bench符号解释任务作为基准测试。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-04T16:01:53Z
- **目录日期**: 2025-05-04
