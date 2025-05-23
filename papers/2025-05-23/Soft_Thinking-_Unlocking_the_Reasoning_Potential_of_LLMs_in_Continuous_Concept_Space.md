# Soft Thinking: Unlocking the Reasoning Potential of LLMs in Continuous Concept Space

**URL**: http://arxiv.org/abs/2505.15778v1

## 原始摘要

Human cognition typically involves thinking through abstract, fluid concepts
rather than strictly using discrete linguistic tokens. Current reasoning
models, however, are constrained to reasoning within the boundaries of human
language, processing discrete token embeddings that represent fixed points in
the semantic space. This discrete constraint restricts the expressive power and
upper potential of such reasoning models, often causing incomplete exploration
of reasoning paths, as standard Chain-of-Thought (CoT) methods rely on sampling
one token per step. In this work, we introduce Soft Thinking, a training-free
method that emulates human-like "soft" reasoning by generating soft, abstract
concept tokens in a continuous concept space. These concept tokens are created
by the probability-weighted mixture of token embeddings, which form the
continuous concept space, enabling smooth transitions and richer
representations that transcend traditional discrete boundaries. In essence,
each generated concept token encapsulates multiple meanings from related
discrete tokens, implicitly exploring various reasoning paths to converge
effectively toward the correct answer. Empirical evaluations on diverse
mathematical and coding benchmarks consistently demonstrate the effectiveness
and efficiency of Soft Thinking, improving pass@1 accuracy by up to 2.48 points
while simultaneously reducing token usage by up to 22.4% compared to standard
CoT. Qualitative analysis further reveals that Soft Thinking outputs remain
highly interpretable and readable, highlighting the potential of Soft Thinking
to break the inherent bottleneck of discrete language-based reasoning. Code is
available at https://github.com/eric-ai-lab/Soft-Thinking.


## AI 摘要

当前推理模型受限于离散语言标记，无法像人类一样进行抽象灵活的思考。本文提出"软思考"(Soft Thinking)方法，通过在连续概念空间中生成混合标记嵌入的抽象概念标记，突破离散边界限制。这些概念标记能同时包含多重含义，隐式探索多种推理路径，从而更高效地收敛到正确答案。实验表明，相比标准思维链方法，该方案在数学和编程基准测试中最高提升2.48%准确率，同时减少22.4%的标记使用量。定性分析显示输出仍保持高度可解释性，有望突破基于离散语言的推理瓶颈。代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-23T00:02:15Z
- **目录日期**: 2025-05-23
