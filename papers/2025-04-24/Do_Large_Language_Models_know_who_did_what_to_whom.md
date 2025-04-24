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

这篇论文研究了大型语言模型(LLMs)是否能理解句子中的"谁对谁做了什么"(主题角色)。通过分析四个LLMs的句子表征发现：与人类不同，LLMs的句子相似性判断主要基于句法而非主题角色分配；主题角色信息在隐藏单元中表现微弱，但某些注意力头能独立于句法捕捉这些信息。结论表明，LLMs虽然能提取主题角色信息，但其表征受此影响的程度远低于人类。研究聚焦语言理解的核心层面，揭示了LLMs在语义角色分配方面的局限性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-24T23:01:40Z
- **目录日期**: 2025-04-24
