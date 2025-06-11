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

现有对话AI基准测试主要模拟单控制环境，仅AI可使用工具与世界交互，而用户被动提供信息。为弥补这一与真实场景（如技术支持）的差距，研究者提出τ²-bench，包含四大创新：1) 电信领域双控制模型（Dec-POMDP框架），测试AI与用户协同工具操作；2) 可组合任务生成器，通过原子组件构建多样化可验证任务；3) 与环境紧密耦合的高保真用户模拟器；4) 细粒度性能分析（区分推理与协作错误）。实验显示双控制场景下AI性能显著下降，凸显指导用户的挑战。该基准为需协同用户行动的AI提供了可控测试平台。（99字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-11T02:31:36Z
- **目录日期**: 2025-06-11
