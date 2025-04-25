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

该研究探讨了大型语言模型(LLMs)是否能在句子表征中捕捉"谁对谁做了什么"(主题角色)这一语言理解核心能力。实验发现：1)与人类不同，LLMs的句子表征相似性主要反映句法而非主题角色；2)隐藏单元中基本不存在主题角色信息；3)部分注意力头能独立于句法捕获主题角色信息。结论表明：虽然LLMs具备提取主题角色的能力，但相比人类，这类信息对其表征的影响较弱。这揭示了LLMs通过词语预测训练形成的语言理解与人类存在系统性差异。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-25T00:01:51Z
- **目录日期**: 2025-04-25
