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

现有对话AI基准测试通常模拟单控制环境，仅AI代理能使用工具与世界互动，而用户是被动信息提供者。这与技术支持等现实场景不符，用户需主动参与修改共享世界状态。为此，研究者提出τ²-bench，包含四大创新：1) 电信双控制领域建模为Dec-POMDP，测试代理协调与沟通；2) 组合任务生成器创建多样化可验证任务；3) 与环境紧密耦合的高保真用户模拟器；4) 细粒度性能分析框架。实验表明，从单控切换到双控时AI性能显著下降，凸显指导用户操作的挑战。该基准为需高效推理和用户引导的AI提供了可控测试平台。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-10T19:01:22Z
- **目录日期**: 2025-06-10
