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

现有对话AI基准测试多模拟单控制环境，仅AI能使用工具与世界互动，而用户被动提供信息。为弥补这一差距，研究者提出τ²-bench基准：(1)创建电信领域双控制环境(Dec-POMDP模型)，要求AI与用户协同操作共享动态环境；(2)通过模块化任务生成器构建多样化可验证任务；(3)开发与环境紧密耦合的用户模拟器，提升仿真真实度；(4)细粒度分析AI表现，区分推理错误与协作/沟通错误。实验表明，从单控转向双控时AI性能显著下降，突显指导用户操作的挑战。该基准为需协同用户操作的AI系统提供了可控测试平台。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-10T23:01:20Z
- **目录日期**: 2025-06-10
