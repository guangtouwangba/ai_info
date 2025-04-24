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

研究表明，大型语言模型（LLMs）的核心训练目标（词语预测）并未使其句法表征充分捕捉人类语言中的"施受关系"（thematic roles）。实验发现，LLMs的句子表征相似性主要反映句法特征，而非施事者与受事者的对应关系（人类判断的核心指标）。尽管部分注意力头能独立于句法提取施受关系信息，但整体上这类信息对模型表征的影响远弱于人类。这表明LLMs虽具备提取施受关系的能力，但其语言理解机制与人类存在本质差异。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-24T10:01:51Z
- **目录日期**: 2025-04-24
