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

大型语言模型（LLMs）常被批评缺乏语言理解能力。本研究聚焦于与语言紧密相关的主题角色理解（如识别句子中的施事和受事）。实验分析了四种LLMs的句子表征，发现其表征相似性主要反映句法而非主题角色分配，且隐藏单元中主题角色信息有限。然而，部分注意力头能独立于句法捕获主题角色信息。结果表明，LLMs虽能提取主题角色，但其表征受此信息的影响远弱于人类。研究揭示了LLMs在语言理解方面的局限性，同时指出注意力机制在主题角色识别中的潜在作用。（100字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-24T04:01:56Z
- **目录日期**: 2025-04-24
