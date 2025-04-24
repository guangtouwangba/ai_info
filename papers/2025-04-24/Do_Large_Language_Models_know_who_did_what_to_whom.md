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

该研究探讨了大型语言模型（LLMs）是否能在句子表征中捕捉主题角色（如施事与受事关系）。实验发现，与人类判断不同，LLMs的句子表征相似性主要反映句法而非主题角色的一致性。尽管部分注意力头能独立于句法提取主题角色信息，但整体表征中这类信息较弱。结果表明，LLMs虽具备提取主题角色的能力，但其表征对该信息的依赖程度远低于人类。研究聚焦语言相关的理解能力，为LLMs的语义处理机制提供了新视角。（100字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-24T11:01:54Z
- **目录日期**: 2025-04-24
