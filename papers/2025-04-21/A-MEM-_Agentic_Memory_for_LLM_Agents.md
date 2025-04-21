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

本文提出了一种新型的LLM智能体记忆系统，采用类似Zettelkasten方法的动态索引和链接机制组织记忆。该系统为每个新记忆创建包含上下文描述、关键词和标签的结构化笔记，并通过分析历史记忆建立关联链接。新记忆的加入能触发现有记忆的上下文表征更新，实现知识网络的持续优化。该方法结合了Zettelkasten的结构化原则与智能体驱动的决策灵活性，在六个基础模型上的实验表明其性能优于现有最优基线。相关评估代码和系统源码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-21T22:02:15Z
- **目录日期**: 2025-04-21
