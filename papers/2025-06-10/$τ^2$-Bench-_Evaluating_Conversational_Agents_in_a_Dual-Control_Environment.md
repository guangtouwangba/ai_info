# $τ^2$-Bench: Evaluating Conversational Agents in a Dual-Control Environment

**URL**: http://arxiv.org/abs/2506.07982v1

## 原始摘要

Existing benchmarks for conversational AI agents simulate single-control
environments, where only the AI agent can use tools to interact with the world,
while the user remains a passive information provider. This differs from
real-world scenarios like technical support, where users need to actively
participate in modifying the state of the (shared) world. In order to address
this gap, we introduce $\tau^2$-bench, with four key contributions:
  1) A novel Telecom dual-control domain modeled as a Dec-POMDP, where both
agent and user make use of tools to act in a shared, dynamic environment that
tests both agent coordination and communication,
  2) A compositional task generator that programmatically creates diverse,
verifiable tasks from atomic components, ensuring domain coverage and
controlled complexity,
  3) A reliable user simulator tightly coupled with the environment, whose
behavior is constrained by tools and observable states, improving simulation
fidelity,
  4) Fine-grained analysis of agent performance through multiple ablations
including separating errors arising from reasoning vs
communication/coordination.
  In particular, our experiments show significant performance drops when agents
shift from no-user to dual-control, highlighting the challenges of guiding
users. Overall, $\tau^2$-bench provides a controlled testbed for agents that
must both reason effectively and guide user actions.


## AI 摘要

现有的对话AI评测基准主要模拟单控制环境，仅AI能使用工具与世界互动，而用户是被动的信息提供者。为弥补这一差距，研究者提出了τ²-bench基准，包含四个关键贡献：1) 电信领域的双控制环境建模为Dec-POMDP，测试AI与用户的协作；2) 可组合的任务生成器，确保多样性和可控复杂度；3) 与工具和环境紧密耦合的高保真用户模拟器；4) 细粒度性能分析框架。实验表明，从单控制转向双控制时AI性能显著下降，突显了指导用户操作的挑战。该基准为需要推理和指导用户的AI提供了可控测试平台。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-10T11:01:34Z
- **目录日期**: 2025-06-10
