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

这篇论文研究了大型语言模型(LLMs)是否能在句子表征中捕捉主题角色(施事者、受事者等)。实验发现，与人类判断不同，LLMs的句子相似性主要反映句法而非主题角色关系。虽然模型整体表征中缺乏明确的主题角色信息，但某些注意力头确实能独立于句法结构识别这些关系。这表明LLMs能够提取主题角色信息，但相比人类，这类信息对其表征的影响较弱。研究聚焦语言理解的核心层面，揭示了LLMs通过词预测训练获得的语言表征特点。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-24T18:01:45Z
- **目录日期**: 2025-04-24
