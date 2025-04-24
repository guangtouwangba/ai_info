# Do Large Language Models know who did what to whom?

**URL**: http://arxiv.org/abs/2504.16884v1

## 原始摘要

Large Language Models (LLMs) are commonly criticized for not understanding
language. However, many critiques focus on cognitive abilities that, in humans,
are distinct from language processing. Here, we instead study a kind of
understanding tightly linked to language: inferring who did what to whom
(thematic roles) in a sentence. Does the central training objective of
LLMs-word prediction-result in sentence representations that capture thematic
roles? In two experiments, we characterized sentence representations in four
LLMs. In contrast to human similarity judgments, in LLMs the overall
representational similarity of sentence pairs reflected syntactic similarity
but not whether their agent and patient assignments were identical vs.
reversed. Furthermore, we found little evidence that thematic role information
was available in any subset of hidden units. However, some attention heads
robustly captured thematic roles, independently of syntax. Therefore, LLMs can
extract thematic roles but, relative to humans, this information influences
their representations more weakly.


## AI 摘要

这篇研究探讨了大语言模型(LLMs)是否能够理解句子中的"谁对谁做了什么"(主题角色)。实验发现，与人类不同，LLMs的句子表征主要反映句法相似性，而非主题角色分配。虽然模型整体表征中缺乏主题角色信息，但某些注意力头能独立于句法结构捕获这些信息。结果表明，LLMs能够提取主题角色，但相比人类，这些信息对其表征的影响较弱。研究聚焦语言理解的核心层面，揭示了LLMs在语义角色识别方面的局限性及其与人类语言处理的差异。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-24T16:02:14Z
- **目录日期**: 2025-04-24
