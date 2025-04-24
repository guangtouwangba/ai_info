# OptimAI: Optimization from Natural Language Using LLM-Powered AI Agents

**URL**: http://arxiv.org/abs/2504.16918v1

## 原始摘要

Optimization plays a vital role in scientific research and practical
applications, but formulating a concrete optimization problem described in
natural language into a mathematical form and selecting a suitable solver to
solve the problem requires substantial domain expertise. We introduce
\textbf{OptimAI}, a framework for solving \underline{Optim}ization problems
described in natural language by leveraging LLM-powered \underline{AI} agents,
achieving superior performance over current state-of-the-art methods. Our
framework is built upon four key roles: (1) a \emph{formulator} that translates
natural language problem descriptions into precise mathematical formulations;
(2) a \emph{planner} that constructs a high-level solution strategy prior to
execution; and (3) a \emph{coder} and a \emph{code critic} capable of
interacting with the environment and reflecting on outcomes to refine future
actions. Ablation studies confirm that all roles are essential; removing the
planner or code critic results in $5.8\times$ and $3.1\times$ drops in
productivity, respectively. Furthermore, we introduce UCB-based debug
scheduling to dynamically switch between alternative plans, yielding an
additional $3.3\times$ productivity gain. Our design emphasizes multi-agent
collaboration, allowing us to conveniently explore the synergistic effect of
combining diverse models within a unified system. Our approach attains 88.1\%
accuracy on the NLP4LP dataset and 71.2\% on the Optibench (non-linear w/o
table) subset, reducing error rates by 58\% and 50\% respectively over prior
best results.


## AI 摘要

研究人员提出了OptimAI框架，利用LLM驱动的AI智能体解决自然语言描述的优化问题，性能超越现有方法。该框架包含四个核心角色：(1) 问题表述器将自然语言转化为数学公式；(2) 规划器制定高层解决方案；(3) 编码器和代码评审器通过环境交互优化执行。消融实验表明，缺少规划器或评审器会分别导致5.8倍和3.1倍的效率下降。引入UCB调试调度策略带来额外3.3倍效率提升。该系统在NLP4LP和Optibench数据集上分别达到88.1%和71.2%准确率，错误率较之前最优结果降低58%和50%。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-24T10:01:13Z
- **目录日期**: 2025-04-24
