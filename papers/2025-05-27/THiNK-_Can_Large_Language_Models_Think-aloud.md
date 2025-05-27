# THiNK: Can Large Language Models Think-aloud?

**URL**: http://arxiv.org/abs/2505.20184v1

## 原始摘要

Assessing higher-order thinking skills in large language models (LLMs)
remains a fundamental challenge, especially in tasks that go beyond
surface-level accuracy. In this work, we propose THiNK (Testing Higher-order
Notion of Knowledge), a multi-agent, feedback-driven evaluation framework
grounded in Bloom's Taxonomy. THiNK frames reasoning assessment as an iterative
task of problem generation, critique, and revision, encouraging LLMs to
think-aloud through step-by-step reflection and refinement. This enables a
systematic evaluation of both lower-order (e.g., remember, understand) and
higher-order (e.g., evaluate, create) thinking skills. We apply THiNK to seven
state-of-the-art LLMs and perform a detailed cognitive analysis of their
outputs. Results reveal that while models reliably perform lower-order
categories well, they struggle with applying knowledge in realistic contexts
and exhibit limited abstraction. Structured feedback loops significantly
improve reasoning performance, particularly in higher-order thinking.
Qualitative evaluations further confirm that THiNK-guided outputs better align
with domain logic and problem structure. The code of our framework provides a
scalable methodology for probing and enhancing LLM reasoning, offering new
directions for evaluation grounded in learning science, which is available at
our GitHub repository.


## AI 摘要

本研究提出THiNK框架，基于布鲁姆分类法，通过多智能体反馈机制评估大语言模型的高阶思维能力。该框架采用"生成-批判-修正"的迭代流程，引导模型分步反思，系统评估从记忆理解到评估创造的不同认知层次。实验表明，当前模型在低阶任务表现良好，但在实际应用和抽象推理方面存在局限。结构化反馈显著提升推理能力，尤其对高阶思维帮助明显。THiNK输出的解决方案更符合领域逻辑，其开源代码为基于学习科学的模型评估提供了可扩展方法。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-27T09:02:00Z
- **目录日期**: 2025-05-27
