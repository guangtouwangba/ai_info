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

大型语言模型(LLMs)常被批评缺乏语言理解能力。本研究聚焦与语言紧密相关的主题角色理解(即识别句子中"谁对谁做了什么")。通过分析四个LLM的句子表征发现：与人类不同，LLM的表征相似性主要反映句法而非主题角色关系；主题角色信息在隐藏单元中表现微弱，但某些注意力头能独立于句法捕捉该信息。结果表明，LLM虽能提取主题角色信息，但其表征受此影响的程度远低于人类。这揭示了LLM语言理解与人类的关键差异。(99字)

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-25T04:01:47Z
- **目录日期**: 2025-04-25
