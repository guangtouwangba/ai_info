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

这项研究探讨了大语言模型(LLMs)是否能在句子表征中捕捉主题角色(施事者与受事者关系)。实验发现，与人类判断不同，LLMs的句子相似性表征主要反映句法而非主题角色关系。虽然模型整体表征中难以检测到主题角色信息，但某些注意力头确实能独立于句法结构识别这些关系。研究表明，LLMs能够提取主题角色信息，但这些信息对其表征的影响程度远低于人类处理语言时的表现。这揭示了LLMs语言理解能力的局限性，特别是在语义角色分配方面与人类认知的差异。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-24T09:01:47Z
- **目录日期**: 2025-04-24
