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

这篇研究探讨了大语言模型（LLMs）是否能在句子中准确推断"谁对谁做了什么"（主题角色）。实验分析了四种LLM的句子表征，发现与人类不同，LLMs的表征相似性主要反映句法而非主题角色（如施事与受事是否相同或颠倒）。虽然某些注意力头能独立于句法捕捉主题角色信息，但整体上LLMs的表征中主题角色信息较弱。研究表明，尽管LLMs能够提取主题角色，但相比人类，这类信息对其表征的影响明显更弱。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-24T05:01:52Z
- **目录日期**: 2025-04-24
