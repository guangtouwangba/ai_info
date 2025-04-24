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

这篇研究探讨了大语言模型(LLMs)是否真正理解语言中的"谁对谁做了什么"(主题角色)关系。实验发现，与人类不同，LLMs的句子表征主要反映句法相似性而非主题角色分配。虽然模型整体表征未能有效捕捉主题角色信息，但某些注意力头确实能独立于句法结构识别这些关系。研究表明，LLMs能够提取主题角色信息，但相比人类，这些信息对模型表征的影响较弱。这显示LLMs的核心训练目标(词语预测)并不必然导致对语言中关键语义关系的深入理解。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-24T21:01:42Z
- **目录日期**: 2025-04-24
