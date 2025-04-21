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

该论文提出了一种新型的LLM智能体记忆系统，采用类似Zettelkasten方法的动态组织方式，通过建立索引和链接形成互联知识网络。系统会为每个新记忆创建包含上下文描述、关键词和标签的结构化笔记，并分析历史记忆建立关联。新记忆的加入还能触发对已有记忆的更新，实现知识网络的持续优化。这种结合Zettelkasten结构化原则与智能体决策灵活性的方法，在6个基础模型上的实验表现优于现有技术。相关代码已开源，包含性能评估工具和记忆系统实现。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-21T19:02:22Z
- **目录日期**: 2025-04-21
