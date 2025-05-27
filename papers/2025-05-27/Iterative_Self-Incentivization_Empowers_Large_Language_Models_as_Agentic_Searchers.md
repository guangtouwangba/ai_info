# Iterative Self-Incentivization Empowers Large Language Models as Agentic Searchers

**URL**: http://arxiv.org/abs/2505.20128v1

## 原始摘要

Large language models (LLMs) have been widely integrated into information
retrieval to advance traditional techniques. However, effectively enabling LLMs
to seek accurate knowledge in complex tasks remains a challenge due to the
complexity of multi-hop queries as well as the irrelevant retrieved content. To
address these limitations, we propose EXSEARCH, an agentic search framework,
where the LLM learns to retrieve useful information as the reasoning unfolds
through a self-incentivized process. At each step, the LLM decides what to
retrieve (thinking), triggers an external retriever (search), and extracts
fine-grained evidence (recording) to support next-step reasoning. To enable LLM
with this capability, EXSEARCH adopts a Generalized Expectation-Maximization
algorithm. In the E-step, the LLM generates multiple search trajectories and
assigns an importance weight to each; the M-step trains the LLM on them with a
re-weighted loss function. This creates a self-incentivized loop, where the LLM
iteratively learns from its own generated data, progressively improving itself
for search. We further theoretically analyze this training process,
establishing convergence guarantees. Extensive experiments on four
knowledge-intensive benchmarks show that EXSEARCH substantially outperforms
baselines, e.g., +7.8% improvement on exact match score. Motivated by these
promising results, we introduce EXSEARCH-Zoo, an extension that extends our
method to broader scenarios, to facilitate future work.


## AI 摘要

该研究提出EXSEARCH框架，通过自激励学习机制提升大语言模型(LLMs)在多跳查询中的信息检索能力。框架采用广义期望最大化算法：E步生成带权重的搜索轨迹，M步通过重加权损失函数训练模型，形成自我提升循环。理论分析证实了训练过程的收敛性。在四个知识密集型基准测试中，EXSEARCH显著优于基线方法（如准确匹配分数提升7.8%）。研究还扩展出EXSEARCH-Zoo以适用更广泛场景。该方法通过"思考-搜索-记录"的迭代过程，有效解决了复杂任务中的无关内容检索问题。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-27T17:01:57Z
- **目录日期**: 2025-05-27
