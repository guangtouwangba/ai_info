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

该研究探讨了语言模型（LM）在探索和推断因果关系时的能力，发现模型能可靠推断常见的析因关系，但系统性地难以处理同样有证据支持的合因关系。这种“析因偏差”存在于不同模型家族、规模和提示策略中，且任务复杂度越高，表现越差。有趣的是，人类成人也存在类似偏差，表明LM可能从训练数据中继承了深层次的推理启发式。研究量化了LM与人类的相似性，发现其推理模式更接近成人而非儿童。最后，提出了一种测试时采样方法，显著减少了析因偏差，提升了因果推理的严谨性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-15T05:00:56Z
- **目录日期**: 2025-05-15
