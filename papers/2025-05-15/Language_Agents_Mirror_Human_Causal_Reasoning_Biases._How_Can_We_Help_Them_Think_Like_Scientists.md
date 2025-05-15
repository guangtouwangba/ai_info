# Language Agents Mirror Human Causal Reasoning Biases. How Can We Help Them Think Like Scientists?

**URL**: http://arxiv.org/abs/2505.09614v1

## 原始摘要

Language model (LM) agents are increasingly used as autonomous
decision-makers who need to actively gather information to guide their
decisions. A crucial cognitive skill for such agents is the efficient
exploration and understanding of the causal structure of the world -- key to
robust, scientifically grounded reasoning. Yet, it remains unclear whether LMs
possess this capability or exhibit systematic biases leading to erroneous
conclusions. In this work, we examine LMs' ability to explore and infer causal
relationships, using the well-established "Blicket Test" paradigm from
developmental psychology. We find that LMs reliably infer the common, intuitive
disjunctive causal relationships but systematically struggle with the unusual,
yet equally (or sometimes even more) evidenced conjunctive ones. This
"disjunctive bias" persists across model families, sizes, and prompting
strategies, and performance further declines as task complexity increases.
Interestingly, an analogous bias appears in human adults, suggesting that LMs
may have inherited deep-seated reasoning heuristics from their training data.
To this end, we quantify similarities between LMs and humans, finding that LMs
exhibit adult-like inference profiles (but not children-like). Finally, we
propose a test-time sampling method which explicitly samples and eliminates
hypotheses about causal relationships from the LM. This scalable approach
significantly reduces the disjunctive bias and moves LMs closer to the goal of
scientific, causally rigorous reasoning.


## AI 摘要

该研究探讨了语言模型（LM）在因果推理中的能力，发现模型能可靠推断常见的析取因果关系，但系统性地难以处理较少见的合取关系，表现出"析取偏见"。这种偏见存在于不同模型家族、规模和提示策略中，且随任务复杂度增加而加剧。有趣的是，类似偏见也存在于人类成人中，表明LM可能从训练数据继承了深层次的推理启发式。研究量化了LM与人类推理的相似性，发现LM表现出成人而非儿童的推理模式。最后，提出一种测试时采样方法，通过显式采样和排除因果假设，显著减少了析取偏见。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-15T06:01:03Z
- **目录日期**: 2025-05-15
