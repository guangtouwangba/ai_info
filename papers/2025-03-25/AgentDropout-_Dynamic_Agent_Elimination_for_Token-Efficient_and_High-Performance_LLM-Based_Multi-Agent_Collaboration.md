# AgentDropout: Dynamic Agent Elimination for Token-Efficient and High-Performance LLM-Based Multi-Agent Collaboration

**URL**: http://arxiv.org/abs/2503.18891v1

## 原始摘要

Multi-agent systems (MAS) based on large language models (LLMs) have
demonstrated significant potential in collaborative problem-solving. However,
they still face substantial challenges of low communication efficiency and
suboptimal task performance, making the careful design of the agents'
communication topologies particularly important. Inspired by the management
theory that roles in an efficient team are often dynamically adjusted, we
propose AgentDropout, which identifies redundant agents and communication
across different communication rounds by optimizing the adjacency matrices of
the communication graphs and eliminates them to enhance both token efficiency
and task performance. Compared to state-of-the-art methods, AgentDropout
achieves an average reduction of 21.6% in prompt token consumption and 18.4% in
completion token consumption, along with a performance improvement of 1.14 on
the tasks. Furthermore, the extended experiments demonstrate that AgentDropout
achieves notable domain transferability and structure robustness, revealing its
reliability and effectiveness. We release our code at
https://github.com/wangzx1219/AgentDropout.


## AI 摘要

基于大语言模型(LLM)的多智能体系统(MAS)在协作解决问题方面展现出巨大潜力，但仍面临通信效率低和任务表现欠佳的挑战。受高效团队中角色动态调整的管理理论启发，研究者提出AgentDropout方法，通过优化通信图的邻接矩阵来识别冗余智能体和通信轮次，从而提升令牌效率和任务表现。相比现有方法，该方法平均减少21.6%的提示令牌和18.4%的完成令牌消耗，任务表现提升1.14分。实验还验证了该方法具有良好的领域迁移性和结构鲁棒性。代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-25T08:01:31Z
- **目录日期**: 2025-03-25
