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

该研究探讨了语言模型(LM)在因果推理中的表现，发现模型能准确推断常见的"析取"因果关系，但对同样有证据支持的"合取"关系存在系统性困难。这种"析取偏好"存在于不同规模和类型的模型中，且任务复杂度越高表现越差。有趣的是，人类成人也存在类似偏见，表明LM可能从训练数据中继承了这种深层推理启发式。研究量化了LM与人类推理的相似性，发现其模式更接近成人而非儿童。最后，作者提出一种采样方法，通过显式排除假设显著减少了这种偏见，提升了模型的科学推理能力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-15T09:01:08Z
- **目录日期**: 2025-05-15
