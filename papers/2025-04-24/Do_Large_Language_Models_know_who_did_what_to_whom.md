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

这篇论文研究了大型语言模型(LLMs)是否能理解句子中的"谁对谁做了什么"(主题角色)。通过两个实验分析四个LLM的句子表征发现：与人类不同，LLMs的句子相似性判断主要基于句法而非主题角色；模型隐藏单元中几乎没有主题角色信息，但某些注意力头能独立于句法捕捉主题角色。结论表明LLMs能够提取主题角色信息，但相比人类，这些信息对模型表征的影响较弱。这揭示了LLMs通过词预测训练获得的语言理解能力与人类存在差异。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-24T15:02:00Z
- **目录日期**: 2025-04-24
