# Unlocking LLM Repair Capabilities in Low-Resource Programming Languages Through Cross-Language Translation and Multi-Agent Refinement

**URL**: http://arxiv.org/abs/2503.22512v3

## 原始摘要

Recent advances in leveraging LLMs for APR have demonstrated impressive
capabilities in fixing software defects. However, current LLM-based approaches
predominantly focus on mainstream programming languages like Java and Python,
neglecting less prevalent but emerging languages such as Rust due to expensive
training resources, limited datasets, and insufficient community support. This
narrow focus creates a significant gap in repair capabilities across the
programming language spectrum, where the full potential of LLMs for
comprehensive multilingual program repair remains largely unexplored. To
address this limitation, we introduce a novel cross-language program repair
approach LANTERN that leverages LLMs' differential proficiency across languages
through a multi-agent iterative repair paradigm. Our technique strategically
translates defective code from languages where LLMs exhibit weaker repair
capabilities to languages where they demonstrate stronger performance, without
requiring additional training. A key innovation of our approach is an LLM-based
decision-making system that dynamically selects optimal target languages based
on bug characteristics and continuously incorporates feedback from previous
repair attempts. We evaluate our method on xCodeEval, a comprehensive
multilingual benchmark comprising 5,068 bugs across 11 programming languages.
Results demonstrate significant enhancement in repair effectiveness,
particularly for underrepresented languages, with Rust showing a 22.09%
improvement in Pass@10 metrics. Our research provides the first empirical
evidence that cross-language translation significantly expands the repair
capabilities of LLMs and effectively bridges the performance gap between
programming languages with different levels of popularity, opening new avenues
for truly language-agnostic automated program repair.


## AI 摘要

近期研究表明，大型语言模型（LLM）在自动程序修复（APR）方面展现出强大能力，但现有方法主要集中于主流语言（如Java、Python），对新兴语言（如Rust）支持不足。为此，研究者提出跨语言修复方法LANTERN，通过多智能体迭代修复范式，将缺陷代码从LLM修复能力较弱的语言翻译至表现更强的语言，无需额外训练。该方法采用基于LLM的动态决策系统，根据错误特征选择最优目标语言。在11种语言的5,068个错误测试中，Rust的修复成功率提升22.09%，首次证明跨语言翻译能显著提升LLM的多语言修复能力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-18T20:01:53Z
- **目录日期**: 2025-04-18
