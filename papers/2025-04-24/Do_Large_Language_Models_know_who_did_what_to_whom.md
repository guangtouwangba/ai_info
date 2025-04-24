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

这篇论文研究了大型语言模型(LLMs)是否能在句子表征中捕捉主题角色(施事者和受事者关系)。实验发现，与人类判断不同，LLMs的句子表征相似性主要反映句法而非主题角色关系。虽然模型隐藏单元中缺乏主题角色信息，但某些注意力头确实能独立于句法结构提取主题角色信息。结果表明，LLMs能够识别主题角色，但这些信息对模型表征的影响程度远低于人类。这说明LLMs通过词语预测训练可以学习到部分语言理解能力，但其表征方式与人类存在显著差异。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-24T20:01:47Z
- **目录日期**: 2025-04-24
