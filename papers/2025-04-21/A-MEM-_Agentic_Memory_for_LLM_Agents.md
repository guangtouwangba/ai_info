# A-MEM: Agentic Memory for LLM Agents

**URL**: http://arxiv.org/abs/2502.12110v5

## 原始摘要

While large language model (LLM) agents can effectively use external tools
for complex real-world tasks, they require memory systems to leverage
historical experiences. Current memory systems enable basic storage and
retrieval but lack sophisticated memory organization, despite recent attempts
to incorporate graph databases. Moreover, these systems' fixed operations and
structures limit their adaptability across diverse tasks. To address this
limitation, this paper proposes a novel agentic memory system for LLM agents
that can dynamically organize memories in an agentic way. Following the basic
principles of the Zettelkasten method, we designed our memory system to create
interconnected knowledge networks through dynamic indexing and linking. When a
new memory is added, we generate a comprehensive note containing multiple
structured attributes, including contextual descriptions, keywords, and tags.
The system then analyzes historical memories to identify relevant connections,
establishing links where meaningful similarities exist. Additionally, this
process enables memory evolution - as new memories are integrated, they can
trigger updates to the contextual representations and attributes of existing
historical memories, allowing the memory network to continuously refine its
understanding. Our approach combines the structured organization principles of
Zettelkasten with the flexibility of agent-driven decision making, allowing for
more adaptive and context-aware memory management. Empirical experiments on six
foundation models show superior improvement against existing SOTA baselines.
The source code for evaluating performance is available at
https://github.com/WujiangXu/AgenticMemory, while the source code of agentic
memory system is available at https://github.com/agiresearch/A-mem.


## AI 摘要

本文提出了一种新型的LLM智能体记忆系统，能够动态组织记忆并建立知识网络。该系统基于Zettelkasten方法，通过动态索引和链接创建互连的记忆网络。每当新增记忆时，系统会生成包含上下文描述、关键词和标签的结构化笔记，并分析历史记忆建立关联。新记忆的加入还能触发对已有记忆的更新，实现记忆网络的持续优化。实验表明，该系统在六个基础模型上表现优于现有方法。源代码已开源，包括性能评估工具和记忆系统实现。这种结合结构化组织和智能决策的方法提升了记忆管理的适应性和情境感知能力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-21T08:02:39Z
- **目录日期**: 2025-04-21
