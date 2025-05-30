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

该研究探讨了语言模型（LM）在探索和推断因果关系时的能力，发现LM能可靠推断常见的析取因果关系，但系统性地难以处理较少见但证据充分的合取关系，表现出"析取偏差"。这种偏差在不同模型家族、规模和提示策略中均存在，且随任务复杂性增加而加剧。有趣的是，人类成人也表现出类似偏差，表明LM可能从训练数据中继承了这种推理启发式。研究量化了LM与人类的相似性，发现LM的推理模式更接近成人而非儿童。最后，提出了一种测试时采样方法，显著减少了析取偏差，提升了LM的科学推理能力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-16T01:28:31Z
- **目录日期**: 2025-05-16
