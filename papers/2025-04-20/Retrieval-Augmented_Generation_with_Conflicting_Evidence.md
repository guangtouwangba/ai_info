# Retrieval-Augmented Generation with Conflicting Evidence

**URL**: http://arxiv.org/abs/2504.13079v1

## 原始摘要

Large language model (LLM) agents are increasingly employing
retrieval-augmented generation (RAG) to improve the factuality of their
responses. However, in practice, these systems often need to handle ambiguous
user queries and potentially conflicting information from multiple sources
while also suppressing inaccurate information from noisy or irrelevant
documents. Prior work has generally studied and addressed these challenges in
isolation, considering only one aspect at a time, such as handling ambiguity or
robustness to noise and misinformation. We instead consider multiple factors
simultaneously, proposing (i) RAMDocs (Retrieval with Ambiguity and
Misinformation in Documents), a new dataset that simulates complex and
realistic scenarios for conflicting evidence for a user query, including
ambiguity, misinformation, and noise; and (ii) MADAM-RAG, a multi-agent
approach in which LLM agents debate over the merits of an answer over multiple
rounds, allowing an aggregator to collate responses corresponding to
disambiguated entities while discarding misinformation and noise, thereby
handling diverse sources of conflict jointly. We demonstrate the effectiveness
of MADAM-RAG using both closed and open-source models on AmbigDocs -- which
requires presenting all valid answers for ambiguous queries -- improving over
strong RAG baselines by up to 11.40% and on FaithEval -- which requires
suppressing misinformation -- where we improve by up to 15.80% (absolute) with
Llama3.3-70B-Instruct. Furthermore, we find that RAMDocs poses a challenge for
existing RAG baselines (Llama3.3-70B-Instruct only obtains 32.60 exact match
score). While MADAM-RAG begins to address these conflicting factors, our
analysis indicates that a substantial gap remains especially when increasing
the level of imbalance in supporting evidence and misinformation.


## AI 摘要

该研究提出RAMDocs数据集和MADAM-RAG多智能体系统，用于解决大语言模型(LLM)在检索增强生成(RAG)中面临的多重挑战：用户查询歧义、多源信息冲突以及噪声干扰。MADAM-RAG通过多轮辩论机制，能同时处理歧义查询（AmbigDocs任务提升11.40%）和过滤错误信息（FaithEval任务提升15.80%）。实验显示现有RAG基线在RAMDocs上表现不佳（Llama3仅32.60分），且当证据支持不平衡时仍有较大改进空间。这项工作首次综合处理了RAG系统中的多重冲突因素。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-20T01:29:58Z
- **目录日期**: 2025-04-20
