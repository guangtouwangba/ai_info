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

研究表明，大型语言模型（LLMs）的核心训练目标（词语预测）并未使其句子表征像人类那样有效捕捉主题角色（施事与受事关系）。实验发现，LLMs的句子相似性判断主要依赖句法而非语义角色分配，且主题角色信息未显著体现在隐藏单元中。不过，部分注意力头能独立于句法结构识别主题角色。这说明LLMs虽具备提取主题角色的能力，但该信息对其表征的影响远弱于人类语言处理机制。研究揭示了LLMs与人类在语言理解层面的关键差异。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-25T01:29:19Z
- **目录日期**: 2025-04-25
