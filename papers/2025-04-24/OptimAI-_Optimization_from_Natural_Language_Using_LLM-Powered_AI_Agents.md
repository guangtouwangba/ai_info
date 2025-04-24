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

OptimAI是一个基于多智能体协作的框架，通过LLM驱动的AI代理将自然语言描述的优化问题转化为数学形式并求解。其核心包含四个角色：(1) **问题建模器**（formulator）将自然语言转换为数学公式；(2) **策略规划器**（planner）制定高层求解策略；(3) **代码生成器**（coder）和**代码评审器**（code critic）通过环境交互优化执行。实验表明，缺少规划器或评审器会导致效率下降5.8倍和3.1倍。结合UCB动态调试策略后，效率进一步提升3.3倍。在NLP4LP和Optibench数据集上，错误率分别降低58%和50%，显著优于现有方法。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-24T06:01:10Z
- **目录日期**: 2025-04-24
